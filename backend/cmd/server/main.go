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
	"strconv"
	"syscall"
	"time"

	"github.com/mtn-server/backend/config"
	"github.com/mtn-server/backend/internal/api"
	"github.com/mtn-server/backend/internal/cache"
	"github.com/mtn-server/backend/internal/data"
	"github.com/mtn-server/backend/internal/database"
	"github.com/mtn-server/backend/internal/history"
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
	var historyDB *database.HistoryDB

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

	if cfg.HistoryMySQL.DSN != "" {
		historyDB, err = database.NewHistoryDB(cfg.HistoryMySQL)
		if err != nil {
			log.Printf("Warning: History DB unavailable: %v", err)
			historyDB = nil
		} else {
			if err := historyDB.EnsureSchema(ctx); err != nil {
				log.Printf("Warning: History DB schema setup failed: %v", err)
				historyDB.Close()
				historyDB = nil
			} else {
				log.Println("History DB connected")
				defer historyDB.Close()
			}
		}
	}

	// Initialize Lucky client
	luckyClient := lucky.New(cfg.Lucky, cfg.Addresses)
	if err := luckyClient.Refresh(ctx); err != nil {
		log.Printf("Warning: Lucky API unavailable: %v", err)
	} else {
		log.Println("Lucky API connected")
	}

	// Start server status monitor
	mon := monitor.NewMonitor(cfg)

	var historyService *history.Service
	if historyDB != nil {
		historyService, err = history.NewService(historyDB, cacheClient, cfg.History)
		if err != nil {
			log.Printf("Warning: Presence history unavailable: %v", err)
		} else {
			mon.SetPresenceRecorder(historyService)
			go historyService.Start(ctx)
			log.Println("Presence history service started")
		}
	}

	// Add static addresses from config
	if cfg.Addresses.JavaIPv4SRV != "" {
		mon.AddTarget("java_ipv4_srv", cfg.Addresses.JavaIPv4SRV, 25565, "java")
	}
	if cfg.Addresses.JavaIPv6SRV != "" {
		mon.AddTarget("java_ipv6_srv", cfg.Addresses.JavaIPv6SRV, 25565, "java")
	}

	mon.Start(ctx)
	defer mon.Stop()
	log.Println("Server status monitor started")

	// Start periodic Lucky refresh and sync with monitor
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
				} else {
					// Sync dynamic targets to monitor
					info, err := luckyClient.GetConnectionInfo()
					if err == nil {
						var targets []monitor.MonitoredAddr
						if info.JavaIPv4 != nil {
							port, _ := strconv.ParseUint(info.JavaIPv4.Port, 10, 16)
							targets = append(targets, monitor.MonitoredAddr{
								Label: "java_ipv4",
								Host:  info.JavaIPv4.IP,
								Port:  uint16(port),
								Type:  "java",
							})
						}
						if info.BedrockIPv4 != nil {
							port, _ := strconv.ParseUint(info.BedrockIPv4.Port, 10, 16)
							targets = append(targets, monitor.MonitoredAddr{
								Label: "bedrock_ipv4",
								Host:  info.BedrockIPv4.IP,
								Port:  uint16(port),
								Type:  "bedrock",
							})
						}
						mon.SyncDynamicTargets(targets)
					}
				}
			}
		}
	}()

	// Create API server
	srv := api.NewServer(cfg, store, mcmmoDB, floodDB, luckyClient, mon, cacheClient)
	srv.SetPresenceHistory(historyService)
	router := srv.SetupRouter()

	// Start periodic mcMMO and Stats ranking refresh
	if cacheClient != nil {
		go func() {
			// Run once on startup
			if mcmmoDB != nil {
				if err := srv.RefreshAllMcmmoRanks(ctx); err != nil {
					log.Printf("Initial mcMMO ranking refresh error: %v", err)
				} else {
					log.Println("Initial mcMMO ranking refresh complete")
				}
			}

			if err := srv.RefreshAllStatRanks(ctx); err != nil {
				log.Printf("Initial stats ranking refresh error: %v", err)
			} else {
				log.Println("Initial stats ranking refresh complete")
			}

			ticker := time.NewTicker(15 * time.Minute) // Refresh every 15 minutes
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if mcmmoDB != nil {
						if err := srv.RefreshAllMcmmoRanks(ctx); err != nil {
							log.Printf("Periodic mcMMO ranking refresh error: %v", err)
						} else {
							log.Println("Periodic mcMMO ranking refresh complete")
						}
					}
					if err := srv.RefreshAllStatRanks(ctx); err != nil {
						log.Printf("Periodic stats ranking refresh error: %v", err)
					} else {
						log.Println("Periodic stats ranking refresh complete")
					}
				}
			}
		}()
	}

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
