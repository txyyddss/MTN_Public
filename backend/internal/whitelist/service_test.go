package whitelist

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"testing"
)

func TestParseCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		raw     string
		matched bool
		wantErr bool
		edition Edition
		action  Action
		player  string
	}{
		{
			name:    "java bind",
			raw:     "bind java Steve",
			matched: true,
			edition: EditionJava,
			action:  ActionAdd,
			player:  "Steve",
		},
		{
			name:    "bedrock unbind",
			raw:     "unbind bedrock Bedrock_User",
			matched: true,
			edition: EditionBedrock,
			action:  ActionRemove,
			player:  "Bedrock_User",
		},
		{
			name:    "leading at mention",
			raw:     "[CQ:at,qq=123456] bind java Steve",
			matched: true,
			edition: EditionJava,
			action:  ActionAdd,
			player:  "Steve",
		},
		{
			name:    "slash command",
			raw:     "/bind bedrock Bedrock_User",
			matched: true,
			edition: EditionBedrock,
			action:  ActionAdd,
			player:  "Bedrock_User",
		},
		{
			name:    "legacy java command invalid",
			raw:     "whitelist add Steve",
			matched: true,
			wantErr: true,
		},
		{
			name:    "legacy bedrock command invalid",
			raw:     "fwhitelist remove Bedrock_User",
			matched: true,
			wantErr: true,
		},
		{
			name:    "not a command",
			raw:     "hello bind java Steve",
			matched: false,
		},
		{
			name:    "missing player",
			raw:     "bind java",
			matched: true,
			wantErr: true,
		},
		{
			name:    "unsafe player token",
			raw:     "bind java Steve;op",
			matched: true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			command, matched, err := ParseCommand(tt.raw)
			if matched != tt.matched {
				t.Fatalf("matched = %v, want %v", matched, tt.matched)
			}
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.matched {
				return
			}
			if command.Edition != tt.edition || command.Action != tt.action || command.PlayerName != tt.player {
				t.Fatalf("unexpected command: %#v", command)
			}
		})
	}
}

func TestParseOneBotGroupEventUsesMessageFallback(t *testing.T) {
	t.Parallel()

	data := []byte(`{
		"post_type": "message",
		"message_type": "group",
		"group_id": 10001,
		"user_id": 20002,
		"message": "bind java Steve",
		"sender": {"role": "member"}
	}`)

	event, ok := parseOneBotGroupEvent(data)
	if !ok {
		t.Fatal("expected event")
	}
	if event.RawMessage != "bind java Steve" {
		t.Fatalf("raw message = %q, want %q", event.RawMessage, "bind java Steve")
	}
	if string(event.GroupID) != "10001" || string(event.UserID) != "20002" {
		t.Fatalf("unexpected ids: group=%q user=%q", event.GroupID, event.UserID)
	}
}

func TestParseOneBotGroupEventUsesSegmentTextFallback(t *testing.T) {
	t.Parallel()

	data := []byte(`{
		"post_type": "message",
		"message_type": "group",
		"group_id": "10001",
		"user_id": "20002",
		"message": [
			{"type": "at", "data": {"qq": "123456"}},
			{"type": "text", "data": {"text": " bind bedrock Bedrock_User"}}
		],
		"sender": {"role": "admin"}
	}`)

	event, ok := parseOneBotGroupEvent(data)
	if !ok {
		t.Fatal("expected event")
	}
	command, matched, err := ParseCommand(event.RawMessage)
	if err != nil {
		t.Fatalf("parse command: %v", err)
	}
	if !matched {
		t.Fatal("expected command match")
	}
	if command.Action != ActionAdd || command.Edition != EditionBedrock || command.PlayerName != "Bedrock_User" {
		t.Fatalf("unexpected command: %#v", command)
	}
}

func TestOneBotGroupMessagePostTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		postType string
		want     bool
	}{
		{postType: "message", want: true},
		{postType: "message_sent", want: true},
		{postType: "notice", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.postType, func(t *testing.T) {
			t.Parallel()

			if got := isOneBotGroupMessagePost(tt.postType); got != tt.want {
				t.Fatalf("isOneBotGroupMessagePost(%q) = %v, want %v", tt.postType, got, tt.want)
			}
		})
	}
}

func TestParseOneBotMessageSentGroupCommand(t *testing.T) {
	t.Parallel()

	data := []byte(`{
		"self_id": 2731818648,
		"user_id": 2731818648,
		"post_type": "message_sent",
		"message_sent_type": "self",
		"message_type": "group",
		"group_id": 1064494318,
		"raw_message": "bind java txro",
		"message": [
			{"type": "text", "data": {"text": "bind java txro"}}
		],
		"sender": {"user_id": 2731818648, "nickname": "TX", "card": "TX", "role": "admin"}
	}`)

	event, ok := parseOneBotGroupEvent(data)
	if !ok {
		t.Fatal("expected event")
	}
	if !isOneBotGroupMessagePost(event.PostType) {
		t.Fatalf("unexpected post type: %q", event.PostType)
	}
	command, matched, err := ParseCommand(event.RawMessage)
	if err != nil {
		t.Fatalf("parse command: %v", err)
	}
	if !matched {
		t.Fatal("expected command match")
	}
	if string(event.GroupID) != "1064494318" || string(event.UserID) != "2731818648" {
		t.Fatalf("unexpected ids: group=%q user=%q", event.GroupID, event.UserID)
	}
	if command.Action != ActionAdd || command.Edition != EditionJava || command.PlayerName != "txro" {
		t.Fatalf("unexpected command: %#v", command)
	}
}

func TestMaskQQID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		qqID string
		want string
	}{
		{qqID: "", want: ""},
		{qqID: "1234", want: "****"},
		{qqID: "123456", want: "12****56"},
		{qqID: "1234567890", want: "123****890"},
	}

	for _, tt := range tests {
		t.Run(tt.qqID, func(t *testing.T) {
			t.Parallel()

			if got := MaskQQID(tt.qqID); got != tt.want {
				t.Fatalf("MaskQQID(%q) = %q, want %q", tt.qqID, got, tt.want)
			}
		})
	}
}

func TestFormatOneBotResultIncludesQuota(t *testing.T) {
	t.Parallel()

	command := &Command{Action: ActionAdd, Edition: EditionJava, PlayerName: "Steve"}
	quota := &QuotaStatus{Used: 2, Limit: 3, Remaining: 1}
	message := formatOneBotResult(command, &OperationResult{Message: "whitelist added", Changed: true, Quota: quota}, nil, nil)

	if !strings.Contains(message, "Bound Steve to the java whitelist.") {
		t.Fatalf("missing success detail: %q", message)
	}
	if !strings.Contains(message, "Quota: 2/3 used, 1 remaining.") {
		t.Fatalf("missing quota detail: %q", message)
	}
}

func TestFormatOneBotResultIncludesFailureQuota(t *testing.T) {
	t.Parallel()

	command := &Command{Action: ActionAdd, Edition: EditionBedrock, PlayerName: "Alex"}
	quota := &QuotaStatus{Used: 3, Limit: 3, Remaining: 0}
	message := formatOneBotResult(command, nil, ErrQuotaExceeded, quota)

	if !strings.Contains(message, "Bind failed for Alex (bedrock): quota reached") {
		t.Fatalf("missing failure detail: %q", message)
	}
	if !strings.Contains(message, "Quota: 3/3 used, 0 remaining.") {
		t.Fatalf("missing quota detail: %q", message)
	}
}

func TestFormatOneBotResultIncludesForbiddenAndNotFoundQuota(t *testing.T) {
	t.Parallel()

	quota := &QuotaStatus{Used: 1, Limit: 3, Remaining: 2}
	tests := []struct {
		name    string
		command *Command
		err     error
		want    string
	}{
		{
			name:    "forbidden unbind",
			command: &Command{Action: ActionRemove, Edition: EditionJava, PlayerName: "Steve"},
			err:     ErrForbidden,
			want:    "Unbind failed for Steve (java): you can only remove whitelist entries linked to your QQ account.",
		},
		{
			name:    "not found unbind",
			command: &Command{Action: ActionRemove, Edition: EditionBedrock, PlayerName: "Alex"},
			err:     ErrNotFound,
			want:    "Unbind failed for Alex (bedrock): no active whitelist entry was found for this player.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			message := formatOneBotResult(tt.command, nil, tt.err, quota)
			if !strings.Contains(message, tt.want) {
				t.Fatalf("missing failure detail: %q", message)
			}
			if !strings.Contains(message, "Quota: 1/3 used, 2 remaining.") {
				t.Fatalf("missing quota detail: %q", message)
			}
		})
	}
}

func TestFormatOneBotResultDoesNotDisplayExemptQuota(t *testing.T) {
	t.Parallel()

	command := &Command{Action: ActionRemove, Edition: EditionJava, PlayerName: "Steve"}
	quota := &QuotaStatus{Used: 1, Limit: 3, Remaining: 2, Exempt: true}
	message := formatOneBotResult(command, &OperationResult{Message: "whitelist removed", Changed: true, Quota: quota}, nil, nil)

	if !strings.Contains(message, "Unbound Steve from the java whitelist.") {
		t.Fatalf("missing success detail: %q", message)
	}
	if strings.Contains(message, "admin exempt") {
		t.Fatalf("unexpected admin quota detail: %q", message)
	}
	if !strings.Contains(message, "Quota: 1/3 used, 2 remaining.") {
		t.Fatalf("missing quota detail: %q", message)
	}
}

func TestFormatOneBotResultUsesQuotaNumbers(t *testing.T) {
	t.Parallel()

	command := &Command{Action: ActionAdd, Edition: EditionJava, PlayerName: "Steve"}
	quota := &QuotaStatus{Used: 2, Limit: 3, Remaining: 1, Exempt: false}
	message := formatOneBotResult(command, &OperationResult{Message: "whitelist added", Changed: true, Quota: quota}, nil, nil)

	if strings.Contains(message, "admin exempt") {
		t.Fatalf("unexpected admin quota detail: %q", message)
	}
	if !strings.Contains(message, "Quota: 2/3 used, 1 remaining.") {
		t.Fatalf("missing quota detail: %q", message)
	}
}

func TestServiceReturnsQuotaForBindDuplicateAndUnbind(t *testing.T) {
	t.Parallel()

	repo := newFakeRepository()
	service := NewService(repo, &fakeExecutor{}, 3)
	ctx := context.Background()
	req := OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionJava,
		PlayerName:   "Steve",
		QQID:         "10001",
		ActorQQID:    "10001",
		Source:       SourceOneBot,
		EnforceQuota: true,
	}

	added, err := service.Add(ctx, req)
	if err != nil {
		t.Fatalf("add: %v", err)
	}
	assertQuota(t, added.Quota, 1, 3, 2)

	duplicate, err := service.Add(ctx, req)
	if err != nil {
		t.Fatalf("duplicate add: %v", err)
	}
	if duplicate.Changed {
		t.Fatal("duplicate add should be idempotent")
	}
	assertQuota(t, duplicate.Quota, 1, 3, 2)

	removed, err := service.Remove(ctx, OperationRequest{
		Action:     ActionRemove,
		Edition:    EditionJava,
		PlayerName: "Steve",
		QQID:       "10001",
		ActorQQID:  "10001",
		Source:     SourceOneBot,
	})
	if err != nil {
		t.Fatalf("remove: %v", err)
	}
	assertQuota(t, removed.Quota, 0, 3, 3)
}

func TestServiceEnforcesQuota(t *testing.T) {
	t.Parallel()

	repo := newFakeRepository()
	executor := &fakeExecutor{}
	service := NewService(repo, executor, 3)
	ctx := context.Background()

	for index := 1; index <= 3; index++ {
		_, err := service.Add(ctx, OperationRequest{
			Action:       ActionAdd,
			Edition:      EditionJava,
			PlayerName:   fmt.Sprintf("Player%d", index),
			QQID:         "10001",
			ActorQQID:    "10001",
			Source:       SourceOneBot,
			EnforceQuota: true,
		})
		if err != nil {
			t.Fatalf("add player %d: %v", index, err)
		}
	}

	_, err := service.Add(ctx, OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionBedrock,
		PlayerName:   "Player4",
		QQID:         "10001",
		ActorQQID:    "10001",
		Source:       SourceOneBot,
		EnforceQuota: true,
	})
	if !errors.Is(err, ErrQuotaExceeded) {
		t.Fatalf("expected quota error, got %v", err)
	}
}

func TestServiceDuplicateSameOwnerIsIdempotent(t *testing.T) {
	t.Parallel()

	repo := newFakeRepository()
	executor := &fakeExecutor{}
	service := NewService(repo, executor, 3)
	ctx := context.Background()
	req := OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionJava,
		PlayerName:   "Steve",
		QQID:         "10001",
		ActorQQID:    "10001",
		Source:       SourceOneBot,
		EnforceQuota: true,
	}

	first, err := service.Add(ctx, req)
	if err != nil {
		t.Fatalf("first add: %v", err)
	}
	second, err := service.Add(ctx, req)
	if err != nil {
		t.Fatalf("second add: %v", err)
	}
	if second.Changed {
		t.Fatal("duplicate add should be idempotent")
	}
	if first.Entry == nil || second.Entry == nil || first.Entry.ID != second.Entry.ID {
		t.Fatalf("expected same entry, got %#v then %#v", first.Entry, second.Entry)
	}
	if executor.Count() != 1 {
		t.Fatalf("expected one RCON execution, got %d", executor.Count())
	}
}

func TestServiceConflictsWithOtherOwner(t *testing.T) {
	t.Parallel()

	repo := newFakeRepository()
	service := NewService(repo, &fakeExecutor{}, 3)
	ctx := context.Background()

	_, err := service.Add(ctx, OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionJava,
		PlayerName:   "Steve",
		QQID:         "10001",
		ActorQQID:    "10001",
		Source:       SourceOneBot,
		EnforceQuota: true,
	})
	if err != nil {
		t.Fatalf("seed add: %v", err)
	}

	_, err = service.Add(ctx, OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionJava,
		PlayerName:   "Steve",
		QQID:         "20002",
		ActorQQID:    "20002",
		Source:       SourceOneBot,
		EnforceQuota: true,
	})
	if !errors.Is(err, ErrConflict) {
		t.Fatalf("expected conflict, got %v", err)
	}
}

func TestServiceRemoveOwnershipRules(t *testing.T) {
	t.Parallel()

	repo := newFakeRepository()
	service := NewService(repo, &fakeExecutor{}, 3)
	ctx := context.Background()

	_, err := service.Add(ctx, OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionJava,
		PlayerName:   "Steve",
		QQID:         "10001",
		ActorQQID:    "10001",
		Source:       SourceOneBot,
		EnforceQuota: true,
	})
	if err != nil {
		t.Fatalf("seed add: %v", err)
	}

	_, err = service.Remove(ctx, OperationRequest{
		Action:     ActionRemove,
		Edition:    EditionJava,
		PlayerName: "Steve",
		QQID:       "20002",
		ActorQQID:  "20002",
		Source:     SourceOneBot,
	})
	if !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected forbidden, got %v", err)
	}

	result, err := service.Remove(ctx, OperationRequest{
		Action:     ActionRemove,
		Edition:    EditionJava,
		PlayerName: "Steve",
		QQID:       "20002",
		ActorQQID:  "20002",
		Source:     SourceOneBot,
		Admin:      true,
	})
	if err != nil {
		t.Fatalf("admin remove: %v", err)
	}
	if !result.Changed {
		t.Fatal("expected admin remove to change state")
	}
}

func TestServiceRCONRejectionDoesNotMutateMirror(t *testing.T) {
	t.Parallel()

	repo := newFakeRepository()
	service := NewService(repo, &fakeExecutor{output: "Unknown command"}, 3)
	ctx := context.Background()

	_, err := service.Add(ctx, OperationRequest{
		Action:       ActionAdd,
		Edition:      EditionJava,
		PlayerName:   "Steve",
		QQID:         "10001",
		ActorQQID:    "10001",
		Source:       SourceOneBot,
		EnforceQuota: true,
	})
	if !errors.Is(err, ErrRCONRejected) {
		t.Fatalf("expected RCON rejection, got %v", err)
	}
	entries, err := repo.ListActive(ctx, EditionAll)
	if err != nil {
		t.Fatalf("list active: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("expected no active entries, got %d", len(entries))
	}
	if len(repo.audits) != 1 || repo.audits[0].Status != "failed" {
		t.Fatalf("expected failed audit, got %#v", repo.audits)
	}
}

type fakeExecutor struct {
	mu     sync.Mutex
	output string
	err    error
	count  int
}

func (f *fakeExecutor) Execute(ctx context.Context, edition Edition, action Action, playerName string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.count++
	if f.output == "" {
		f.output = "OK"
	}
	return f.output, f.err
}

func (f *fakeExecutor) Count() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.count
}

type fakeRepository struct {
	mu      sync.Mutex
	nextID  int64
	entries map[string]Entry
	audits  []AuditRecord
}

func newFakeRepository() *fakeRepository {
	return &fakeRepository{
		nextID:  1,
		entries: make(map[string]Entry),
	}
}

func (f *fakeRepository) ListActive(ctx context.Context, edition Edition) ([]Entry, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	var entries []Entry
	for _, entry := range f.entries {
		if !entry.Active {
			continue
		}
		if edition != EditionAll && entry.Edition != edition {
			continue
		}
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i int, j int) bool {
		return entries[i].ID < entries[j].ID
	})
	return entries, nil
}

func (f *fakeRepository) FindActive(ctx context.Context, edition Edition, normalizedPlayerName string) (*Entry, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	entry, ok := f.entries[f.key(edition, normalizedPlayerName)]
	if !ok || !entry.Active {
		return nil, nil
	}
	return &entry, nil
}

func (f *fakeRepository) CountActiveByQQ(ctx context.Context, qqID string) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	count := 0
	for _, entry := range f.entries {
		if entry.Active && entry.QQID == qqID {
			count++
		}
	}
	return count, nil
}

func (f *fakeRepository) UpsertLink(ctx context.Context, qqID string, edition Edition, playerName string, normalizedPlayerName string) error {
	return nil
}

func (f *fakeRepository) UpsertEntry(ctx context.Context, qqID string, edition Edition, playerName string, normalizedPlayerName string) (*Entry, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	key := f.key(edition, normalizedPlayerName)
	entry, ok := f.entries[key]
	if !ok {
		entry.ID = f.nextID
		f.nextID++
		entry.CreatedAt = "now"
	}
	entry.QQID = qqID
	entry.Edition = edition
	entry.PlayerName = playerName
	entry.NormalizedPlayerName = normalizedPlayerName
	entry.Active = true
	entry.UpdatedAt = "now"
	entry.RemovedAt = nil
	f.entries[key] = entry
	return &entry, nil
}

func (f *fakeRepository) DeactivateEntry(ctx context.Context, id int64) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	for key, entry := range f.entries {
		if entry.ID == id {
			entry.Active = false
			removed := "now"
			entry.RemovedAt = &removed
			f.entries[key] = entry
			return nil
		}
	}
	return ErrNotFound
}

func (f *fakeRepository) RecordAudit(ctx context.Context, record AuditRecord) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.audits = append(f.audits, record)
	return nil
}

func (f *fakeRepository) key(edition Edition, normalizedPlayerName string) string {
	return string(edition) + ":" + normalizedPlayerName
}

func assertQuota(t *testing.T, quota *QuotaStatus, used int, limit int, remaining int) {
	t.Helper()

	if quota == nil {
		t.Fatal("expected quota")
	}
	if quota.Used != used || quota.Limit != limit || quota.Remaining != remaining {
		t.Fatalf("quota = %#v, want used=%d limit=%d remaining=%d", quota, used, limit, remaining)
	}
}
