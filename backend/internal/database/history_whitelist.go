package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mtn-server/backend/internal/whitelist"
)

const whitelistTimeFormat = "%Y-%m-%d %H:%i:%s"

// ListActive returns active whitelist entries, optionally scoped by edition.
func (h *HistoryDB) ListActive(ctx context.Context, edition whitelist.Edition) ([]whitelist.Entry, error) {
	query := `
		SELECT id,
		       COALESCE(qq_id, '') AS qq_id,
		       edition,
		       player_name,
		       normalized_player_name,
		       active,
		       DATE_FORMAT(created_at, ?) AS created_at,
		       DATE_FORMAT(updated_at, ?) AS updated_at,
		       DATE_FORMAT(removed_at, ?) AS removed_at
		FROM whitelist_entries
		WHERE active = TRUE
	`
	args := []any{whitelistTimeFormat, whitelistTimeFormat, whitelistTimeFormat}
	if edition != whitelist.EditionAll {
		query += " AND edition = ?"
		args = append(args, string(edition))
	}
	query += " ORDER BY updated_at DESC, id DESC"

	rows, err := h.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query active whitelist entries: %w", err)
	}
	defer rows.Close()

	var entries []whitelist.Entry
	for rows.Next() {
		entry, err := scanWhitelistEntry(rows)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, rows.Err()
}

// FindActive returns an active whitelist entry by edition and normalized player name.
func (h *HistoryDB) FindActive(ctx context.Context, edition whitelist.Edition, normalizedPlayerName string) (*whitelist.Entry, error) {
	query := `
		SELECT id,
		       COALESCE(qq_id, '') AS qq_id,
		       edition,
		       player_name,
		       normalized_player_name,
		       active,
		       DATE_FORMAT(created_at, ?) AS created_at,
		       DATE_FORMAT(updated_at, ?) AS updated_at,
		       DATE_FORMAT(removed_at, ?) AS removed_at
		FROM whitelist_entries
		WHERE edition = ? AND normalized_player_name = ? AND active = TRUE
		LIMIT 1
	`

	row := h.db.QueryRowContext(ctx, query, whitelistTimeFormat, whitelistTimeFormat, whitelistTimeFormat, string(edition), normalizedPlayerName)
	entry, err := scanWhitelistEntry(row)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// CountActiveByQQ returns the active whitelist count owned by a QQ ID.
func (h *HistoryDB) CountActiveByQQ(ctx context.Context, qqID string) (int, error) {
	if qqID == "" {
		return 0, nil
	}

	var count int
	query := `SELECT COUNT(*) FROM whitelist_entries WHERE qq_id = ? AND active = TRUE`
	if err := h.db.QueryRowContext(ctx, query, qqID).Scan(&count); err != nil {
		return 0, fmt.Errorf("count whitelist entries by qq: %w", err)
	}

	return count, nil
}

// UpsertLink stores the QQ-to-Minecraft account link history.
func (h *HistoryDB) UpsertLink(ctx context.Context, qqID string, edition whitelist.Edition, playerName string, normalizedPlayerName string) error {
	if qqID == "" {
		return nil
	}

	query := `
		INSERT INTO mc_account_links (qq_id, edition, player_name, normalized_player_name)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE player_name = VALUES(player_name), last_seen_at = CURRENT_TIMESTAMP
	`
	if _, err := h.db.ExecContext(ctx, query, qqID, string(edition), playerName, normalizedPlayerName); err != nil {
		return fmt.Errorf("upsert mc account link: %w", err)
	}

	return nil
}

// UpsertEntry creates or reactivates a whitelist mirror entry.
func (h *HistoryDB) UpsertEntry(ctx context.Context, qqID string, edition whitelist.Edition, playerName string, normalizedPlayerName string) (*whitelist.Entry, error) {
	query := `
		INSERT INTO whitelist_entries (qq_id, edition, player_name, normalized_player_name, active, removed_at)
		VALUES (?, ?, ?, ?, TRUE, NULL)
		ON DUPLICATE KEY UPDATE
			qq_id = VALUES(qq_id),
			player_name = VALUES(player_name),
			active = TRUE,
			removed_at = NULL,
			updated_at = CURRENT_TIMESTAMP
	`
	if _, err := h.db.ExecContext(ctx, query, nullableString(qqID), string(edition), playerName, normalizedPlayerName); err != nil {
		return nil, fmt.Errorf("upsert whitelist entry: %w", err)
	}

	entry, err := h.FindActive(ctx, edition, normalizedPlayerName)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, sql.ErrNoRows
	}

	return entry, nil
}

// DeactivateEntry marks an active whitelist mirror entry as removed.
func (h *HistoryDB) DeactivateEntry(ctx context.Context, id int64) error {
	query := `
		UPDATE whitelist_entries
		SET active = FALSE, removed_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND active = TRUE
	`
	if _, err := h.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("deactivate whitelist entry: %w", err)
	}

	return nil
}

// RecordAudit records a whitelist operation attempt.
func (h *HistoryDB) RecordAudit(ctx context.Context, record whitelist.AuditRecord) error {
	query := `
		INSERT INTO whitelist_audit (
			action, edition, player_name, normalized_player_name,
			qq_id, actor_qq_id, source, status, rcon_output, error_message
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := h.db.ExecContext(
		ctx,
		query,
		string(record.Action),
		string(record.Edition),
		record.PlayerName,
		record.NormalizedPlayerName,
		nullableString(record.QQID),
		nullableString(record.ActorQQID),
		string(record.Source),
		record.Status,
		nullableString(record.RCONOutput),
		nullableString(record.ErrorMessage),
	)
	if err != nil {
		return fmt.Errorf("insert whitelist audit: %w", err)
	}

	return nil
}

type whitelistEntryScanner interface {
	Scan(dest ...any) error
}

func scanWhitelistEntry(scanner whitelistEntryScanner) (whitelist.Entry, error) {
	var (
		entry     whitelist.Entry
		removedAt sql.NullString
	)
	err := scanner.Scan(
		&entry.ID,
		&entry.QQID,
		&entry.Edition,
		&entry.PlayerName,
		&entry.NormalizedPlayerName,
		&entry.Active,
		&entry.CreatedAt,
		&entry.UpdatedAt,
		&removedAt,
	)
	if err != nil {
		return whitelist.Entry{}, err
	}
	if removedAt.Valid {
		entry.RemovedAt = &removedAt.String
	}

	return entry, nil
}

func nullableString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: value, Valid: true}
}
