// MTN Server Backend - Entry Point
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/api"
	"github.com/mtn-server/backend/internal/cache"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/database"
	"github.com/mtn-server/backend/internal/lucky"
	"github.com/mtn-server/backend/internal/monitor"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded from %s", *configPath)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize Redis cache
	cacheClient, err := cache.New(cfg.Redis)
	if err != nil {
		log.Printf("Warning: Redis unavailable: %v (continuing without cache)", err)
		cacheClient = nil
	} else {
		log.Println("Redis connected")
	}

	// Load world data
	store := data.NewStore(cfg.WorldDir)
	if err := store.LoadAll(); err != nil {
		log.Printf("Warning: Failed to load world data: %v", err)
	} else {
		log.Printf("World data loaded from %s", cfg.WorldDir)
	}

	// Start periodic data refresh
	go func() {
		ticker := time.NewTicker(time.Duration(cfg.DataRefreshSec) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := store.LoadAll(); err != nil {
					log.Printf("Data refresh error: %v", err)
				} else {
					log.Println("Data refreshed")
				}
			}
		}
	}()

	// Initialize MySQL databases
	var mcmmoDB *database.McmmoDB
	var floodDB *database.FloodgateDB

	mcmmoDB, err = database.NewMcmmoDB(cfg.McmmoMySQL)
	if err != nil {
		log.Printf("Warning: McMMO DB unavailable: %v", err)
		mcmmoDB = nil
	} else {
		log.Println("McMMO DB connected")
		defer mcmmoDB.Close()
	}

	floodDB, err = database.NewFloodgateDB(cfg.FloodgateMySQL)
	if err != nil {
		log.Printf("Warning: Floodgate DB unavailable: %v", err)
		floodDB = nil
	} else {
		log.Println("Floodgate DB connected")
		defer floodDB.Close()
	}

	// Initialize Lucky client
	luckyClient := lucky.New(cfg.Lucky, cfg.Addresses)
	if err := luckyClient.Refresh(ctx); err != nil {
		log.Printf("Warning: Lucky API unavailable: %v", err)
	} else {
		log.Println("Lucky API connected")
	}

	// Start periodic Lucky refresh
	go func() {
		ticker := time.NewTicker(time.Duration(cfg.StatusRefreshSec) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := luckyClient.Refresh(ctx); err != nil {
					log.Printf("Lucky refresh error: %v", err)
				}
			}
		}
	}()

	// Start server status monitor
	mon := monitor.NewMonitor(
		cfg.LocalConnection.Java,
		cfg.LocalConnection.Bedrock,
		cfg.StatusRefreshSec,
	)
	mon.Start(ctx)
	defer mon.Stop()
	log.Println("Server status monitor started")

	// Create API server
	srv := api.NewServer(cfg, store, mcmmoDB, floodDB, luckyClient, mon, cacheClient)
	router := srv.SetupRouter()

	// Start HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	go func() {
		log.Printf("API server starting on :%d", cfg.Server.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	cancel()

	if cacheClient != nil {
		cacheClient.Close()
	}

	log.Println("Server stopped")
}
