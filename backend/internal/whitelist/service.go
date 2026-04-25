package whitelist

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// Service applies whitelist rules and coordinates persistence with RCON.
type Service struct {
	repo     Repository
	executor Executor
	maxPerQQ int
}

// NewService creates a whitelist service.
func NewService(repo Repository, executor Executor, maxPerQQ int) *Service {
	if maxPerQQ <= 0 {
		maxPerQQ = 3
	}

	return &Service{
		repo:     repo,
		executor: executor,
		maxPerQQ: maxPerQQ,
	}
}

// List returns active whitelist entries from the MySQL mirror.
func (s *Service) List(ctx context.Context, edition Edition) ([]Entry, error) {
	if s == nil || s.repo == nil {
		return nil, ErrUnavailable
	}
	if edition != EditionAll && edition != EditionJava && edition != EditionBedrock {
		return nil, ErrInvalidInput
	}

	entries, err := s.repo.ListActive(ctx, edition)
	if err != nil {
		return nil, fmt.Errorf("list active whitelist entries: %w", err)
	}

	return entries, nil
}

// FindActive returns one active whitelist entry by edition and normalized player name.
func (s *Service) FindActive(ctx context.Context, edition Edition, normalizedPlayerName string) (*Entry, error) {
	if s == nil || s.repo == nil {
		return nil, ErrUnavailable
	}
	if edition != EditionJava && edition != EditionBedrock {
		return nil, ErrInvalidInput
	}
	if strings.TrimSpace(normalizedPlayerName) == "" {
		return nil, ErrInvalidInput
	}

	entry, err := s.repo.FindActive(ctx, edition, strings.TrimSpace(normalizedPlayerName))
	if err != nil {
		return nil, fmt.Errorf("find active whitelist entry: %w", err)
	}

	return entry, nil
}

// Quota returns current quota usage for a QQ account.
func (s *Service) Quota(ctx context.Context, qqID string, exempt bool) (*QuotaStatus, error) {
	return s.quota(ctx, qqID, exempt)
}

// Add adds or reactivates a whitelist entry after a successful RCON command.
func (s *Service) Add(ctx context.Context, req OperationRequest) (*OperationResult, error) {
	if err := s.validateMutation(req); err != nil {
		return nil, err
	}

	normalized := NormalizePlayerName(req.PlayerName)
	existing, err := s.repo.FindActive(ctx, req.Edition, normalized)
	if err != nil {
		return nil, fmt.Errorf("find active whitelist entry: %w", err)
	}
	if existing != nil {
		if req.QQID == "" || existing.QQID == req.QQID {
			if req.QQID != "" {
				if err := s.repo.UpsertLink(ctx, req.QQID, req.Edition, req.PlayerName, normalized); err != nil {
					return nil, fmt.Errorf("upsert account link: %w", err)
				}
			}
			quota, err := s.quota(ctx, req.QQID, !req.EnforceQuota)
			if err != nil {
				return nil, err
			}
			return &OperationResult{
				Entry:   existing,
				Message: "already whitelisted",
				Changed: false,
				Quota:   quota,
			}, nil
		}

		return nil, ErrConflict
	}

	if req.EnforceQuota && req.QQID != "" {
		count, err := s.repo.CountActiveByQQ(ctx, req.QQID)
		if err != nil {
			return nil, fmt.Errorf("count active whitelist entries: %w", err)
		}
		if count >= s.maxPerQQ {
			return nil, ErrQuotaExceeded
		}
	}

	output, err := s.executeAndAudit(ctx, req, normalized)
	if err != nil {
		return nil, err
	}

	if req.QQID != "" {
		if err := s.repo.UpsertLink(ctx, req.QQID, req.Edition, req.PlayerName, normalized); err != nil {
			return nil, fmt.Errorf("upsert account link: %w", err)
		}
	}

	entry, err := s.repo.UpsertEntry(ctx, req.QQID, req.Edition, req.PlayerName, normalized)
	if err != nil {
		return nil, fmt.Errorf("upsert whitelist entry: %w", err)
	}
	quota, err := s.quota(ctx, req.QQID, !req.EnforceQuota)
	if err != nil {
		return nil, err
	}

	return &OperationResult{
		Entry:      entry,
		RCONOutput: output,
		Message:    "whitelist added",
		Changed:    true,
		Quota:      quota,
	}, nil
}

// Remove removes a whitelist entry after a successful RCON command.
func (s *Service) Remove(ctx context.Context, req OperationRequest) (*OperationResult, error) {
	if err := s.validateMutation(req); err != nil {
		return nil, err
	}

	normalized := NormalizePlayerName(req.PlayerName)
	existing, err := s.repo.FindActive(ctx, req.Edition, normalized)
	if err != nil {
		return nil, fmt.Errorf("find active whitelist entry: %w", err)
	}
	if existing == nil && !req.Admin {
		return nil, ErrNotFound
	}
	if existing != nil && !req.Admin && existing.QQID != req.QQID {
		return nil, ErrForbidden
	}

	output, err := s.executeAndAudit(ctx, req, normalized)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		quota, err := s.quota(ctx, req.QQID, !req.EnforceQuota)
		if err != nil {
			return nil, err
		}
		return &OperationResult{
			RCONOutput: output,
			Message:    "whitelist remove command sent",
			Changed:    false,
			Quota:      quota,
		}, nil
	}

	if err := s.repo.DeactivateEntry(ctx, existing.ID); err != nil {
		return nil, fmt.Errorf("deactivate whitelist entry: %w", err)
	}
	entry := *existing
	entry.Active = false
	quota, err := s.quota(ctx, req.QQID, !req.EnforceQuota)
	if err != nil {
		return nil, err
	}

	return &OperationResult{
		Entry:      &entry,
		RCONOutput: output,
		Message:    "whitelist removed",
		Changed:    true,
		Quota:      quota,
	}, nil
}

func (s *Service) quota(ctx context.Context, qqID string, exempt bool) (*QuotaStatus, error) {
	if strings.TrimSpace(qqID) == "" {
		return nil, nil
	}
	if s == nil || s.repo == nil {
		return nil, ErrUnavailable
	}

	count, err := s.repo.CountActiveByQQ(ctx, qqID)
	if err != nil {
		return nil, fmt.Errorf("count active whitelist entries: %w", err)
	}
	remaining := s.maxPerQQ - count
	if remaining < 0 {
		remaining = 0
	}

	return &QuotaStatus{
		Used:      count,
		Limit:     s.maxPerQQ,
		Remaining: remaining,
		Exempt:    exempt,
	}, nil
}

func (s *Service) validateMutation(req OperationRequest) error {
	if s == nil || s.repo == nil || s.executor == nil {
		return ErrUnavailable
	}
	if req.Action != ActionAdd && req.Action != ActionRemove {
		return ErrInvalidInput
	}
	if req.Edition != EditionJava && req.Edition != EditionBedrock {
		return ErrInvalidInput
	}
	if req.Source != SourceAPI && req.Source != SourceOneBot {
		return ErrInvalidInput
	}
	if err := ValidatePlayerName(req.PlayerName); err != nil {
		return err
	}
	if strings.TrimSpace(req.QQID) != req.QQID || strings.TrimSpace(req.ActorQQID) != req.ActorQQID {
		return ErrInvalidInput
	}
	return nil
}

func (s *Service) executeAndAudit(ctx context.Context, req OperationRequest, normalized string) (string, error) {
	output, err := s.executor.Execute(ctx, req.Edition, req.Action, req.PlayerName)
	if err == nil {
		err = rejectObviousRCONError(output)
	}

	record := AuditRecord{
		Action:               req.Action,
		Edition:              req.Edition,
		PlayerName:           req.PlayerName,
		NormalizedPlayerName: normalized,
		QQID:                 req.QQID,
		ActorQQID:            req.ActorQQID,
		Source:               req.Source,
		Status:               "success",
		RCONOutput:           output,
	}
	if err != nil {
		record.Status = "failed"
		record.ErrorMessage = err.Error()
	}

	if auditErr := s.repo.RecordAudit(ctx, record); auditErr != nil && err == nil {
		return "", fmt.Errorf("record whitelist audit: %w", auditErr)
	}

	if err != nil {
		return output, err
	}

	return output, nil
}

func rejectObviousRCONError(output string) error {
	normalized := strings.ToLower(output)
	errorMarkers := []string{
		"unknown command",
		"incorrect argument",
		"invalid argument",
		"syntax error",
		"usage:",
		"exception",
		"failed",
		"error:",
	}

	for _, marker := range errorMarkers {
		if strings.Contains(normalized, marker) {
			return fmt.Errorf("%w: %s", ErrRCONRejected, strings.TrimSpace(output))
		}
	}

	return nil
}

// ErrorCode returns a stable short code for service errors.
func ErrorCode(err error) string {
	switch {
	case errors.Is(err, ErrUnavailable):
		return "unavailable"
	case errors.Is(err, ErrInvalidInput):
		return "invalid_input"
	case errors.Is(err, ErrQuotaExceeded):
		return "quota_exceeded"
	case errors.Is(err, ErrConflict):
		return "conflict"
	case errors.Is(err, ErrForbidden):
		return "forbidden"
	case errors.Is(err, ErrNotFound):
		return "not_found"
	case errors.Is(err, ErrRCONRejected):
		return "rcon_rejected"
	default:
		return "internal_error"
	}
}
