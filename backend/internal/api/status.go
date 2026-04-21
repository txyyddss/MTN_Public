package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/lucky"
	"github.com/mtn-server/backend/internal/monitor"
)

type statusEditionResponse struct {
	Online  bool `json:"online"`
	Players int  `json:"players"`
}

type statusSystemResponse struct {
	Platform   string  `json:"platform"`
	CPUModel   string  `json:"cpu_model"`
	MemUsed    uint64  `json:"mem_used"`
	MemPercent float64 `json:"mem_percent"`
	Load1      float64 `json:"load_1"`
	Load5      float64 `json:"load_5"`
	Load15     float64 `json:"load_15"`
}

type statusResponse struct {
	Java          *statusEditionResponse `json:"java"`
	Bedrock       *statusEditionResponse `json:"bedrock"`
	System        *statusSystemResponse  `json:"system"`
	OnlinePlayers []string               `json:"online_players"`
	Updated       any                    `json:"updated"`
}

// ConnectionResponse holds detailed connection output.
type ConnectionResponse struct {
	Connection *lucky.ConnectionInfo  `json:"connection"`
	Addresses  config.AddressesConfig `json:"addresses"`
}

func slimEditionStatus(online bool, players int) *statusEditionResponse {
	return &statusEditionResponse{
		Online:  online,
		Players: players,
	}
}

func slimSystemStatus(system *monitor.SystemStats) *statusSystemResponse {
	if system == nil {
		return nil
	}

	return &statusSystemResponse{
		Platform:   system.Platform,
		CPUModel:   system.CPUModel,
		MemUsed:    system.MemUsed,
		MemPercent: system.MemPercent,
		Load1:      system.Load1,
		Load5:      system.Load5,
		Load15:     system.Load15,
	}
}

// handleStatus returns real-time server status.
func (s *Server) handleStatus(c *gin.Context) {
	status := s.monitor.GetStatus()
	c.Header("Cache-Control", "public, max-age=5, stale-while-revalidate=10")

	if status == nil {
		c.JSON(http.StatusOK, statusResponse{})
		return
	}

	var java *statusEditionResponse
	if status.Stats != nil {
		java = slimEditionStatus(status.Stats.Online, status.Stats.Players)
	}

	var bedrock *statusEditionResponse
	if status.BedrockStats != nil {
		bedrock = slimEditionStatus(status.BedrockStats.Online, status.BedrockStats.Players)
	}

	c.JSON(http.StatusOK, statusResponse{
		Java:          java,
		Bedrock:       bedrock,
		System:        slimSystemStatus(status.System),
		OnlinePlayers: status.OnlinePlayers,
		Updated:       status.Updated,
	})
}

func (s *Server) handleConnection(c *gin.Context) {
	connInfo, err := s.luckyClient.GetConnectionInfo()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   "connection data unavailable",
			"details": err.Error(),
		})
		return
	}

	c.Header("Cache-Control", "public, max-age=300, stale-while-revalidate=600")
	c.JSON(http.StatusOK, ConnectionResponse{
		Connection: connInfo,
		Addresses:  s.cfg.Addresses,
	})
}
