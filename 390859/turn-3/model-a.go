package main

/*

import (
	"fmt"
	"net/http"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
	rate   time.Duration
	burst  int
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, burst),
		rate:   time.Second / time.Duration(rate),
		burst:  burst,
	}

	// Start a goroutine to refill tokens at the specified rate
	go func() {
		ticker := time.NewTicker(rl.rate)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case rl.tokens <- struct{}{}: // Add a token
				default:
					// No need to do anything if the bucket is full
				}
			}
		}
	}()

	return rl
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens: // Try to take a token
		return true
	default:
		return false // No tokens available
	}
}

func rateLimitedHandler(rl *RateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Request successful!")
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, 5) // 5 requests per second, burst of 5
	http.HandleFunc("/api", rateLimitedHandler(rateLimiter))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

*/
