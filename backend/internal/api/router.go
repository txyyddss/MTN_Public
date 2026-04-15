// Package api provides the HTTP REST API for the MTN backend.
package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/cache"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/database"
	"github.com/mtn-server/backend/internal/lucky"
	"github.com/mtn-server/backend/internal/monitor"
)

// Server holds all dependencies for the API.
type Server struct {
	cfg         *config.Config
	store       *data.Store
	mcmmoDB     *database.McmmoDB
	floodDB     *database.FloodgateDB
	luckyClient *lucky.Client
	monitor     *monitor.Monitor
	cache       *cache.Client
}

// NewServer creates an API server with all dependencies.
func NewServer(
	cfg *config.Config,
	store *data.Store,
	mcmmoDB *database.McmmoDB,
	floodDB *database.FloodgateDB,
	luckyClient *lucky.Client,
	mon *monitor.Monitor,
	cacheClient *cache.Client,
) *Server {
	return &Server{
		cfg:         cfg,
		store:       store,
		mcmmoDB:     mcmmoDB,
		floodDB:     floodDB,
		luckyClient: luckyClient,
		monitor:     mon,
		cache:       cacheClient,
	}
}

// SetupRouter creates the Gin router with all routes.
func (s *Server) SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     s.cfg.Server.CORSOrigins,
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/players", s.handlePlayers)
		api.GET("/players/random", s.handleRandomPlayer)
		api.GET("/players/:uuid", s.handlePlayerDetail)
		api.GET("/leaderboards/:type", s.handleLeaderboard)
		api.GET("/status", s.handleStatus)
		api.GET("/connection", s.handleConnection)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
