// Package whitelist coordinates Minecraft whitelist operations across API,
// RCON, persistence, and OneBot command sources.
package whitelist

import (
	"context"
	"errors"
	"strings"
)

// Edition identifies the Minecraft edition targeted by a whitelist command.
type Edition string

const (
	// EditionAll is used by list operations to include all editions.
	EditionAll Edition = "all"
	// EditionJava targets Java Edition whitelist commands.
	EditionJava Edition = "java"
	// EditionBedrock targets Bedrock Edition fwhitelist commands.
	EditionBedrock Edition = "bedrock"
)

// Action identifies a whitelist mutation.
type Action string

const (
	// ActionAdd adds a player to the whitelist.
	ActionAdd Action = "add"
	// ActionRemove removes a player from the whitelist.
	ActionRemove Action = "remove"
)

// Source identifies the caller surface that requested a whitelist operation.
type Source string

const (
	// SourceAPI indicates an HTTP API request.
	SourceAPI Source = "api"
	// SourceOneBot indicates a OneBot group command.
	SourceOneBot Source = "onebot"
)

var (
	// ErrUnavailable indicates that persistence or RCON is not configured.
	ErrUnavailable = errors.New("whitelist service unavailable")
	// ErrInvalidInput indicates malformed edition, action, QQ ID, or player name.
	ErrInvalidInput = errors.New("invalid whitelist input")
	// ErrQuotaExceeded indicates that a QQ account reached its active entry limit.
	ErrQuotaExceeded = errors.New("whitelist quota exceeded")
	// ErrConflict indicates that an active entry is owned by another QQ account.
	ErrConflict = errors.New("whitelist entry is owned by another QQ account")
	// ErrForbidden indicates that a caller is not allowed to mutate the entry.
	ErrForbidden = errors.New("whitelist operation forbidden")
	// ErrNotFound indicates that a requested active entry does not exist.
	ErrNotFound = errors.New("whitelist entry not found")
	// ErrRCONRejected indicates that RCON returned an obvious command failure.
	ErrRCONRejected = errors.New("rcon command rejected")
)

// Entry is an active or historical whitelist row mirrored in MySQL.
type Entry struct {
	ID                   int64   `json:"id"`
	QQID                 string  `json:"qq_id,omitempty"`
	Edition              Edition `json:"edition"`
	PlayerName           string  `json:"player_name"`
	NormalizedPlayerName string  `json:"normalized_player_name"`
	Active               bool    `json:"active"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	RemovedAt            *string `json:"removed_at,omitempty"`
}

// AuditRecord captures a whitelist attempt and its RCON result.
type AuditRecord struct {
	Action               Action
	Edition              Edition
	PlayerName           string
	NormalizedPlayerName string
	QQID                 string
	ActorQQID            string
	Source               Source
	Status               string
	RCONOutput           string
	ErrorMessage         string
}

// OperationRequest describes one whitelist mutation.
type OperationRequest struct {
	Action       Action
	Edition      Edition
	PlayerName   string
	QQID         string
	ActorQQID    string
	Source       Source
	Admin        bool
	EnforceQuota bool
}

// OperationResult describes the result of a whitelist mutation.
type OperationResult struct {
	Entry      *Entry       `json:"entry,omitempty"`
	RCONOutput string       `json:"rcon_output,omitempty"`
	Message    string       `json:"message"`
	Changed    bool         `json:"changed"`
	Quota      *QuotaStatus `json:"quota,omitempty"`
}

// QuotaStatus describes a QQ account's current whitelist quota usage.
type QuotaStatus struct {
	Used      int  `json:"used"`
	Limit     int  `json:"limit"`
	Remaining int  `json:"remaining"`
	Exempt    bool `json:"exempt,omitempty"`
}

// Repository is the persistence contract used by the whitelist service.
type Repository interface {
	ListActive(ctx context.Context, edition Edition) ([]Entry, error)
	FindActive(ctx context.Context, edition Edition, normalizedPlayerName string) (*Entry, error)
	CountActiveByQQ(ctx context.Context, qqID string) (int, error)
	UpsertLink(ctx context.Context, qqID string, edition Edition, playerName string, normalizedPlayerName string) error
	UpsertEntry(ctx context.Context, qqID string, edition Edition, playerName string, normalizedPlayerName string) (*Entry, error)
	DeactivateEntry(ctx context.Context, id int64) error
	RecordAudit(ctx context.Context, record AuditRecord) error
}

// Executor runs edition-specific RCON whitelist commands.
type Executor interface {
	Execute(ctx context.Context, edition Edition, action Action, playerName string) (string, error)
}

// NormalizePlayerName returns the canonical key used for quota and ownership checks.
func NormalizePlayerName(playerName string) string {
	return strings.ToLower(strings.TrimSpace(playerName))
}

// MaskQQID returns a privacy-preserving QQ identifier for public responses.
func MaskQQID(qqID string) string {
	trimmed := strings.TrimSpace(qqID)
	if trimmed == "" {
		return ""
	}
	if len(trimmed) <= 4 {
		return strings.Repeat("*", len(trimmed))
	}
	if len(trimmed) <= 7 {
		return trimmed[:2] + "****" + trimmed[len(trimmed)-2:]
	}
	return trimmed[:3] + "****" + trimmed[len(trimmed)-3:]
}

// NormalizeEdition returns a canonical edition.
func NormalizeEdition(value string) (Edition, error) {
	switch Edition(strings.ToLower(strings.TrimSpace(value))) {
	case EditionAll:
		return EditionAll, nil
	case EditionJava:
		return EditionJava, nil
	case EditionBedrock:
		return EditionBedrock, nil
	default:
		return "", ErrInvalidInput
	}
}
