package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/mtn-server/backend/config"
)

// HistoryDB provides access to the presence history MySQL database.
type HistoryDB struct {
	db *sql.DB
}

// NewHistoryDB opens a connection to the presence history database.
func NewHistoryDB(cfg config.MySQLConfig) (*HistoryDB, error) {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open history db: %w", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping history db: %w", err)
	}

	return &HistoryDB{db: db}, nil
}

// EnsureSchema creates the required history tables if they do not exist.
func (h *HistoryDB) EnsureSchema(ctx context.Context) error {
	statements := []string{
		`
		CREATE TABLE IF NOT EXISTS server_hourly_presence (
			bucket_start DATETIME NOT NULL PRIMARY KEY,
			max_players INT NOT NULL
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS player_hourly_presence (
			player_uuid VARCHAR(36) NOT NULL,
			bucket_start DATETIME NOT NULL,
			PRIMARY KEY (player_uuid, bucket_start),
			KEY idx_player_hourly_presence_bucket (bucket_start)
		)
		`,
	}

	for _, statement := range statements {
		if _, err := h.db.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("ensure history schema: %w", err)
		}
	}

	return nil
}

// InsertPlayerPresenceBatch stores one presence row per player UUID for the given bucket.
func (h *HistoryDB) InsertPlayerPresenceBatch(ctx context.Context, bucketStart string, uuids []string) error {
	if len(uuids) == 0 {
		return nil
	}

	var builder strings.Builder
	builder.WriteString("INSERT IGNORE INTO player_hourly_presence (player_uuid, bucket_start) VALUES ")

	args := make([]any, 0, len(uuids)*2)
	for index, uuid := range uuids {
		if index > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString("(?, ?)")
		args = append(args, uuid, bucketStart)
	}

	if _, err := h.db.ExecContext(ctx, builder.String(), args...); err != nil {
		return fmt.Errorf("insert player hourly presence: %w", err)
	}

	return nil
}

// UpsertServerHourlyMax stores the maximum concurrent players seen for the bucket.
func (h *HistoryDB) UpsertServerHourlyMax(ctx context.Context, bucketStart string, maxPlayers int) error {
	query := `
		INSERT INTO server_hourly_presence (bucket_start, max_players)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE max_players = GREATEST(max_players, VALUES(max_players))
	`

	if _, err := h.db.ExecContext(ctx, query, bucketStart, maxPlayers); err != nil {
		return fmt.Errorf("upsert server hourly presence: %w", err)
	}

	return nil
}

// LoadServerCounts returns server counts keyed by bucket timestamp string.
func (h *HistoryDB) LoadServerCounts(ctx context.Context, startBucket string, endBucket string) (map[string]int, error) {
	query := `
		SELECT DATE_FORMAT(bucket_start, '%Y-%m-%d %H:%i:%s') AS bucket_key, max_players
		FROM server_hourly_presence
		WHERE bucket_start >= ? AND bucket_start < ?
	`

	rows, err := h.db.QueryContext(ctx, query, startBucket, endBucket)
	if err != nil {
		return nil, fmt.Errorf("query server hourly presence: %w", err)
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var bucketKey string
		var maxPlayers int
		if err := rows.Scan(&bucketKey, &maxPlayers); err != nil {
			return nil, fmt.Errorf("scan server hourly presence: %w", err)
		}
		result[bucketKey] = maxPlayers
	}

	return result, rows.Err()
}

// LoadPlayerPresence returns player online buckets keyed by timestamp string.
func (h *HistoryDB) LoadPlayerPresence(ctx context.Context, playerUUID string, startBucket string, endBucket string) (map[string]bool, error) {
	query := `
		SELECT DATE_FORMAT(bucket_start, '%Y-%m-%d %H:%i:%s') AS bucket_key
		FROM player_hourly_presence
		WHERE player_uuid = ? AND bucket_start >= ? AND bucket_start < ?
	`

	rows, err := h.db.QueryContext(ctx, query, playerUUID, startBucket, endBucket)
	if err != nil {
		return nil, fmt.Errorf("query player hourly presence: %w", err)
	}
	defer rows.Close()

	result := make(map[string]bool)
	for rows.Next() {
		var bucketKey string
		if err := rows.Scan(&bucketKey); err != nil {
			return nil, fmt.Errorf("scan player hourly presence: %w", err)
		}
		result[bucketKey] = true
	}

	return result, rows.Err()
}

// DeleteOlderThan prunes history older than the provided local bucket start.
func (h *HistoryDB) DeleteOlderThan(ctx context.Context, cutoffBucket string) error {
	statements := []string{
		"DELETE FROM server_hourly_presence WHERE bucket_start < ?",
		"DELETE FROM player_hourly_presence WHERE bucket_start < ?",
	}

	for _, statement := range statements {
		if _, err := h.db.ExecContext(ctx, statement, cutoffBucket); err != nil {
			return fmt.Errorf("delete old history rows: %w", err)
		}
	}

	return nil
}

// Close shuts down the history database connection.
func (h *HistoryDB) Close() error {
	return h.db.Close()
}
