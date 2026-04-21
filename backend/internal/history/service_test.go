package history

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/mtn-server/backend/config"
)

type fakeRepo struct {
	playerBuckets map[string]map[string]bool
	serverBuckets map[string]int
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{
		playerBuckets: make(map[string]map[string]bool),
		serverBuckets: make(map[string]int),
	}
}

func (f *fakeRepo) InsertPlayerPresenceBatch(_ context.Context, bucketStart string, uuids []string) error {
	for _, uuid := range uuids {
		if _, ok := f.playerBuckets[uuid]; !ok {
			f.playerBuckets[uuid] = make(map[string]bool)
		}
		f.playerBuckets[uuid][bucketStart] = true
	}
	return nil
}

func (f *fakeRepo) UpsertServerHourlyMax(_ context.Context, bucketStart string, maxPlayers int) error {
	if maxPlayers > f.serverBuckets[bucketStart] {
		f.serverBuckets[bucketStart] = maxPlayers
	}
	return nil
}

func (f *fakeRepo) LoadServerCounts(_ context.Context, startBucket string, endBucket string) (map[string]int, error) {
	result := make(map[string]int)
	for key, value := range f.serverBuckets {
		if key >= startBucket && key < endBucket {
			result[key] = value
		}
	}
	return result, nil
}

func (f *fakeRepo) LoadPlayerPresence(_ context.Context, playerUUID string, startBucket string, endBucket string) (map[string]bool, error) {
	result := make(map[string]bool)
	for key, value := range f.playerBuckets[playerUUID] {
		if key >= startBucket && key < endBucket {
			result[key] = value
		}
	}
	return result, nil
}

func (f *fakeRepo) DeleteOlderThan(_ context.Context, cutoffBucket string) error {
	for key := range f.serverBuckets {
		if key < cutoffBucket {
			delete(f.serverBuckets, key)
		}
	}

	for uuid, buckets := range f.playerBuckets {
		for key := range buckets {
			if key < cutoffBucket {
				delete(buckets, key)
			}
		}
		if len(buckets) == 0 {
			delete(f.playerBuckets, uuid)
		}
	}

	return nil
}

type fakeCache struct {
	values map[string]int
}

func newFakeCache() *fakeCache {
	return &fakeCache{values: make(map[string]int)}
}

func (f *fakeCache) Set(_ context.Context, key string, value any, _ time.Duration) error {
	intValue, _ := value.(int)
	f.values[key] = intValue
	return nil
}

func (f *fakeCache) Get(_ context.Context, key string, dest any) (bool, error) {
	value, ok := f.values[key]
	if !ok {
		return false, nil
	}
	if intDest, ok := dest.(*int); ok {
		*intDest = value
	}
	return true, nil
}

func (f *fakeCache) Delete(_ context.Context, key string) error {
	delete(f.values, key)
	return nil
}

func TestRecordSampleUsesConfiguredTimezoneBucket(t *testing.T) {
	repo := newFakeRepo()
	cache := newFakeCache()
	service, err := NewService(repo, cache, config.HistoryConfig{
		Timezone:      "Asia/Shanghai",
		RetentionDays: 7,
	})
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	sampledAt := time.Date(2026, time.April, 20, 16, 45, 0, 0, time.UTC)
	if err := service.RecordSample(context.Background(), sampledAt, []string{"player-a"}, 4); err != nil {
		t.Fatalf("record sample: %v", err)
	}

	expectedBucket := "2026-04-21 00:00:00"
	if !repo.playerBuckets["player-a"][expectedBucket] {
		t.Fatalf("expected player bucket %q to be recorded", expectedBucket)
	}
	if cache.values[serverHourCachePrefix+expectedBucket] != 4 {
		t.Fatalf("expected cached server max for %q", expectedBucket)
	}
}

func TestRecordSampleDeduplicatesPlayersWithinHour(t *testing.T) {
	repo := newFakeRepo()
	service, err := NewService(repo, nil, config.HistoryConfig{
		Timezone:      "Asia/Shanghai",
		RetentionDays: 7,
	})
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	sampledAt := time.Date(2026, time.April, 21, 9, 10, 0, 0, time.UTC)
	if err := service.RecordSample(context.Background(), sampledAt, []string{"player-a", "player-a", "player-b"}, 2); err != nil {
		t.Fatalf("record sample: %v", err)
	}

	bucket := "2026-04-21 17:00:00"
	if len(repo.playerBuckets["player-a"]) != 1 || !repo.playerBuckets["player-a"][bucket] {
		t.Fatalf("expected player-a to have exactly one bucket entry")
	}
	if len(repo.playerBuckets["player-b"]) != 1 || !repo.playerBuckets["player-b"][bucket] {
		t.Fatalf("expected player-b to have exactly one bucket entry")
	}
}

func TestFlushAndPrunePersistsServerMaxFromCache(t *testing.T) {
	repo := newFakeRepo()
	cache := newFakeCache()
	service, err := NewService(repo, cache, config.HistoryConfig{
		Timezone:      "Asia/Shanghai",
		RetentionDays: 7,
	})
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	cache.values[serverHourCachePrefix+"2026-04-20 23:00:00"] = 5
	now := time.Date(2026, time.April, 21, 1, 15, 0, 0, service.location)
	if err := service.FlushAndPrune(context.Background(), now); err != nil {
		t.Fatalf("flush and prune: %v", err)
	}

	if repo.serverBuckets["2026-04-20 23:00:00"] != 5 {
		t.Fatalf("expected flushed server max to be persisted")
	}
	if _, ok := cache.values[serverHourCachePrefix+"2026-04-20 23:00:00"]; ok {
		t.Fatalf("expected flushed cache key to be removed")
	}
}

func TestFlushAndPruneRemovesExpiredRows(t *testing.T) {
	repo := newFakeRepo()
	service, err := NewService(repo, nil, config.HistoryConfig{
		Timezone:      "Asia/Shanghai",
		RetentionDays: 7,
	})
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	repo.serverBuckets["2026-04-13 23:00:00"] = 2
	repo.serverBuckets["2026-04-15 00:00:00"] = 4
	repo.playerBuckets["player-a"] = map[string]bool{
		"2026-04-13 23:00:00": true,
		"2026-04-15 00:00:00": true,
	}

	now := time.Date(2026, time.April, 21, 12, 0, 0, 0, service.location)
	if err := service.FlushAndPrune(context.Background(), now); err != nil {
		t.Fatalf("flush and prune: %v", err)
	}

	if _, ok := repo.serverBuckets["2026-04-13 23:00:00"]; ok {
		t.Fatalf("expected expired server bucket to be removed")
	}
	if repo.serverBuckets["2026-04-15 00:00:00"] != 4 {
		t.Fatalf("expected retained server bucket to remain")
	}
	if repo.playerBuckets["player-a"]["2026-04-13 23:00:00"] {
		t.Fatalf("expected expired player bucket to be removed")
	}
	if !repo.playerBuckets["player-a"]["2026-04-15 00:00:00"] {
		t.Fatalf("expected retained player bucket to remain")
	}
}

func TestGetHeatmapsShapeRollingGrid(t *testing.T) {
	repo := newFakeRepo()
	cache := newFakeCache()
	service, err := NewService(repo, cache, config.HistoryConfig{
		Timezone:      "Asia/Shanghai",
		RetentionDays: 7,
	})
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	repo.serverBuckets["2026-04-16 01:00:00"] = 3
	repo.serverBuckets["2026-04-21 12:00:00"] = 7
	repo.playerBuckets["player-a"] = map[string]bool{
		"2026-04-16 01:00:00": true,
		"2026-04-21 12:00:00": true,
	}
	cache.values[serverHourCachePrefix+"2026-04-21 13:00:00"] = 9

	now := time.Date(2026, time.April, 21, 13, 30, 0, 0, service.location)
	serverHeatmap, err := service.GetServerHeatmap(context.Background(), now)
	if err != nil {
		t.Fatalf("get server heatmap: %v", err)
	}
	playerHeatmap, err := service.GetPlayerHeatmap(context.Background(), "player-a", now)
	if err != nil {
		t.Fatalf("get player heatmap: %v", err)
	}

	expectedDays := []HeatmapDay{
		{Date: "2026-04-15", Weekday: "Wed"},
		{Date: "2026-04-16", Weekday: "Thu"},
		{Date: "2026-04-17", Weekday: "Fri"},
		{Date: "2026-04-18", Weekday: "Sat"},
		{Date: "2026-04-19", Weekday: "Sun"},
		{Date: "2026-04-20", Weekday: "Mon"},
		{Date: "2026-04-21", Weekday: "Tue"},
	}
	if !reflect.DeepEqual(serverHeatmap.Days, expectedDays) {
		t.Fatalf("unexpected heatmap days: %#v", serverHeatmap.Days)
	}
	if len(serverHeatmap.Hours) != 24 || serverHeatmap.Hours[0] != 1 || serverHeatmap.Hours[23] != 24 {
		t.Fatalf("unexpected hour labels: %#v", serverHeatmap.Hours)
	}
	if serverHeatmap.Cells[1][1] != 3 {
		t.Fatalf("expected 2026-04-16 01:00 count to be 3")
	}
	if serverHeatmap.Cells[6][12] != 7 {
		t.Fatalf("expected stored current-day count to be 7")
	}
	if serverHeatmap.Cells[6][13] != 9 {
		t.Fatalf("expected cached current-hour count to be 9")
	}
	if serverHeatmap.WeeklyMaxPlayers != 9 {
		t.Fatalf("expected weekly max of 9, got %d", serverHeatmap.WeeklyMaxPlayers)
	}
	if !playerHeatmap.Cells[1][1] || !playerHeatmap.Cells[6][12] {
		t.Fatalf("expected player online cells to be true")
	}
	if playerHeatmap.Cells[0][0] {
		t.Fatalf("expected empty cell to remain false")
	}
}
