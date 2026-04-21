// Package config provides JSON-based configuration for the MTN backend.
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds all configurable parameters for the backend.
type Config struct {
	Server           ServerConfig          `json:"server"`
	WorldDir         string                `json:"world_dir"`
	Redis            RedisConfig           `json:"redis"`
	McmmoMySQL       MySQLConfig           `json:"mcmmo_mysql"`
	FloodgateMySQL   MySQLConfig           `json:"floodgate_mysql"`
	HistoryMySQL     MySQLConfig           `json:"history_mysql"`
	History          HistoryConfig         `json:"history"`
	Lucky            LuckyConfig           `json:"lucky"`
	LocalConnection  LocalConnectionConfig `json:"local_connection"`
	Addresses        AddressesConfig       `json:"addresses"`
	StatusRefreshSec int                   `json:"status_refresh_seconds"`
	DataRefreshSec   int                   `json:"data_refresh_seconds"`
	ActiveDays       int                   `json:"active_days"`
}

// LocalConnectionConfig holds local query addresses.
type LocalConnectionConfig struct {
	Java    string `json:"java"`
	Bedrock string `json:"bedrock"`
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Port        int      `json:"port"`
	CORSOrigins []string `json:"cors_origins"`
}

// RedisConfig holds Redis connection settings.
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// MySQLConfig holds MySQL connection settings.
type MySQLConfig struct {
	DSN string `json:"dsn"`
}

// HistoryConfig holds online history settings.
type HistoryConfig struct {
	Timezone      string `json:"timezone"`
	RetentionDays int    `json:"retention_days"`
}

// LuckyConfig holds Lucky STUN/DDNS API settings.
type LuckyConfig struct {
	ServerURL string `json:"server_url"`
	Token     string `json:"token"`
}

// AddressesConfig holds static server address configuration.
type AddressesConfig struct {
	JavaIPv6    string `json:"java_ipv6"`
	BedrockIPv6 string `json:"bedrock_ipv6"`
	JavaIPv4SRV string `json:"java_ipv4_srv"`
	JavaIPv6SRV string `json:"java_ipv6_srv"`
}

// Load reads and parses configuration from the given JSON file path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	cfg := &Config{
		Server: ServerConfig{
			Port:        8080,
			CORSOrigins: []string{"*"},
		},
		StatusRefreshSec: 5,
		DataRefreshSec:   300,
		ActiveDays:       7,
		History: HistoryConfig{
			Timezone:      "Local",
			RetentionDays: 7,
		},
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}
