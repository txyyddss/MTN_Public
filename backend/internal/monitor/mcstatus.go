// Package monitor provides server status polling and system stats.
package monitor

import (
	"context"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/mcstatus-io/mcutil/v4/query"
	"github.com/mcstatus-io/mcutil/v4/response"
	"github.com/mcstatus-io/mcutil/v4/status"
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

// Monitor periodically polls server status.
type Monitor struct {
	javaHost    string
	javaPort    uint16
	bedrockHost string
	bedrockPort uint16
	interval    time.Duration

	mu     sync.RWMutex
	status *ServerStatus
	cancel context.CancelFunc
}

// NewMonitor creates a new server status monitor.
// javaAddr and bedrockAddr should be "host:port" format.
func NewMonitor(javaAddr, bedrockAddr string, intervalSec int) *Monitor {
	jh, jp := parseHostPort(javaAddr, 25565)
	bh, bp := parseHostPort(bedrockAddr, 19132)

	return &Monitor{
		javaHost:    jh,
		javaPort:    jp,
		bedrockHost: bh,
		bedrockPort: bp,
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
	s := &ServerStatus{
		Updated: time.Now(),
	}

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

	result, err := query.Full(ctx, m.javaHost, m.javaPort)
	if err != nil {
		if resultMod, errMod := status.Modern(ctx, m.javaHost, m.javaPort); errMod == nil {
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
	js.PlayerList = result.Players

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

	result, err := status.Bedrock(ctx, m.bedrockHost, m.bedrockPort)
	if err != nil {
		return &BedrockStatus{Online: false}
	}

	return parseBedrockResult(result)
}

func parseBedrockResult(r *response.StatusBedrock) *BedrockStatus {
	bs := &BedrockStatus{Online: true}

	if r.OnlinePlayers != nil {
		bs.Players = int(*r.OnlinePlayers)
	}
	if r.Version != nil {
		bs.Version = *r.Version
	}
	bs.MOTD = r.MOTD.Clean

	return bs
}

func parseHostPort(addr string, defaultPort uint16) (string, uint16) {
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
