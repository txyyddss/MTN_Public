package api

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mtn-server/backend/internal/database"
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

var mcmmoSkillNames = []string{
	"taming",
	"mining",
	"woodcutting",
	"repair",
	"unarmed",
	"herbalism",
	"excavation",
	"archery",
	"swords",
	"axes",
	"acrobatics",
	"fishing",
	"alchemy",
	"crossbows",
	"tridents",
	"maces",
	"spears",
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

	entries, ok := s.resolveLeaderboardEntries(c.Request.Context(), lbType)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid leaderboard type"})
		return
	}

	if s.cache != nil && len(entries) > 0 {
		s.cache.Set(c.Request.Context(), cacheKey, entries, 5*time.Minute)
	}

	s.sendPaginatedLeaderboard(c, lbType, entries, page, pageSize)
}

func (s *Server) resolveLeaderboardEntries(ctx context.Context, lbType string) ([]LeaderboardEntry, bool) {
	switch lbType {
	case "skills":
		return s.skillsLeaderboard(ctx), true
	case "playtime":
		return s.statLeaderboard(ctx, "minecraft:custom", "minecraft:play_time"), true
	case "mining":
		return s.minedTotalLeaderboard(ctx), true
	case "killing":
		return s.statLeaderboard(ctx, "minecraft:custom", "minecraft:mob_kills"), true
	case "deaths":
		return s.statLeaderboard(ctx, "minecraft:custom", "minecraft:deaths"), true
	case "walking":
		return s.statLeaderboard(ctx, "minecraft:custom", "minecraft:walk_one_cm"), true
	case "pvp":
		return s.statLeaderboard(ctx, "minecraft:custom", "minecraft:player_kills"), true
	}

	if skillName, ok := parseMcmmoLeaderboardType(lbType); ok {
		return s.mcmmoSkillLeaderboard(ctx, skillName), true
	}

	if category, stat, ok := parseStatLeaderboardType(lbType); ok {
		return s.statLeaderboard(ctx, category, stat), true
	}

	return nil, false
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
	skills, err := s.loadAllMcmmoSkills(ctx)
	if err != nil || len(skills) == 0 {
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

func (s *Server) mcmmoSkillLeaderboard(ctx context.Context, skillName string) []LeaderboardEntry {
	skills, err := s.loadAllMcmmoSkills(ctx)
	if err != nil || len(skills) == 0 {
		return nil
	}

	var entries []LeaderboardEntry
	for _, sk := range skills {
		value, ok := mcmmoSkillValue(sk, skillName)
		if !ok || value <= 0 {
			continue
		}
		if !s.isValidPlayer(sk.UUID) {
			continue
		}

		name := sk.User
		if info := s.store.GetPlayer(sk.UUID); info != nil {
			name = info.LastKnownName
		}

		entries = append(entries, LeaderboardEntry{
			UUID:  sk.UUID,
			Name:  name,
			Value: int64(value),
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

func (s *Server) loadAllMcmmoSkills(ctx context.Context) ([]database.McmmoSkills, error) {
	if s.mcmmoSkillsLoader == nil {
		return nil, nil
	}

	return s.mcmmoSkillsLoader(ctx)
}

func parseMcmmoLeaderboardType(lbType string) (string, bool) {
	skillName, ok := strings.CutPrefix(lbType, "mcmmo:")
	if !ok || skillName == "" {
		return "", false
	}

	for _, candidate := range mcmmoSkillNames {
		if candidate == skillName {
			return skillName, true
		}
	}

	return "", false
}

func parseStatLeaderboardType(lbType string) (string, string, bool) {
	remainder, ok := strings.CutPrefix(lbType, "stat:")
	if !ok || remainder == "" {
		return "", "", false
	}

	parts := strings.Split(remainder, ":")
	if len(parts) < 4 {
		return "", "", false
	}

	category := strings.Join(parts[:2], ":")
	stat := strings.Join(parts[2:], ":")
	if category == "" || stat == "" {
		return "", "", false
	}

	return category, stat, true
}

func mcmmoSkillValue(sk database.McmmoSkills, skillName string) (int, bool) {
	switch skillName {
	case "taming":
		return sk.Taming, true
	case "mining":
		return sk.Mining, true
	case "woodcutting":
		return sk.Woodcutting, true
	case "repair":
		return sk.Repair, true
	case "unarmed":
		return sk.Unarmed, true
	case "herbalism":
		return sk.Herbalism, true
	case "excavation":
		return sk.Excavation, true
	case "archery":
		return sk.Archery, true
	case "swords":
		return sk.Swords, true
	case "axes":
		return sk.Axes, true
	case "acrobatics":
		return sk.Acrobatics, true
	case "fishing":
		return sk.Fishing, true
	case "alchemy":
		return sk.Alchemy, true
	case "crossbows":
		return sk.Crossbows, true
	case "tridents":
		return sk.Tridents, true
	case "maces":
		return sk.Maces, true
	case "spears":
		return sk.Spears, true
	default:
		return 0, false
	}
}
