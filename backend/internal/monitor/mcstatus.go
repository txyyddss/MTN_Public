// Package monitor provides server status polling and system stats.
package monitor

import (
	"context"
	"fmt"
	"sync"
	"time"

	mcutil "github.com/mcstatus-io/mcutil/v4"
)

// ServerStatus holds the current status of the Minecraft server.
type ServerStatus struct {
	Java    *JavaStatus    `json:"java"`
	Bedrock *BedrockStatus `json:"bedrock"`
	System  *SystemStats   `json:"system"`
	Updated time.Time      `json:"updated"`
}

// JavaStatus holds Java server ping result.
type JavaStatus struct {
	Online     bool   `json:"online"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
	Version    string `json:"version"`
	MOTD       string `json:"motd"`
	Latency    int64  `json:"latency_ms"`
}

// BedrockStatus holds Bedrock server ping result.
type BedrockStatus struct {
	Online     bool   `json:"online"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
	Version    string `json:"version"`
	MOTD       string `json:"motd"`
}

// Monitor periodically polls server status.
type Monitor struct {
	javaAddr    string
	bedrockAddr string
	interval    time.Duration

	mu     sync.RWMutex
	status *ServerStatus
	cancel context.CancelFunc
}

// NewMonitor creates a new server status monitor.
func NewMonitor(javaAddr, bedrockAddr string, intervalSec int) *Monitor {
	return &Monitor{
		javaAddr:    javaAddr,
		bedrockAddr: bedrockAddr,
		interval:    time.Duration(intervalSec) * time.Second,
		status:      &ServerStatus{},
	}
}

// Start begins periodic polling. Call Stop() to cancel.
func (m *Monitor) Start(ctx context.Context) {
	ctx, m.cancel = context.WithCancel(ctx)

	// Initial poll
	m.poll(ctx)

	go func() {
		ticker := time.NewTicker(m.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				m.poll(ctx)
			}
		}
	}()
}

// Stop cancels the polling goroutine.
func (m *Monitor) Stop() {
	if m.cancel != nil {
		m.cancel()
	}
}

// GetStatus returns the latest server status.
func (m *Monitor) GetStatus() *ServerStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.status
}

func (m *Monitor) poll(ctx context.Context) {
	status := &ServerStatus{
		Updated: time.Now(),
	}

	// Poll Java
	status.Java = m.pollJava(ctx)

	// Poll Bedrock
	status.Bedrock = m.pollBedrock(ctx)

	// Get system stats
	status.System = getSystemStats()

	m.mu.Lock()
	m.status = status
	m.mu.Unlock()
}

func (m *Monitor) pollJava(ctx context.Context) *JavaStatus {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := mcutil.Status(ctx, m.javaAddr)
	if err != nil {
		return &JavaStatus{Online: false}
	}

	motd := ""
	if result.MOTD != nil {
		motd = fmt.Sprintf("%v", result.MOTD.Clean)
	}

	return &JavaStatus{
		Online:     true,
		Players:    int(*result.Players.Online),
		MaxPlayers: int(*result.Players.Max),
		Version:    result.Version.NameClean,
		MOTD:       motd,
	}
}

func (m *Monitor) pollBedrock(ctx context.Context) *BedrockStatus {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := mcutil.StatusBedrock(ctx, m.bedrockAddr)
	if err != nil {
		return &BedrockStatus{Online: false}
	}

	return &BedrockStatus{
		Online:     true,
		Players:    int(*result.OnlinePlayers),
		MaxPlayers: int(*result.MaxPlayers),
		Version:    *result.Version,
		MOTD:       result.MOTD.Clean,
	}
}
