package main

/*

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens     int
	refill     int
	period     time.Duration
	lastRefill time.Time
	mu         sync.Mutex
}

func NewRateLimiter(tokens, refill int, period time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:     tokens,
		refill:     refill,
		period:     period,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Acquire() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.lastRefill.Add(rl.period)) {
		rl.tokens = rl.refill
		rl.lastRefill = now
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

func rateLimitedHandler(rl *RateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.Acquire() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Request successful!")
	}
}

func main() {
	rateLimiter := NewRateLimiter(10, 10, time.Second) // 10 requests per second
	http.HandleFunc("/api", rateLimitedHandler(rateLimiter))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
*/
