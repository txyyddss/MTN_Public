package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/lucky"
)

// ConnectionResponse holds detailed connection output.
type ConnectionResponse struct {
	Connection *lucky.ConnectionInfo  `json:"connection"`
	Addresses  config.AddressesConfig `json:"addresses"`
}

// handleStatus returns real-time server status.
func (s *Server) handleStatus(c *gin.Context) {
	status := s.monitor.GetStatus()
	c.JSON(http.StatusOK, status)
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
	c.JSON(http.StatusOK, ConnectionResponse{
		Connection: connInfo,
		Addresses:  s.cfg.Addresses,
	})
}
