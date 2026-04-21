package history

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mtn-server/backend/config"
)

const serverHourCachePrefix = "presence:server-hour:"

// ErrUnavailable reports that presence history storage is not configured.
var ErrUnavailable = errors.New("presence history unavailable")

type historyRepository interface {
	InsertPlayerPresenceBatch(ctx context.Context, bucketStart string, uuids []string) error
	UpsertServerHourlyMax(ctx context.Context, bucketStart string, maxPlayers int) error
	LoadServerCounts(ctx context.Context, startBucket string, endBucket string) (map[string]int, error)
	LoadPlayerPresence(ctx context.Context, playerUUID string, startBucket string, endBucket string) (map[string]bool, error)
	DeleteOlderThan(ctx context.Context, cutoffBucket string) error
}

type cacheStore interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, dest any) (bool, error)
	Delete(ctx context.Context, key string) error
}

// HeatmapDay describes one row in the weekly heatmap grid.
type HeatmapDay struct {
	Date    string `json:"date"`
	Weekday string `json:"weekday"`
}

// ServerHeatmapResponse contains weekly server presence data.
type ServerHeatmapResponse struct {
	Timezone         string       `json:"timezone"`
	Days             []HeatmapDay `json:"days"`
	Hours            []int        `json:"hours"`
	Cells            [][]int      `json:"cells"`
	WeeklyMaxPlayers int          `json:"weekly_max_players"`
}

// PlayerHeatmapResponse contains weekly player presence data.
type PlayerHeatmapResponse struct {
	Timezone string       `json:"timezone"`
	Days     []HeatmapDay `json:"days"`
	Hours    []int        `json:"hours"`
	Cells    [][]bool     `json:"cells"`
}

// Service records and queries rolling online presence history.
type Service struct {
	repo          historyRepository
	cache         cacheStore
	location      *time.Location
	timezoneLabel string
	retentionDays int
}

// NewService creates a history service using the provided storage backends.
func NewService(repo historyRepository, cache cacheStore, cfg config.HistoryConfig) (*Service, error) {
	if repo == nil {
		return nil, ErrUnavailable
	}

	locationName := cfg.Timezone
	if locationName == "" || locationName == "Local" {
		locationName = time.Local.String()
		if locationName == "" {
			locationName = "Local"
		}
	}

	location, err := time.LoadLocation(locationName)
	if err != nil {
		return nil, fmt.Errorf("load history timezone %q: %w", locationName, err)
	}

	retentionDays := cfg.RetentionDays
	if retentionDays <= 0 {
		retentionDays = 7
	}

	return &Service{
		repo:          repo,
		cache:         cache,
		location:      location,
		timezoneLabel: locationName,
		retentionDays: retentionDays,
	}, nil
}

// Start runs periodic flush and prune work until the context is cancelled.
func (s *Service) Start(ctx context.Context) {
	if s == nil {
		return
	}

	if err := s.FlushAndPrune(ctx, time.Now()); err != nil {
		log.Printf("Presence history initial flush error: %v", err)
	}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			if err := s.FlushAndPrune(ctx, now); err != nil {
				log.Printf("Presence history flush error: %v", err)
			}
		}
	}
}

// RecordSample stores one live monitor sample.
func (s *Service) RecordSample(ctx context.Context, sampledAt time.Time, onlineUUIDs []string, onlinePlayers int) error {
	if s == nil || s.repo == nil {
		return ErrUnavailable
	}

	bucketStart := s.hourBucket(sampledAt)
	bucketKey := s.bucketKey(bucketStart)

	if len(onlineUUIDs) > 0 {
		uniqueUUIDs := uniqueStrings(onlineUUIDs)
		if err := s.repo.InsertPlayerPresenceBatch(ctx, bucketKey, uniqueUUIDs); err != nil {
			return err
		}
	}

	if s.cache == nil {
		return s.repo.UpsertServerHourlyMax(ctx, bucketKey, onlinePlayers)
	}

	currentMax := 0
	if ok, err := s.cache.Get(ctx, s.serverHourKey(bucketStart), &currentMax); err != nil {
		return err
	} else if !ok {
		currentMax = 0
	}

	if onlinePlayers > currentMax {
		ttl := time.Duration((s.retentionDays+2)*24) * time.Hour
		if err := s.cache.Set(ctx, s.serverHourKey(bucketStart), onlinePlayers, ttl); err != nil {
			return err
		}
	}

	return nil
}

// FlushAndPrune finalizes past Redis buckets into MySQL and removes expired data.
func (s *Service) FlushAndPrune(ctx context.Context, now time.Time) error {
	if s == nil || s.repo == nil {
		return ErrUnavailable
	}

	cutoffDay := s.windowStart(now)
	if err := s.repo.DeleteOlderThan(ctx, s.bucketKey(cutoffDay)); err != nil {
		return err
	}

	if s.cache == nil {
		return nil
	}

	currentHour := s.hourBucket(now)
	for bucket := cutoffDay; bucket.Before(currentHour); bucket = bucket.Add(time.Hour) {
		var maxPlayers int
		ok, err := s.cache.Get(ctx, s.serverHourKey(bucket), &maxPlayers)
		if err != nil {
			return err
		}
		if !ok {
			continue
		}
		if err := s.repo.UpsertServerHourlyMax(ctx, s.bucketKey(bucket), maxPlayers); err != nil {
			return err
		}
		if err := s.cache.Delete(ctx, s.serverHourKey(bucket)); err != nil {
			return err
		}
	}

	pruneStart := cutoffDay.Add(-48 * time.Hour)
	for bucket := pruneStart; bucket.Before(cutoffDay); bucket = bucket.Add(time.Hour) {
		if err := s.cache.Delete(ctx, s.serverHourKey(bucket)); err != nil {
			return err
		}
	}

	return nil
}

// GetServerHeatmap returns the rolling 7-day server online heatmap.
func (s *Service) GetServerHeatmap(ctx context.Context, now time.Time) (*ServerHeatmapResponse, error) {
	if s == nil || s.repo == nil {
		return nil, ErrUnavailable
	}

	windowStart := s.windowStart(now)
	windowEnd := windowStart.AddDate(0, 0, 7)
	counts, err := s.repo.LoadServerCounts(ctx, s.bucketKey(windowStart), s.bucketKey(windowEnd))
	if err != nil {
		return nil, err
	}

	if s.cache != nil {
		for bucket := windowStart; bucket.Before(windowEnd); bucket = bucket.Add(time.Hour) {
			var pending int
			ok, err := s.cache.Get(ctx, s.serverHourKey(bucket), &pending)
			if err != nil {
				return nil, err
			}
			if ok && pending > counts[s.bucketKey(bucket)] {
				counts[s.bucketKey(bucket)] = pending
			}
		}
	}

	response := &ServerHeatmapResponse{
		Timezone: s.timezoneLabel,
		Days:     s.buildDays(windowStart),
		Hours:    buildHours(),
		Cells:    make([][]int, 0, 7),
	}

	for dayIndex := 0; dayIndex < 7; dayIndex++ {
		dayStart := windowStart.AddDate(0, 0, dayIndex)
		row := make([]int, 24)
		for hour := 0; hour < 24; hour++ {
			bucket := dayStart.Add(time.Duration(hour) * time.Hour)
			value := counts[s.bucketKey(bucket)]
			row[hour] = value
			if value > response.WeeklyMaxPlayers {
				response.WeeklyMaxPlayers = value
			}
		}
		response.Cells = append(response.Cells, row)
	}

	return response, nil
}

// GetPlayerHeatmap returns the rolling 7-day online heatmap for one player.
func (s *Service) GetPlayerHeatmap(ctx context.Context, playerUUID string, now time.Time) (*PlayerHeatmapResponse, error) {
	if s == nil || s.repo == nil {
		return nil, ErrUnavailable
	}

	windowStart := s.windowStart(now)
	windowEnd := windowStart.AddDate(0, 0, 7)
	presence, err := s.repo.LoadPlayerPresence(ctx, playerUUID, s.bucketKey(windowStart), s.bucketKey(windowEnd))
	if err != nil {
		return nil, err
	}

	response := &PlayerHeatmapResponse{
		Timezone: s.timezoneLabel,
		Days:     s.buildDays(windowStart),
		Hours:    buildHours(),
		Cells:    make([][]bool, 0, 7),
	}

	for dayIndex := 0; dayIndex < 7; dayIndex++ {
		dayStart := windowStart.AddDate(0, 0, dayIndex)
		row := make([]bool, 24)
		for hour := 0; hour < 24; hour++ {
			bucket := dayStart.Add(time.Duration(hour) * time.Hour)
			row[hour] = presence[s.bucketKey(bucket)]
		}
		response.Cells = append(response.Cells, row)
	}

	return response, nil
}

func (s *Service) hourBucket(value time.Time) time.Time {
	local := value.In(s.location)
	return time.Date(local.Year(), local.Month(), local.Day(), local.Hour(), 0, 0, 0, s.location)
}

func (s *Service) windowStart(now time.Time) time.Time {
	local := now.In(s.location)
	todayStart := time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, s.location)
	return todayStart.AddDate(0, 0, -(s.retentionDays - 1))
}

func (s *Service) serverHourKey(bucket time.Time) string {
	return serverHourCachePrefix + s.bucketKey(bucket)
}

func (s *Service) bucketKey(bucket time.Time) string {
	return bucket.In(s.location).Format("2006-01-02 15:04:05")
}

func (s *Service) buildDays(windowStart time.Time) []HeatmapDay {
	days := make([]HeatmapDay, 0, 7)
	for dayIndex := 0; dayIndex < 7; dayIndex++ {
		day := windowStart.AddDate(0, 0, dayIndex)
		days = append(days, HeatmapDay{
			Date:    day.Format("2006-01-02"),
			Weekday: day.Format("Mon"),
		})
	}
	return days
}

func buildHours() []int {
	hours := make([]int, 24)
	for hour := 0; hour < 24; hour++ {
		hours[hour] = hour + 1
	}
	return hours
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
