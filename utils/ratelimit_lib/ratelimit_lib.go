package ratelimit_lib

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	RateLimiter struct {
		Prefix    string
		RedisAddr string
		RedisPwd  string

		client *redis.Client
	}
)

var (
	rateLimiter RateLimiter
)

func Initialize(r RateLimiter) error {
	client := redis.NewClient(&redis.Options{
		Addr:     r.RedisAddr,
		Password: r.RedisPwd,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	rateLimiter = RateLimiter{
		Prefix:    r.Prefix,
		RedisAddr: r.RedisAddr,
		RedisPwd:  r.RedisPwd,

		client: client,
	}

	return nil
}

func Check(ctx context.Context, key string, limit int64, duration time.Duration) (bool, error) {
	pipe := rateLimiter.client.Pipeline()

	currentCount := pipe.Incr(
		ctx, fmt.Sprintf("%s:RATELIMIT:%s", rateLimiter.Prefix, key),
	)
	pipe.Expire(ctx, key, duration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return true, fmt.Errorf("failed to execute Redis pipeline: %w", err)
	}

	count := currentCount.Val()
	if count > limit {
		return false, fmt.Errorf("rate limit exceeded for key %s", key)
	}

	return true, nil
}
