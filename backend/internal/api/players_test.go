package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/whitelist"
)

func TestHandleRandomPlayerFiltersInvalidPlayers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	now := time.Now().UnixMilli()
	store := data.NewStore("")
	store.Players["valid-uuid"] = &data.PlayerInfo{UUID: "valid-uuid", LastSeen: now}
	store.Players["invalid-uuid"] = &data.PlayerInfo{UUID: "invalid-uuid", LastSeen: now}
	store.Stats["valid-uuid"] = &data.PlayerStats{
		UUID:  "valid-uuid",
		Stats: map[string]map[string]int64{"minecraft:custom": {"minecraft:jump": 1}},
	}
	store.Advancements["valid-uuid"] = &data.PlayerAdvancements{
		UUID:         "valid-uuid",
		Advancements: []data.PlayerAdvancement{{Key: "minecraft:story/mine_stone", Done: true}},
	}

	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api/players/random", nil)

	server.handleRandomPlayer(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var response struct {
		UUID string `json:"uuid"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	if response.UUID != "valid-uuid" {
		t.Fatalf("expected valid uuid, got %q", response.UUID)
	}
}

func TestHandleRandomPlayerReturnsNotFoundWithoutValidPlayers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	now := time.Now().UnixMilli()
	store := data.NewStore("")
	store.Players["invalid-uuid"] = &data.PlayerInfo{UUID: "invalid-uuid", LastSeen: now}

	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api/players/random", nil)

	server.handleRandomPlayer(ctx)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestHandlePlayerDetailIncludesMaskedWhitelistAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := validPlayerDetailStore("player-uuid", "Steve", "Java")
	repo := newAPIWhitelistRepo()
	service := whitelist.NewService(repo, &apiWhitelistExecutor{}, 3)
	if _, err := service.Add(t.Context(), whitelist.OperationRequest{
		Action:     whitelist.ActionAdd,
		Edition:    whitelist.EditionJava,
		PlayerName: "Steve",
		QQID:       "1234567890",
		ActorQQID:  "api",
		Source:     whitelist.SourceAPI,
		Admin:      true,
	}); err != nil {
		t.Fatalf("seed whitelist: %v", err)
	}

	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)
	server.SetWhitelist(service)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "uuid", Value: "player-uuid"}}
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api/players/player-uuid", nil)

	server.handlePlayerDetail(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusOK, recorder.Code, recorder.Body.String())
	}

	var response PlayerDetailResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if response.WhitelistAccount == nil {
		t.Fatal("expected whitelist account")
	}
	if response.WhitelistAccount.QQIDMasked != "123****890" {
		t.Fatalf("masked qq = %q, want %q", response.WhitelistAccount.QQIDMasked, "123****890")
	}
}

func TestHandlePlayerDetailOmitsWhitelistAccountWithoutOwner(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := validPlayerDetailStore("player-uuid", "Steve", "Java")
	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)
	server.SetWhitelist(whitelist.NewService(newAPIWhitelistRepo(), &apiWhitelistExecutor{}, 3))

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "uuid", Value: "player-uuid"}}
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api/players/player-uuid", nil)

	server.handlePlayerDetail(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusOK, recorder.Code, recorder.Body.String())
	}

	var response PlayerDetailResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if response.WhitelistAccount != nil {
		t.Fatalf("expected no whitelist account, got %#v", response.WhitelistAccount)
	}
}

func validPlayerDetailStore(uuid string, name string, playerType string) *data.Store {
	now := time.Now().UnixMilli()
	store := data.NewStore("")
	store.Players[uuid] = &data.PlayerInfo{
		UUID:          uuid,
		LastKnownName: name,
		Type:          playerType,
		LastSeen:      now,
		FirstPlayed:   now,
	}
	store.Stats[uuid] = &data.PlayerStats{
		UUID:  uuid,
		Stats: map[string]map[string]int64{"minecraft:custom": {"minecraft:jump": 1}},
	}
	store.Advancements[uuid] = &data.PlayerAdvancements{
		UUID:         uuid,
		Advancements: []data.PlayerAdvancement{{Key: "minecraft:story/mine_stone", Done: true}},
	}
	return store
}
