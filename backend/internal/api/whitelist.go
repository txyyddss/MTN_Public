package api

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mtn-server/backend/internal/whitelist"
)

type whitelistMutationRequest struct {
	Edition    string `json:"edition"`
	PlayerName string `json:"player_name"`
	QQID       string `json:"qq_id"`
}

type whitelistListResponse struct {
	Entries []whitelist.Entry `json:"entries"`
	Count   int               `json:"count"`
}

func (s *Server) handleWhitelistList(c *gin.Context) {
	if !s.requireWhitelistToken(c) {
		return
	}
	if s.whitelist == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "whitelist service unavailable", "code": "unavailable"})
		return
	}

	edition, err := whitelist.NormalizeEdition(c.DefaultQuery("edition", string(whitelist.EditionAll)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid edition", "code": whitelist.ErrorCode(err)})
		return
	}

	entries, err := s.whitelist.List(c.Request.Context(), edition)
	if err != nil {
		s.writeWhitelistError(c, err)
		return
	}

	c.Header("Cache-Control", "no-store")
	c.JSON(http.StatusOK, whitelistListResponse{
		Entries: entries,
		Count:   len(entries),
	})
}

func (s *Server) handleWhitelistAdd(c *gin.Context) {
	s.handleWhitelistMutation(c, whitelist.ActionAdd)
}

func (s *Server) handleWhitelistRemove(c *gin.Context) {
	s.handleWhitelistMutation(c, whitelist.ActionRemove)
}

func (s *Server) handleWhitelistMutation(c *gin.Context, action whitelist.Action) {
	if !s.requireWhitelistToken(c) {
		return
	}
	if s.whitelist == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "whitelist service unavailable", "code": "unavailable"})
		return
	}

	var payload whitelistMutationRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "code": "invalid_input"})
		return
	}

	edition, err := whitelist.NormalizeEdition(payload.Edition)
	if err != nil || edition == whitelist.EditionAll {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid edition", "code": "invalid_input"})
		return
	}

	req := whitelist.OperationRequest{
		Action:       action,
		Edition:      edition,
		PlayerName:   strings.TrimSpace(payload.PlayerName),
		QQID:         strings.TrimSpace(payload.QQID),
		ActorQQID:    "api",
		Source:       whitelist.SourceAPI,
		Admin:        true,
		EnforceQuota: strings.TrimSpace(payload.QQID) != "",
	}

	var result *whitelist.OperationResult
	switch action {
	case whitelist.ActionAdd:
		result, err = s.whitelist.Add(c.Request.Context(), req)
	case whitelist.ActionRemove:
		result, err = s.whitelist.Remove(c.Request.Context(), req)
	default:
		err = whitelist.ErrInvalidInput
	}
	if err != nil {
		s.writeWhitelistError(c, err)
		return
	}

	c.Header("Cache-Control", "no-store")
	c.JSON(http.StatusOK, result)
}

func (s *Server) requireWhitelistToken(c *gin.Context) bool {
	expected := s.cfg.Whitelist.APIToken
	if expected == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "whitelist API token is not configured", "code": "unavailable"})
		return false
	}

	const prefix = "Bearer "
	header := c.GetHeader("Authorization")
	if !strings.HasPrefix(header, prefix) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token", "code": "unauthorized"})
		return false
	}

	actual := strings.TrimSpace(strings.TrimPrefix(header, prefix))
	if subtle.ConstantTimeCompare([]byte(actual), []byte(expected)) != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid bearer token", "code": "unauthorized"})
		return false
	}

	return true
}

func (s *Server) writeWhitelistError(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, whitelist.ErrUnavailable):
		status = http.StatusServiceUnavailable
	case errors.Is(err, whitelist.ErrInvalidInput):
		status = http.StatusBadRequest
	case errors.Is(err, whitelist.ErrQuotaExceeded), errors.Is(err, whitelist.ErrConflict):
		status = http.StatusConflict
	case errors.Is(err, whitelist.ErrForbidden):
		status = http.StatusForbidden
	case errors.Is(err, whitelist.ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, whitelist.ErrRCONRejected):
		status = http.StatusBadGateway
	}

	c.JSON(status, gin.H{
		"error": err.Error(),
		"code":  whitelist.ErrorCode(err),
	})
}
