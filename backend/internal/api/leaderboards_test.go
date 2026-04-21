package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/database"
)

func TestHandleLeaderboardPaginatesAndFallsBackToFirstPage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	now := time.Now().UnixMilli()
	store := data.NewStore("")
	advancements := []data.PlayerAdvancement{{Key: "minecraft:story/mine_stone", Done: true}}

	for index := 1; index <= 25; index++ {
		uuid := "uuid-" + strconv.Itoa(index)
		name := "Player" + strconv.Itoa(index)

		store.Players[uuid] = &data.PlayerInfo{
			UUID:          uuid,
			LastKnownName: name,
			LastSeen:      now,
		}
		store.Stats[uuid] = &data.PlayerStats{
			UUID: uuid,
			Stats: map[string]map[string]int64{
				"minecraft:mined": {
					"minecraft:stone": int64(100 - index),
				},
			},
		}
		store.Advancements[uuid] = &data.PlayerAdvancements{
			UUID:         uuid,
			Advancements: advancements,
		}
	}

	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)

	secondPage := performLeaderboardRequest(t, server, "mining", "/api/leaderboards/mining?page=2")
	if secondPage.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, secondPage.Code)
	}

	var paged LeaderboardResponse
	if err := json.Unmarshal(secondPage.Body.Bytes(), &paged); err != nil {
		t.Fatalf("unmarshal second page: %v", err)
	}

	if paged.Count != 25 {
		t.Fatalf("expected total count 25, got %d", paged.Count)
	}
	if len(paged.Entries) != 5 {
		t.Fatalf("expected 5 entries on second page, got %d", len(paged.Entries))
	}
	if paged.Entries[0].Rank != 21 {
		t.Fatalf("expected first second-page rank 21, got %d", paged.Entries[0].Rank)
	}

	fallbackPage := performLeaderboardRequest(t, server, "mining", "/api/leaderboards/mining?page=99")
	if fallbackPage.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, fallbackPage.Code)
	}

	var fallback LeaderboardResponse
	if err := json.Unmarshal(fallbackPage.Body.Bytes(), &fallback); err != nil {
		t.Fatalf("unmarshal fallback page: %v", err)
	}

	if len(fallback.Entries) != 20 {
		t.Fatalf("expected fallback page size 20, got %d", len(fallback.Entries))
	}
	if fallback.Entries[0].Rank != 1 {
		t.Fatalf("expected fallback to restart at rank 1, got %d", fallback.Entries[0].Rank)
	}
}

func TestHandleLeaderboardSupportsDynamicStatTypes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	now := time.Now().UnixMilli()
	store := data.NewStore("")
	advancements := []data.PlayerAdvancement{{Key: "minecraft:story/mine_stone", Done: true}}

	store.Players["alpha"] = &data.PlayerInfo{UUID: "alpha", LastKnownName: "Alpha", LastSeen: now}
	store.Players["beta"] = &data.PlayerInfo{UUID: "beta", LastKnownName: "Beta", LastSeen: now}
	store.Players["gamma"] = &data.PlayerInfo{UUID: "gamma", LastKnownName: "Gamma", LastSeen: now}

	for _, uuid := range []string{"alpha", "beta", "gamma"} {
		store.Advancements[uuid] = &data.PlayerAdvancements{
			UUID:         uuid,
			Advancements: advancements,
		}
	}

	store.Stats["alpha"] = &data.PlayerStats{
		UUID: "alpha",
		Stats: map[string]map[string]int64{
			"minecraft:custom": {
				"minecraft:play_time": 300,
			},
		},
	}
	store.Stats["beta"] = &data.PlayerStats{
		UUID: "beta",
		Stats: map[string]map[string]int64{
			"minecraft:custom": {
				"minecraft:play_time": 500,
			},
		},
	}
	store.Stats["gamma"] = &data.PlayerStats{
		UUID: "gamma",
		Stats: map[string]map[string]int64{
			"minecraft:custom": {
				"minecraft:play_time": 200,
			},
		},
	}

	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)

	response := performLeaderboardRequest(t, server, "stat:minecraft:custom:minecraft:play_time", "/api/leaderboards/stat:minecraft:custom:minecraft:play_time")
	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload LeaderboardResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("unmarshal dynamic stat leaderboard: %v", err)
	}

	if payload.Type != "stat:minecraft:custom:minecraft:play_time" {
		t.Fatalf("expected dynamic type to round-trip, got %q", payload.Type)
	}
	if len(payload.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(payload.Entries))
	}
	if payload.Entries[0].UUID != "beta" || payload.Entries[0].Rank != 1 {
		t.Fatalf("expected Beta to lead the stat board, got %+v", payload.Entries[0])
	}
}

func TestHandleLeaderboardSupportsDynamicMcmmoTypes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	now := time.Now().UnixMilli()
	store := data.NewStore("")
	advancements := []data.PlayerAdvancement{{Key: "minecraft:story/mine_stone", Done: true}}

	for _, player := range []struct {
		uuid string
		name string
	}{
		{uuid: "alpha", name: "Alpha"},
		{uuid: "beta", name: "Beta"},
		{uuid: "gamma", name: "Gamma"},
	} {
		store.Players[player.uuid] = &data.PlayerInfo{
			UUID:          player.uuid,
			LastKnownName: player.name,
			LastSeen:      now,
		}
		store.Stats[player.uuid] = &data.PlayerStats{
			UUID: player.uuid,
			Stats: map[string]map[string]int64{
				"minecraft:custom": {
					"minecraft:play_time": 100,
				},
			},
		}
		store.Advancements[player.uuid] = &data.PlayerAdvancements{
			UUID:         player.uuid,
			Advancements: advancements,
		}
	}

	server := NewServer(&config.Config{ActiveDays: 30}, store, nil, nil, nil, nil, nil)
	server.mcmmoSkillsLoader = func(_ context.Context) ([]database.McmmoSkills, error) {
		return []database.McmmoSkills{
			{UUID: "alpha", User: "Alpha", Mining: 45, Total: 45},
			{UUID: "beta", User: "Beta", Mining: 60, Total: 60},
			{UUID: "gamma", User: "Gamma", Mining: 12, Total: 12},
		}, nil
	}

	response := performLeaderboardRequest(t, server, "mcmmo:mining", "/api/leaderboards/mcmmo:mining")
	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload LeaderboardResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("unmarshal dynamic mcmmo leaderboard: %v", err)
	}

	if len(payload.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(payload.Entries))
	}
	if payload.Entries[0].UUID != "beta" || payload.Entries[0].Value != 60 {
		t.Fatalf("expected Beta to lead the mining board, got %+v", payload.Entries[0])
	}
}

func TestHandleLeaderboardRejectsInvalidDynamicType(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := NewServer(&config.Config{ActiveDays: 30}, data.NewStore(""), nil, nil, nil, nil, nil)
	response := performLeaderboardRequest(t, server, "mcmmo:invalid", "/api/leaderboards/mcmmo:invalid")

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func performLeaderboardRequest(t *testing.T, server *Server, leaderboardType, url string) *httptest.ResponseRecorder {
	t.Helper()

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "type", Value: leaderboardType}}
	ctx.Request = httptest.NewRequest(http.MethodGet, url, nil)

	server.handleLeaderboard(ctx)

	return recorder
}
