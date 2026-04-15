// Package data provides parsers for Minecraft world data files.
package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// PlayerStats holds parsed stats for a single player.
type PlayerStats struct {
	UUID  string                      `json:"uuid"`
	Stats map[string]map[string]int64 `json:"stats"`
}

// PlayerAdvancement holds a single advancement entry.
type PlayerAdvancement struct {
	Key      string            `json:"key"`
	Done     bool              `json:"done"`
	Criteria map[string]string `json:"criteria"`
}

// PlayerAdvancements holds all advancements for a player.
type PlayerAdvancements struct {
	UUID         string              `json:"uuid"`
	Advancements []PlayerAdvancement `json:"advancements"`
}

// PlayerInfo holds parsed NBT playerdata.
type PlayerInfo struct {
	UUID          string  `json:"uuid"`
	XpLevel       int     `json:"xp_level"`
	TicksLived    int64   `json:"ticks_lived"`
	FoodLevel     int     `json:"food_level"`
	Health        float64 `json:"health"`
	LastSeen      int64   `json:"last_seen"`    // milliseconds
	FirstPlayed   int64   `json:"first_played"` // milliseconds
	LastKnownName string  `json:"last_known_name"`
}

// Store holds all preloaded player data in memory.
type Store struct {
	mu           sync.RWMutex
	worldDir     string
	Players      map[string]*PlayerInfo         // uuid -> info
	Stats        map[string]*PlayerStats        // uuid -> stats
	Advancements map[string]*PlayerAdvancements // uuid -> advancements
}

// NewStore creates a new data store for the given world directory.
func NewStore(worldDir string) *Store {
	return &Store{
		worldDir:     worldDir,
		Players:      make(map[string]*PlayerInfo),
		Stats:        make(map[string]*PlayerStats),
		Advancements: make(map[string]*PlayerAdvancements),
	}
}

// LoadAll loads all player data from disk into memory.
func (s *Store) LoadAll() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.loadPlayerData(); err != nil {
		return fmt.Errorf("load playerdata: %w", err)
	}
	if err := s.loadStats(); err != nil {
		return fmt.Errorf("load stats: %w", err)
	}
	if err := s.loadAdvancements(); err != nil {
		return fmt.Errorf("load advancements: %w", err)
	}

	return nil
}

// GetActivePlayers returns players active within the given number of days.
func (s *Store) GetActivePlayers(days int) []*PlayerInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cutoff := time.Now().Add(-time.Duration(days) * 24 * time.Hour).UnixMilli()
	var result []*PlayerInfo
	for _, p := range s.Players {
		if p.LastSeen >= cutoff {
			result = append(result, p)
		}
	}
	return result
}

// GetPlayer returns player info by UUID.
func (s *Store) GetPlayer(uuid string) *PlayerInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Players[uuid]
}

// GetPlayerStats returns stats for a player by UUID.
func (s *Store) GetPlayerStats(uuid string) *PlayerStats {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Stats[uuid]
}

// GetPlayerAdvancements returns advancements for a player by UUID.
func (s *Store) GetPlayerAdvancements(uuid string) *PlayerAdvancements {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Advancements[uuid]
}

// SearchPlayers does fuzzy search by name or UUID.
func (s *Store) SearchPlayers(query string) []*PlayerInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	query = strings.ToLower(query)
	var result []*PlayerInfo
	for uuid, p := range s.Players {
		if strings.Contains(strings.ToLower(p.LastKnownName), query) ||
			strings.Contains(strings.ToLower(uuid), query) {
			result = append(result, p)
		}
	}
	return result
}

func (s *Store) loadStats() error {
	dir := filepath.Join(s.worldDir, "stats")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read stats dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		uuid := strings.TrimSuffix(entry.Name(), ".json")
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}

		var raw struct {
			Stats map[string]map[string]int64 `json:"stats"`
		}
		if err := json.Unmarshal(data, &raw); err != nil {
			continue
		}

		s.Stats[uuid] = &PlayerStats{
			UUID:  uuid,
			Stats: raw.Stats,
		}
	}
	return nil
}

func (s *Store) loadAdvancements() error {
	dir := filepath.Join(s.worldDir, "advancements")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read advancements dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		uuid := strings.TrimSuffix(entry.Name(), ".json")
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}

		// Advancements JSON is a flat map of advancement_key -> { criteria, done }
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			continue
		}

		var advs []PlayerAdvancement
		for key, val := range raw {
			// Skip DataVersion key and recipe advancements
			if key == "DataVersion" {
				continue
			}
			if strings.Contains(key, "recipes/") {
				continue
			}

			var entry struct {
				Done     bool              `json:"done"`
				Criteria map[string]string `json:"criteria"`
			}
			if err := json.Unmarshal(val, &entry); err != nil {
				continue
			}

			advs = append(advs, PlayerAdvancement{
				Key:      key,
				Done:     entry.Done,
				Criteria: entry.Criteria,
			})
		}

		s.Advancements[uuid] = &PlayerAdvancements{
			UUID:         uuid,
			Advancements: advs,
		}
	}
	return nil
}
