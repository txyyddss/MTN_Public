package api

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/database"
)

// PlayerListResponse represents list of players
type PlayerListResponse struct {
	Players    []*data.PlayerInfo `json:"players"`
	Count      int                `json:"count"`
	ActiveDays int                `json:"active_days,omitempty"`
}

// PlayerDetailResponse represents detailed player info
type PlayerDetailResponse struct {
	Info          *data.PlayerInfo         `json:"info"`
	Stats         *data.PlayerStats        `json:"stats"`
	Advancements  *data.PlayerAdvancements `json:"advancements"`
	McMMO         interface{}              `json:"mcmmo"`
	LinkedAccount *database.LinkedPlayer   `json:"linked_account"`
	OreStats      []OreData                `json:"ore_stats"`
	Ranks         map[string]int           `json:"ranks"`
}

// handlePlayers returns active players or search results.
func (s *Server) handlePlayers(c *gin.Context) {
	search := c.Query("search")
	all := c.Query("all") == "true"

	var players []*data.PlayerInfo
	if search != "" {
		players = s.store.SearchPlayers(search)
	} else if all {
		players = s.store.GetAllPlayers()
	} else {
		players = s.store.GetActivePlayers(s.cfg.ActiveDays)
	}

	// Sort by last seen descending
	sort.Slice(players, func(i, j int) bool {
		return players[i].LastSeen > players[j].LastSeen
	})

	validPlayers := s.filterValidPlayers(c.Request.Context(), players)

	c.JSON(http.StatusOK, PlayerListResponse{
		Players:    validPlayers,
		Count:      len(validPlayers),
		ActiveDays: s.cfg.ActiveDays,
	})
}

func (s *Server) getMcmmoValidUUIDs(ctx context.Context) map[string]bool {
	valid := make(map[string]bool)
	if s.mcmmoDB == nil {
		return valid
	}

	if s.cache != nil {
		var cached []string
		if ok, _ := s.cache.Get(ctx, "mcmmo_valid_uuids", &cached); ok {
			for _, u := range cached {
				valid[u] = true
			}
			return valid
		}
	}

	skills, err := s.mcmmoDB.GetAllSkills(ctx)
	if err == nil {
		var uuids []string
		for _, sk := range skills {
			if sk.Total > 0 {
				valid[sk.UUID] = true
				uuids = append(uuids, sk.UUID)
			}
		}
		if s.cache != nil {
			s.cache.Set(ctx, "mcmmo_valid_uuids", uuids, 5*time.Minute)
		}
	}
	return valid
}

func (s *Server) isValidPlayer(ctx context.Context, uuid string, mcmmoValid map[string]bool) bool {
	ps := s.store.GetPlayerStats(uuid)
	if ps == nil || len(ps.Stats) == 0 {
		return false
	}
	pa := s.store.GetPlayerAdvancements(uuid)
	if pa == nil || len(pa.Advancements) == 0 {
		return false
	}
	if s.mcmmoDB != nil && !mcmmoValid[uuid] {
		return false
	}
	return true
}

func (s *Server) filterValidPlayers(ctx context.Context, players []*data.PlayerInfo) []*data.PlayerInfo {
	mcmmoValid := s.getMcmmoValidUUIDs(ctx)

	var valid []*data.PlayerInfo
	for _, p := range players {
		if !s.isValidPlayer(ctx, p.UUID, mcmmoValid) {
			continue
		}
		valid = append(valid, p)
	}
	return valid
}

// handleRandomPlayer redirects to a random player's detail.
func (s *Server) handleRandomPlayer(c *gin.Context) {
	players := s.store.GetActivePlayers(s.cfg.ActiveDays)
	if len(players) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no active players"})
		return
	}
	idx := rand.Intn(len(players))
	c.JSON(http.StatusOK, gin.H{"uuid": players[idx].UUID})
}

// handlePlayerDetail returns full player information.
func (s *Server) handlePlayerDetail(c *gin.Context) {
	uuid := c.Param("uuid")

	info := s.store.GetPlayer(uuid)
	if info == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "player not found"})
		return
	}

	// Only return detail for valid users
	if !s.isValidPlayer(c.Request.Context(), uuid, s.getMcmmoValidUUIDs(c.Request.Context())) {
		c.JSON(http.StatusForbidden, gin.H{"error": "player is not valid (missing data)"})
		return
	}

	stats := s.store.GetPlayerStats(uuid)
	advancements := s.store.GetPlayerAdvancements(uuid)

	// McMMO skills
	var mcmmoSkills interface{}
	if s.mcmmoDB != nil {
		skills, err := s.mcmmoDB.GetSkillsByUUID(c.Request.Context(), uuid)
		if err == nil && skills != nil {
			mcmmoSkills = skills
		}
	}

	// Linked account
	var linkedAccount *database.LinkedPlayer
	if s.floodDB != nil {
		if strings.HasPrefix(uuid, "00000000-0000-0000") {
			linked, err := s.floodDB.GetLinkedByBedrockUUID(c.Request.Context(), uuid)
			if err == nil && linked != nil {
				linkedAccount = linked
			}
		} else {
			linked, err := s.floodDB.GetLinkedByJavaUUID(c.Request.Context(), uuid)
			if err == nil && linked != nil {
				linkedAccount = linked
				// Look up Bedrock username in store
				if bi := s.store.GetPlayer(linked.BedrockUUID); bi != nil {
					linkedAccount.BedrockUsername = bi.CleanName
				}
			}
		}
	}

	// Compute ore stats
	oreStats := computeOreStats(stats)

	// Compute ranks for major categories AND all sub-items
	ranks := s.computeAllRanks(c.Request.Context(), uuid, stats, mcmmoSkills)

	c.JSON(http.StatusOK, PlayerDetailResponse{
		Info:          info,
		Stats:         stats,
		Advancements:  advancements,
		McMMO:         mcmmoSkills,
		LinkedAccount: linkedAccount,
		OreStats:      oreStats,
		Ranks:         ranks,
	})
}

func (s *Server) computeAllRanks(ctx context.Context, uuid string, ps *data.PlayerStats, mcmmo interface{}) map[string]int {
	ranks := make(map[string]int)

	// Broad categories from leaderboards
	categories := []string{"skills", "playtime", "mining", "killing", "deaths", "walking", "pvp"}
	for _, cat := range categories {
		if rank := s.getPlayerRank(ctx, cat, uuid); rank > 0 {
			ranks[cat] = rank
		}
	}

	// Per-item Stats rankings
	if ps != nil {
		for cat, items := range ps.Stats {
			for stat := range items {
				rankKey := fmt.Sprintf("stat:%s:%s", cat, stat)
				// We only calculate ranks for items the player actually has value in
				if rank := s.calculateOrGetRank(ctx, rankKey, uuid); rank > 0 {
					ranks[rankKey] = rank
				}
			}
		}
	}

	// Per-skill McMMO rankings
	if mcmmo != nil {
		if m, ok := mcmmo.(*database.McmmoSkills); ok {
			skills := []struct {
				name string
				val  int
			}{
				{"taming", m.Taming}, {"mining", m.Mining}, {"woodcutting", m.Woodcutting},
				{"repair", m.Repair}, {"unarmed", m.Unarmed}, {"herbalism", m.Herbalism},
				{"excavation", m.Excavation}, {"archery", m.Archery}, {"swords", m.Swords},
				{"axes", m.Axes}, {"acrobatics", m.Acrobatics}, {"fishing", m.Fishing},
				{"alchemy", m.Alchemy}, {"crossbows", m.Crossbows}, {"tridents", m.Tridents},
				{"maces", m.Maces}, {"spears", m.Spears},
			}
			for _, sk := range skills {
				if sk.val > 0 {
					rankKey := "mcmmo:" + sk.name
					if rank := s.calculateOrGetRank(ctx, rankKey, uuid); rank > 0 {
						ranks[rankKey] = rank
					}
				}
			}
		}
	}

	return ranks
}

func (s *Server) calculateOrGetRank(ctx context.Context, lbType, uuid string) int {
	// Try getting from cache first (populated by leaderboards or background)
	if rank := s.getPlayerRank(ctx, lbType, uuid); rank > 0 {
		return rank
	}

	// Fallback/Draft ranking: this is expensive but ensures "New" stats get ranks
	// In a production environment, we'd have a worker pre-calculating common ones.
	return 0
}

func (s *Server) getPlayerRank(ctx context.Context, lbType, uuid string) int {
	var entries []LeaderboardEntry
	cacheKey := "lb_" + lbType

	// Try cache first
	if s.cache != nil {
		if ok, _ := s.cache.Get(ctx, cacheKey, &entries); ok {
			for _, e := range entries {
				if e.UUID == uuid {
					return e.Rank
				}
			}
		}
	}

	// We don't calculate individual ranks on the fly here to avoid performance issues
	return 0
}

// OreData holds mined and used counts for an ore type.
type OreData struct {
	Name  string `json:"name"`
	Mined int64  `json:"mined"`
	Used  int64  `json:"used"`
}

func computeOreStats(ps *data.PlayerStats) []OreData {
	if ps == nil {
		return nil
	}

	ores := []string{
		"minecraft:coal_ore", "minecraft:deepslate_coal_ore",
		"minecraft:iron_ore", "minecraft:deepslate_iron_ore",
		"minecraft:gold_ore", "minecraft:deepslate_gold_ore",
		"minecraft:diamond_ore", "minecraft:deepslate_diamond_ore",
		"minecraft:emerald_ore", "minecraft:deepslate_emerald_ore",
		"minecraft:lapis_ore", "minecraft:deepslate_lapis_ore",
		"minecraft:redstone_ore", "minecraft:deepslate_redstone_ore",
		"minecraft:copper_ore", "minecraft:deepslate_copper_ore",
		"minecraft:nether_quartz_ore", "minecraft:nether_gold_ore",
		"minecraft:ancient_debris",
	}

	oreNames := map[string]string{
		"coal":     "Coal",
		"iron":     "Iron",
		"gold":     "Gold",
		"diamond":  "Diamond",
		"emerald":  "Emerald",
		"lapis":    "Lapis",
		"redstone": "Redstone",
		"copper":   "Copper",
		"quartz":   "Quartz",
		"ancient":  "Ancient Debris",
	}

	minedMap := make(map[string]int64)
	if mined, ok := ps.Stats["minecraft:mined"]; ok {
		for _, ore := range ores {
			if count, exists := mined[ore]; exists {
				category := getOreCategory(ore)
				minedMap[category] += count
			}
		}
	}

	usedMap := make(map[string]int64)
	if used, ok := ps.Stats["minecraft:used"]; ok {
		for _, ore := range ores {
			if count, exists := used[ore]; exists {
				category := getOreCategory(ore)
				usedMap[category] += count
			}
		}
	}

	result := []OreData{}
	for category, displayName := range oreNames {
		mined := minedMap[category]
		used := usedMap[category]
		if mined > 0 || used > 0 {
			result = append(result, OreData{
				Name:  displayName,
				Mined: mined,
				Used:  used,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Mined > result[j].Mined
	})

	return result
}

func getOreCategory(ore string) string {
	switch {
	case strings.Contains(ore, "coal"):
		return "coal"
	case strings.Contains(ore, "iron"):
		return "iron"
	case strings.Contains(ore, "gold") && !strings.Contains(ore, "nether"):
		return "gold"
	case strings.Contains(ore, "nether_gold"):
		return "gold"
	case strings.Contains(ore, "diamond"):
		return "diamond"
	case strings.Contains(ore, "emerald"):
		return "emerald"
	case strings.Contains(ore, "lapis"):
		return "lapis"
	case strings.Contains(ore, "redstone"):
		return "redstone"
	case strings.Contains(ore, "copper"):
		return "copper"
	case strings.Contains(ore, "quartz"):
		return "quartz"
	case strings.Contains(ore, "ancient"):
		return "ancient"
	default:
		return "other"
	}
}

// RefreshAllMcmmoRanks calculates and caches rankings for all mcMMO skills for all players.
func (s *Server) RefreshAllMcmmoRanks(ctx context.Context) error {
	if s.mcmmoDB == nil || s.cache == nil {
		return fmt.Errorf("mcmmo db or cache unavailable")
	}

	skills, err := s.mcmmoDB.GetAllSkills(ctx)
	if err != nil {
		return fmt.Errorf("get all skills: %w", err)
	}

	mcmmoValid := s.getMcmmoValidUUIDs(ctx)

	skillNames := []string{
		"taming", "mining", "woodcutting", "repair", "unarmed",
		"herbalism", "excavation", "archery", "swords", "axes",
		"acrobatics", "fishing", "alchemy", "crossbows", "tridents",
		"maces", "spears",
	}

	for _, skillName := range skillNames {
		var entries []LeaderboardEntry
		for _, sk := range skills {
			if !s.isValidPlayer(ctx, sk.UUID, mcmmoValid) {
				continue
			}

			val := 0
			switch skillName {
			case "taming":
				val = sk.Taming
			case "mining":
				val = sk.Mining
			case "woodcutting":
				val = sk.Woodcutting
			case "repair":
				val = sk.Repair
			case "unarmed":
				val = sk.Unarmed
			case "herbalism":
				val = sk.Herbalism
			case "excavation":
				val = sk.Excavation
			case "archery":
				val = sk.Archery
			case "swords":
				val = sk.Swords
			case "axes":
				val = sk.Axes
			case "acrobatics":
				val = sk.Acrobatics
			case "fishing":
				val = sk.Fishing
			case "alchemy":
				val = sk.Alchemy
			case "crossbows":
				val = sk.Crossbows
			case "tridents":
				val = sk.Tridents
			case "maces":
				val = sk.Maces
			case "spears":
				val = sk.Spears
			}

			if val > 0 {
				name := sk.User
				if info := s.store.GetPlayer(sk.UUID); info != nil {
					name = info.LastKnownName
				}
				entries = append(entries, LeaderboardEntry{
					UUID:  sk.UUID,
					Name:  name,
					Value: int64(val),
				})
			}
		}

		if len(entries) > 0 {
			ranked := rankEntries(entries)
			cacheKey := "lb_mcmmo:" + skillName
			if err := s.cache.Set(ctx, cacheKey, ranked, 1*time.Hour); err != nil {
				fmt.Printf("Error caching mcmmo rank for %s: %v\n", skillName, err)
			}
		}
	}

	return nil
}
