// Package monitor provides server status polling and system stats.
package monitor

import (
	"context"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mcstatus-io/mcutil/v4/query"
	"github.com/mcstatus-io/mcutil/v4/response"
	"github.com/mcstatus-io/mcutil/v4/status"
	"github.com/mtn-server/backend/config"
)

// ServerStatus holds the current status and individual connection latencies.
type ServerStatus struct {
	Java        *JavaStatus                  `json:"java"`
	Bedrock     *BedrockStatus               `json:"bedrock"`
	System      *SystemStats                 `json:"system"`
	Connections map[string]*ConnectionStatus `json:"connections"`
	Updated     time.Time                    `json:"updated"`
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
	s.Java = m.pollJava(ctx)
	s.Bedrock = m.pollBedrock(ctx)
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
			// Use split query port if it's the main local address
			if m.cfg != nil && (host == m.cfg.LocalConnection.Java || strings.Contains(m.cfg.LocalConnection.Java, host)) {
				port = m.cfg.LocalConnection.JavaQueryPort
			} else {
				port = t.Port
			}
			break
		}
	}
	m.mu.RUnlock()

	result, err := query.Full(ctx, host, port)
	if err != nil {
		if resultMod, errMod := status.Modern(ctx, host, port); errMod == nil {
			return parseJavaResult(resultMod)
		}
		return &JavaStatus{Online: false}
	}

	js := &JavaStatus{Online: true}

	if val, ok := result.Data["version"]; ok {
		js.Version = val
	}
	if val, ok := result.Data["motd"]; ok {
		js.MOTD = val
	}
	if val, ok := result.Data["numplayers"]; ok {
		if i, err := strconv.Atoi(val); err == nil {
			js.Players = i
		}
	}
	js.PlayerList = []string{}
	for _, p := range result.Players {
		if !strings.HasPrefix(p, ".") && !strings.HasPrefix(p, "BE_") {
			js.PlayerList = append(js.PlayerList, p)
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
	// If it's a Geyser/Floodgate server, we should count players from the Java list with prefixes
	if m.status != nil && m.status.Java != nil {
		javaFull, err := query.Full(ctx, m.cfg.LocalConnection.Java, m.cfg.LocalConnection.JavaQueryPort)
		if err == nil {
			beCount := 0
			for _, p := range javaFull.Players {
				if strings.HasPrefix(p, ".") || strings.HasPrefix(p, "BE_") {
					beCount++
				}
			}
			if beCount > 0 {
				bs.Players = beCount
			}
		}
	}

	if result.Version != nil {
		bs.Version = *result.Version
	}
	bs.MOTD = result.MOTD.Clean

	return bs
}

func parseHostPort(addr string, defaultPort uint16) (string, uint16) {
	// Handle bracketed IPv6 without port: [2001:db8::1]
	if strings.HasPrefix(addr, "[") && strings.HasSuffix(addr, "]") {
		return addr[1 : len(addr)-1], defaultPort
	}

	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		// No port, use default
		return addr, defaultPort
	}
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return host, defaultPort
	}
	return host, uint16(port)
}
