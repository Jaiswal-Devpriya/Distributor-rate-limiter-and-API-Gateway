package middleware

import (
	"net/http"

	"github.com/Jaiswal-Devpriya/Distributor-rate-limiter-and-API-Gateway/internal/limiter"
)

func RateLimitMiddleware(rateLimiter *limiter.TokenBucketLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Header.Get("X-Client-ID")
		if clientID == "" {
			clientID = r.RemoteAddr
		}

		allowed, err := rateLimiter.Allow(r.Context(), clientID)
		if err != nil {
			http.Error(w, "rate limiter error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":"rate limit exceeded"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}