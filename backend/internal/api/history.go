package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mtn-server/backend/internal/history"
)

func (s *Server) handleStatusHistory(c *gin.Context) {
	if s.history == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "presence history unavailable"})
		return
	}

	heatmap, err := s.history.GetServerHeatmap(c.Request.Context(), time.Now())
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, history.ErrUnavailable) {
			statusCode = http.StatusServiceUnavailable
		}
		c.JSON(statusCode, gin.H{"error": "failed to load server history"})
		return
	}

	c.JSON(http.StatusOK, heatmap)
}
