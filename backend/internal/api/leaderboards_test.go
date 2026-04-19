package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/data"
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

	secondPage := performLeaderboardRequest(t, server, "/api/leaderboards/mining?page=2")
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

	fallbackPage := performLeaderboardRequest(t, server, "/api/leaderboards/mining?page=99")
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

func performLeaderboardRequest(t *testing.T, server *Server, url string) *httptest.ResponseRecorder {
	t.Helper()

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "type", Value: "mining"}}
	ctx.Request = httptest.NewRequest(http.MethodGet, url, nil)

	server.handleLeaderboard(ctx)

	return recorder
}
