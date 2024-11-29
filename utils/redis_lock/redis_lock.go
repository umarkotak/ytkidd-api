package redis_lock

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	RedisLock struct {
		Prefix    string
		RedisAddr string
		RedisPwd  string

		client *redis.Client
	}
)

var (
	redisLock RedisLock
)

func Initialize(r RedisLock) error {
	client := redis.NewClient(&redis.Options{
		Addr:     r.RedisAddr,
		Password: r.RedisPwd,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	redisLock = RedisLock{
		Prefix:    r.Prefix,
		RedisAddr: r.RedisAddr,
		RedisPwd:  r.RedisPwd,

		client: client,
	}

	return nil
}

func Lock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	key = fmt.Sprintf("%s:%s", redisLock.Prefix, key)

	// Use SETNX command with NX option to acquire the lock
	result, err := redisLock.client.SetNX(ctx, key, "locked", ttl).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set lock: %w", err)
	}

	if result {
		return true, nil
	}
	return false, nil
}

func Unlock(key string) error {
	// Use DEL command to remove the lock
	_, err := redisLock.client.Del(context.Background(), key).Result()
	if err != nil {
		return fmt.Errorf("failed to unlock: %w", err)
	}
	return nil
}
