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
