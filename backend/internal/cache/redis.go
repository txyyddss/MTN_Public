// Package cache provides Redis-based caching for computed backend data.
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/mtn-server/backend/config"
)

// Client wraps a Redis client with convenience methods.
type Client struct {
	rdb *redis.Client
}

// New creates a new Redis cache client.
func New(cfg config.RedisConfig) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}

	return &Client{rdb: rdb}, nil
}

// Set stores a JSON-serializable value with expiration.
func (c *Client) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal cache value: %w", err)
	}
	return c.rdb.Set(ctx, key, data, ttl).Err()
}

// Get retrieves and unmarshals a cached value. Returns false if key doesn't exist.
func (c *Client) Get(ctx context.Context, key string, dest any) (bool, error) {
	data, err := c.rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("redis get: %w", err)
	}
	if err := json.Unmarshal(data, dest); err != nil {
		return false, fmt.Errorf("unmarshal cache: %w", err)
	}
	return true, nil
}

// Delete removes a key from cache.
func (c *Client) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

// Close shuts down the Redis connection.
func (c *Client) Close() error {
	return c.rdb.Close()
}
