// Package monitor provides server status polling and system stats.
package monitor

import (
	"context"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mcstatus-io/mcutil/v4/response"
	"github.com/mcstatus-io/mcutil/v4/status"
	"github.com/mtn-server/backend/config"
)

// ServerStatus holds the current status and individual connection latencies.
type ServerStatus struct {
	Stats        *JavaStatus                  `json:"stats"`
	BedrockStats *BedrockStatus               `json:"bedrock_stats"`
	System       *SystemStats                 `json:"system"`
	Connections  map[string]*ConnectionStatus `json:"connections"`
	Updated      time.Time                    `json:"updated"`
}

// ConnectionStatus holds latency and status for a specific connection method.
type ConnectionStatus struct {
	Online    bool          `json:"online"`
	Latency   time.Duration `json:"latency"`
	LatencyMs int64         `json:"latency_ms"`
}

// JavaStatus holds Java server ping result.
type JavaStatus struct {
	Online     bool     `json:"online"`
	Players    int      `json:"players"`
	PlayerList []string `json:"player_list"`
	Version    string   `json:"version"`
	MOTD       string   `json:"motd"`
}

// BedrockStatus holds Bedrock server ping result.
type BedrockStatus struct {
	Online     bool     `json:"online"`
	Players    int      `json:"players"`
	PlayerList []string `json:"player_list"`
	Version    string   `json:"version"`
	MOTD       string   `json:"motd"`
}

type MonitoredAddr struct {
	Label   string
	Host    string
	Port    uint16
	Type    string // "java" or "bedrock"
	Dynamic bool
}

// Monitor periodically polls server status.
type Monitor struct {
	targets  []MonitoredAddr
	interval time.Duration
	cfg      *config.Config // Keep config for query ports

	mu     sync.RWMutex
	status *ServerStatus
	cancel context.CancelFunc
}

// NewMonitor creates a new server status monitor for multiple addresses.
func NewMonitor(cfg *config.Config) *Monitor {
	jh, jp := parseHostPort(cfg.Addresses.JavaIPv6, 25565)
	bh, bp := parseHostPort(cfg.Addresses.BedrockIPv6, 19132)

	targets := []MonitoredAddr{
		{Label: "java_ipv6", Host: jh, Port: jp, Type: "java"},
		{Label: "bedrock_ipv6", Host: bh, Port: bp, Type: "bedrock"},
	}

	return &Monitor{
		targets:  targets,
		interval: time.Duration(cfg.StatusRefreshSec) * time.Second,
		cfg:      cfg,
		status: &ServerStatus{
			Connections: make(map[string]*ConnectionStatus),
		},
	}
}

// AddTarget adds a new address to monitor.
func (m *Monitor) AddTarget(label, addr string, port uint16, t string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.targets = append(m.targets, MonitoredAddr{Label: label, Host: addr, Port: port, Type: t})
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

// SyncDynamicTargets replaces all dynamic targets with new ones.
func (m *Monitor) SyncDynamicTargets(targets []MonitoredAddr) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var newTargets []MonitoredAddr
	// Keep static targets
	for _, t := range m.targets {
		if !t.Dynamic {
			newTargets = append(newTargets, t)
		}
	}
	// Add new dynamic targets
	for _, t := range targets {
		t.Dynamic = true
		newTargets = append(newTargets, t)
	}
	m.targets = newTargets
}

// GetStatus returns the latest server status.
func (m *Monitor) GetStatus() *ServerStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.status
}

func (m *Monitor) poll(ctx context.Context) {
	s := &ServerStatus{
		Connections: make(map[string]*ConnectionStatus),
		Updated:     time.Now(),
	}

	// Poll all targets for latencies
	var wg sync.WaitGroup
	m.mu.RLock()
	targets := m.targets
	m.mu.RUnlock()

	for _, target := range targets {
		wg.Add(1)
		go func(t MonitoredAddr) {
			defer wg.Done()
			start := time.Now()
			online := false

			pollCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			if t.Type == "java" {
				if _, err := status.Modern(pollCtx, t.Host, t.Port); err == nil {
					online = true
				}
			} else {
				if _, err := status.Bedrock(pollCtx, t.Host, t.Port); err == nil {
					online = true
				}
			}

			m.mu.Lock()
			lat := time.Since(start)
			s.Connections[t.Label] = &ConnectionStatus{
				Online:    online,
				Latency:   lat,
				LatencyMs: lat.Milliseconds(),
			}
			m.mu.Unlock()
		}(target)
	}
	wg.Wait()

	// Pick the first java/bedrock targets as the "main" status
	// (Usually java_ipv6/bedrock_ipv6)
	s.Stats = m.pollJava(ctx)
	s.BedrockStats = m.pollBedrock(ctx)
	s.System = getSystemStats()

	m.mu.Lock()
	m.status = s
	m.mu.Unlock()
}

func (m *Monitor) pollJava(ctx context.Context) *JavaStatus {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Default to first java target
	m.mu.RLock()
	var host string
	var port uint16
	for _, t := range m.targets {
		if t.Type == "java" {
			host = t.Host
			port = t.Port
			break
		}
	}
	m.mu.RUnlock()

	result, err := status.Modern(ctx, host, port)
	if err != nil {
		return &JavaStatus{Online: false}
	}

	js := parseJavaResult(result)

	// Populate player list from sample
	js.PlayerList = []string{}
	for _, p := range result.Players.Sample {
		// Filter out players with zero UUIDs (often fake/bot indicators)
		if p.ID == "00000000-0000-0000-0000-000000000000" {
			continue
		}
		name := p.Name.Clean
		if name == "" {
			continue
		}
		// Filter out BE_ or . prefixed players for the main Java list
		if !strings.HasPrefix(name, ".") && !strings.HasPrefix(name, "BE_") {
			js.PlayerList = append(js.PlayerList, name)
		}
	}
	js.Players = len(js.PlayerList)

	return js
}

func parseJavaResult(r *response.StatusModern) *JavaStatus {
	js := &JavaStatus{Online: true}

	if r.Players.Online != nil {
		js.Players = int(*r.Players.Online)
	}
	js.Version = r.Version.Name.Clean
	js.MOTD = r.MOTD.Clean

	return js
}

func (m *Monitor) pollBedrock(ctx context.Context) *BedrockStatus {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m.mu.RLock()
	var host string
	var port uint16
	for _, t := range m.targets {
		if t.Type == "bedrock" {
			host = t.Host
			port = t.Port
			break
		}
	}
	m.mu.RUnlock()

	result, err := status.Bedrock(ctx, host, port)
	if err != nil {
		return &BedrockStatus{Online: false}
	}

	bs := &BedrockStatus{Online: true}

	if result.OnlinePlayers != nil {
		bs.Players = int(*result.OnlinePlayers)
	}

	// Bedrock online count adjustment:
	// If it's a Geyser/Floodgate server, we should count players from the Java status sample
	if m.status != nil && m.status.Stats != nil {
		// We use the already polled Java status to count Bedrock players
		beCount := 0
		bs.PlayerList = []string{}

		// Note: We need to re-poll or use the latest Java status sample.
		// Since pollJava was called just before pollBedrock, m.status.Stats might not be updated yet.
		// But in this implementation, pollBedrock is called after pollJava in same poll cycle,
		// and poll() update m.status at the END. So we should use the one we just got.
		// Wait, let's pass it in or rely on the fact that Geyser players are in the Java sample.

		// Re-poll Java to get the latest sample for Bedrock counting if needed,
		// but for efficiency we can just use the one we just calculated in poll()
		// Actually, let's simplify and just use the Java status if available.
		for _, t := range m.targets {
			if t.Type == "java" {
				res, err := status.Modern(ctx, t.Host, t.Port)
				if err == nil {
					for _, p := range res.Players.Sample {
						name := p.Name.Clean
						if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "BE_") {
							beCount++
							bs.PlayerList = append(bs.PlayerList, name)
						}
					}
				}
				break
			}
		}

		if beCount > 0 {
			bs.Players = beCount
		}
	}

	if result.Version != nil {
		bs.Version = *result.Version
	}
	bs.MOTD = result.MOTD.Clean

	return bs
}

func parseHostPort(addr string, defaultPort uint16) (string, uint16) {
	if addr == "" {
		return "", defaultPort
	}

	// Case 1: Bracketed IPv6: [2001:db8::1] or [2001:db8::1]:port
	if strings.HasPrefix(addr, "[") {
		end := strings.LastIndex(addr, "]")
		if end > 0 {
			host := addr[1:end]
			portPart := addr[end+1:]
			if strings.HasPrefix(portPart, ":") {
				if p, err := strconv.ParseUint(portPart[1:], 10, 16); err == nil {
					return host, uint16(p)
				}
			}
			return host, defaultPort
		}
	}

	// Case 2: Standard net.SplitHostPort (handles host:port, ipv4:port)
	host, portStr, err := net.SplitHostPort(addr)
	if err == nil {
		if p, err := strconv.ParseUint(portStr, 10, 16); err == nil {
			return host, uint16(p)
		}
		return host, defaultPort
	}

	// Case 3: Literal IPv6 without brackets but with multiple colons
	// We check if it looks like an IPv6 by counting colons.
	// If it has >= 2 colons but was not handled by SplitHostPort, it's likely an IP without port.
	if strings.Count(addr, ":") >= 2 {
		return addr, defaultPort
	}

	// Case 4: Just a host or IPv4 address without port
	return addr, defaultPort
}
