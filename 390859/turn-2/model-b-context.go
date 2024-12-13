package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu      sync.Mutex
	counter int
	limit   int
	reset   time.Time
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		counter: 0,
		limit:   limit,
		reset:   time.Now(),
	}
}

func (rl *RateLimiter) Limit(ctx context.Context) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.counter >= rl.limit {
		return false
	}

	now := time.Now()
	if now.After(rl.reset.Add(time.Minute)) {
		rl.counter = 0
		rl.reset = now
	}

	rl.counter++
	return true
}

func rateLimitedHandler(rl *RateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if !rl.Limit(ctx) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Request successful!")
	}
}

func main() {
	rateLimiter := NewRateLimiter(10, time.Minute) // 10 requests per minute
	http.HandleFunc("/api", rateLimitedHandler(rateLimiter))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
