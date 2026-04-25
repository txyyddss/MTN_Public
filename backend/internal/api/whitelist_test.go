package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/whitelist"
)

func TestWhitelistAPIAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := whitelistTestRouter()

	tests := []struct {
		name   string
		token  string
		status int
	}{
		{name: "missing", status: http.StatusUnauthorized},
		{name: "wrong", token: "wrong", status: http.StatusUnauthorized},
		{name: "correct", token: "secret", status: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/api/whitelist", nil)
			if tt.token != "" {
				request.Header.Set("Authorization", "Bearer "+tt.token)
			}

			router.ServeHTTP(recorder, request)
			if recorder.Code != tt.status {
				t.Fatalf("status = %d, want %d, body=%s", recorder.Code, tt.status, recorder.Body.String())
			}
		})
	}
}

func TestWhitelistAPIAddAndList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := whitelistTestRouter()

	body, err := json.Marshal(map[string]string{
		"edition":     "java",
		"player_name": "Steve",
		"qq_id":       "10001",
	})
	if err != nil {
		t.Fatalf("marshal body: %v", err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/whitelist/add", bytes.NewReader(body))
	request.Header.Set("Authorization", "Bearer secret")
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("add status = %d, body=%s", recorder.Code, recorder.Body.String())
	}

	recorder = httptest.NewRecorder()
	request = httptest.NewRequest(http.MethodGet, "/api/whitelist?edition=java", nil)
	request.Header.Set("Authorization", "Bearer secret")
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("list status = %d, body=%s", recorder.Code, recorder.Body.String())
	}

	var response whitelistListResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal list: %v", err)
	}
	if response.Count != 1 || response.Entries[0].PlayerName != "Steve" || response.Entries[0].QQID != "10001" {
		t.Fatalf("unexpected response: %#v", response)
	}
}

func whitelistTestRouter() http.Handler {
	cfg := &config.Config{
		Server: config.ServerConfig{
			CORSOrigins: []string{"*"},
		},
		Whitelist: config.WhitelistConfig{
			APIToken: "secret",
			MaxPerQQ: 3,
		},
	}
	server := NewServer(cfg, data.NewStore(""), nil, nil, nil, nil, nil)
	server.SetWhitelist(whitelist.NewService(newAPIWhitelistRepo(), &apiWhitelistExecutor{}, 3))
	return server.SetupRouter()
}

type apiWhitelistExecutor struct{}

func (a *apiWhitelistExecutor) Execute(ctx context.Context, edition whitelist.Edition, action whitelist.Action, playerName string) (string, error) {
	return "OK", nil
}

type apiWhitelistRepo struct {
	mu      sync.Mutex
	nextID  int64
	entries map[string]whitelist.Entry
}

func newAPIWhitelistRepo() *apiWhitelistRepo {
	return &apiWhitelistRepo{
		nextID:  1,
		entries: make(map[string]whitelist.Entry),
	}
}

func (a *apiWhitelistRepo) ListActive(ctx context.Context, edition whitelist.Edition) ([]whitelist.Entry, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var entries []whitelist.Entry
	for _, entry := range a.entries {
		if !entry.Active {
			continue
		}
		if edition != whitelist.EditionAll && entry.Edition != edition {
			continue
		}
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i int, j int) bool {
		return entries[i].ID < entries[j].ID
	})
	return entries, nil
}

func (a *apiWhitelistRepo) FindActive(ctx context.Context, edition whitelist.Edition, normalizedPlayerName string) (*whitelist.Entry, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	entry, ok := a.entries[a.key(edition, normalizedPlayerName)]
	if !ok || !entry.Active {
		return nil, nil
	}
	return &entry, nil
}

func (a *apiWhitelistRepo) CountActiveByQQ(ctx context.Context, qqID string) (int, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	count := 0
	for _, entry := range a.entries {
		if entry.Active && entry.QQID == qqID {
			count++
		}
	}
	return count, nil
}

func (a *apiWhitelistRepo) UpsertLink(ctx context.Context, qqID string, edition whitelist.Edition, playerName string, normalizedPlayerName string) error {
	return nil
}

func (a *apiWhitelistRepo) UpsertEntry(ctx context.Context, qqID string, edition whitelist.Edition, playerName string, normalizedPlayerName string) (*whitelist.Entry, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	key := a.key(edition, normalizedPlayerName)
	entry, ok := a.entries[key]
	if !ok {
		entry.ID = a.nextID
		a.nextID++
	}
	entry.QQID = qqID
	entry.Edition = edition
	entry.PlayerName = playerName
	entry.NormalizedPlayerName = normalizedPlayerName
	entry.Active = true
	entry.CreatedAt = "now"
	entry.UpdatedAt = "now"
	a.entries[key] = entry
	return &entry, nil
}

func (a *apiWhitelistRepo) DeactivateEntry(ctx context.Context, id int64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	for key, entry := range a.entries {
		if entry.ID == id {
			entry.Active = false
			a.entries[key] = entry
			return nil
		}
	}
	return whitelist.ErrNotFound
}

func (a *apiWhitelistRepo) RecordAudit(ctx context.Context, record whitelist.AuditRecord) error {
	return nil
}

func (a *apiWhitelistRepo) key(edition whitelist.Edition, normalizedPlayerName string) string {
	return string(edition) + ":" + normalizedPlayerName
}
