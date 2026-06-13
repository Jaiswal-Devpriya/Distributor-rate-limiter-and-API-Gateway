package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBucketLimiter struct {
	redisClient *redis.Client
	capacity    int
	refillRate  int
	window      time.Duration
}

func NewTokenBucketLimiter(redisClient *redis.Client, capacity int, refillRate int, window time.Duration) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		redisClient: redisClient,
		capacity:    capacity,
		refillRate:  refillRate,
		window:      window,
	}
}

func (l *TokenBucketLimiter) Allow(ctx context.Context, clientID string) (bool, error) {
	key := "rate_limit:" + clientID

	count, err := l.redisClient.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		l.redisClient.Expire(ctx, key, l.window)
	}

	return count <= int64(l.capacity), nil
}