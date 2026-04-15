package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleStatus returns real-time server status.
func (s *Server) handleStatus(c *gin.Context) {
	status := s.monitor.GetStatus()
	c.JSON(http.StatusOK, status)
}

// handleConnection returns server connection information.
func (s *Server) handleConnection(c *gin.Context) {
	connInfo := s.luckyClient.GetConnectionInfo()
	c.JSON(http.StatusOK, gin.H{
		"connection": connInfo,
		"addresses":  s.cfg.Addresses,
	})
}
