package api

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// LeaderboardEntry holds a single leaderboard row.
type LeaderboardEntry struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value int64  `json:"value"`
	Rank  int    `json:"rank"`
}

// LeaderboardResponse holds a leaderboard response.
type LeaderboardResponse struct {
	Type    string             `json:"type"`
	Entries []LeaderboardEntry `json:"entries"`
	Count   int                `json:"count"`
}

// handleLeaderboard returns leaderboard data by type.
func (s *Server) handleLeaderboard(c *gin.Context) {
	lbType := c.Param("type")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize := 20

	var entries []LeaderboardEntry
	cacheKey := "lb_" + lbType

	// Try checking cache first
	if s.cache != nil {
		if ok, _ := s.cache.Get(c.Request.Context(), cacheKey, &entries); ok {
			s.sendPaginatedLeaderboard(c, lbType, entries, page, pageSize)
			return
		}
	}

	switch lbType {
	case "skills":
		entries = s.skillsLeaderboard(c.Request.Context())
	case "playtime":
		entries = s.statLeaderboard(c.Request.Context(), "minecraft:custom", "minecraft:play_time")
	case "mining":
		entries = s.minedTotalLeaderboard(c.Request.Context())
	case "killing":
		entries = s.statLeaderboard(c.Request.Context(), "minecraft:custom", "minecraft:mob_kills")
	case "deaths":
		entries = s.statLeaderboard(c.Request.Context(), "minecraft:custom", "minecraft:deaths")
	case "walking":
		entries = s.statLeaderboard(c.Request.Context(), "minecraft:custom", "minecraft:walk_one_cm")
	case "pvp":
		entries = s.statLeaderboard(c.Request.Context(), "minecraft:custom", "minecraft:player_kills")
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid leaderboard type"})
		return
	}

	if s.cache != nil && len(entries) > 0 {
		s.cache.Set(c.Request.Context(), cacheKey, entries, 5*time.Minute)
	}

	s.sendPaginatedLeaderboard(c, lbType, entries, page, pageSize)
}

func (s *Server) sendPaginatedLeaderboard(c *gin.Context, lbType string, entries []LeaderboardEntry, page, pageSize int) {
	start := (page - 1) * pageSize
	if start >= len(entries) {
		start = 0
		page = 1
	}
	end := start + pageSize
	if end > len(entries) {
		end = len(entries)
	}

	c.JSON(http.StatusOK, LeaderboardResponse{
		Type:    lbType,
		Entries: entries[start:end],
		Count:   len(entries),
	})
}

func (s *Server) skillsLeaderboard(ctx context.Context) []LeaderboardEntry {
	if s.mcmmoDB == nil {
		return nil
	}

	skills, err := s.mcmmoDB.GetAllSkills(ctx)
	if err != nil {
		return nil
	}

	var entries []LeaderboardEntry
	for _, sk := range skills {
		if sk.Total <= 0 {
			continue
		}
		if !s.isValidPlayer(sk.UUID) {
			continue
		}
		name := sk.User
		// Try to get the display name from playerdata
		if info := s.store.GetPlayer(sk.UUID); info != nil {
			name = info.LastKnownName
		}
		entries = append(entries, LeaderboardEntry{
			UUID:  sk.UUID,
			Name:  name,
			Value: int64(sk.Total),
		})
	}

	return rankEntries(entries)
}

func (s *Server) statLeaderboard(ctx context.Context, category, stat string) []LeaderboardEntry {
	allStats := s.store.GetAllStats()
	var entries []LeaderboardEntry
	for uuid, ps := range allStats {
		if catMap, ok := ps.Stats[category]; ok {
			if val, ok := catMap[stat]; ok && val > 0 {
				if !s.isValidPlayer(uuid) {
					continue
				}
				name := uuid
				if info := s.store.GetPlayer(uuid); info != nil {
					name = info.LastKnownName
				}
				entries = append(entries, LeaderboardEntry{
					UUID:  uuid,
					Name:  name,
					Value: val,
				})
			}
		}
	}

	return rankEntries(entries)
}

func (s *Server) minedTotalLeaderboard(ctx context.Context) []LeaderboardEntry {
	allStats := s.store.GetAllStats()
	var entries []LeaderboardEntry
	for uuid, ps := range allStats {
		if minedMap, ok := ps.Stats["minecraft:mined"]; ok {
			var total int64
			for _, count := range minedMap {
				total += count
			}
			if total > 0 {
				if !s.isValidPlayer(uuid) {
					continue
				}
				name := uuid
				if info := s.store.GetPlayer(uuid); info != nil {
					name = info.LastKnownName
				}
				entries = append(entries, LeaderboardEntry{
					UUID:  uuid,
					Name:  name,
					Value: total,
				})
			}
		}
	}

	return rankEntries(entries)
}

func rankEntries(entries []LeaderboardEntry) []LeaderboardEntry {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Value > entries[j].Value
	})

	for i := range entries {
		entries[i].Rank = i + 1
	}

	return entries
}
