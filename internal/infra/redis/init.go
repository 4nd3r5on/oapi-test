package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

// NewCache accepts a standard Redis URL:
//
//	redis://:password@host:port/db
//	rediss://:password@host:port/db  (TLS)
func NewCache(ctx context.Context, cacheURL string) (*Cache, error) {
	opts, err := redis.ParseURL(cacheURL)
	if err != nil {
		return nil, fmt.Errorf("redis: parse url: %w", err)
	}
	return &Cache{Client: redis.NewClient(opts)}, nil
}
