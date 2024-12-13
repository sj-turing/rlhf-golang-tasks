package main

import (
	"fmt"
	"net/http"
	"time"
)

// This will act as the limiter
type RateLimiter struct {
	rate    int
	burst   int
	limiter *time.Ticker
}

// NewRateLimiter creates a RateLimiter that allows `rate` requests per second
func NewRateLimiter(rate int, burst int) *RateLimiter {
	limiter := time.NewTicker(time.Second / time.Duration(rate))
	return &RateLimiter{
		rate:    rate,
		burst:   burst,
		limiter: limiter,
	}
}

// LimitRequest handles the rate limiting logic
func (rl *RateLimiter) LimitRequest(next http.HandlerFunc) http.HandlerFunc {
	semaphore := make(chan struct{}, rl.burst)

	return func(w http.ResponseWriter, r *http.Request) {
		// Wait for the ticker to tick which indicates available request
		semaphore <- struct{}{}
		<-rl.limiter.C
		defer func() { <-semaphore }()

		next(w, r) // Call the next handler
	}
}

// Example API endpoint
func apiEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Request successful!")
}

func main() {
	rateLimiter := NewRateLimiter(10, 5) // 10 requests per second, burst of 5

	http.HandleFunc("/api", rateLimiter.LimitRequest(apiEndpoint))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
