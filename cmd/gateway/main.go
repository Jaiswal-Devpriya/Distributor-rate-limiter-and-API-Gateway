package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Jaiswal-Devpriya/Distributor-rate-limiter-and-API-Gateway/internal/limiter"
	"github.com/Jaiswal-Devpriya/Distributor-rate-limiter-and-API-Gateway/internal/middleware"
	"github.com/Jaiswal-Devpriya/Distributor-rate-limiter-and-API-Gateway/internal/proxy"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	port := getEnv("PORT", "8080")

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	rateLimiter := limiter.NewTokenBucketLimiter(redisClient, 5, 5, time.Minute)

	router := proxy.NewRouter()
	handler := middleware.RateLimitMiddleware(rateLimiter, router)

	log.Printf("API Gateway running on port %s", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}