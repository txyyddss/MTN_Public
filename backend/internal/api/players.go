package api

import (
	"math/rand"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mtn-server/backend/internal/data"
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
	LinkedAccount interface{}              `json:"linked_account"`
	OreStats      []OreData                `json:"ore_stats"`
}

// handlePlayers returns active players or search results.
func (s *Server) handlePlayers(c *gin.Context) {
	search := c.Query("search")

	if search != "" {
		players := s.store.SearchPlayers(search)
		c.JSON(http.StatusOK, PlayerListResponse{
			Players: players,
			Count:   len(players),
		})
		return
	}

	players := s.store.GetActivePlayers(s.cfg.ActiveDays)
	c.JSON(http.StatusOK, PlayerListResponse{
		Players:    players,
		Count:      len(players),
		ActiveDays: s.cfg.ActiveDays,
	})
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
	var linkedAccount interface{}
	if s.floodDB != nil {
		// Check if this is a Bedrock player (UUID starts with 00000000-0000-0000)
		if strings.HasPrefix(uuid, "00000000-0000-0000") {
			linked, err := s.floodDB.GetLinkedByBedrockUUID(c.Request.Context(), uuid)
			if err == nil && linked != nil {
				linkedAccount = linked
			}
		} else {
			linked, err := s.floodDB.GetLinkedByJavaUUID(c.Request.Context(), uuid)
			if err == nil && linked != nil {
				linkedAccount = linked
			}
		}
	}

	// Compute ore stats (mined vs used/crafted)
	oreStats := computeOreStats(stats)

	c.JSON(http.StatusOK, PlayerDetailResponse{
		Info:          info,
		Stats:         stats,
		Advancements:  advancements,
		McMMO:         mcmmoSkills,
		LinkedAccount: linkedAccount,
		OreStats:      oreStats,
	})
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

	var result []OreData
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
