package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimiter controls the rate of requests
type RateLimiter struct {
	tokens    chan struct{}
	waitGroup sync.WaitGroup
}

// NewRateLimiter creates a new RateLimiter allowing `rate` requests per second
func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, rate),
	}
	// Fill the token bucket at the specified rate
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}()
	return rl
}

// LimitRequest limits incoming requests
func (rl *RateLimiter) LimitRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-rl.tokens:
			defer func() {
				rl.waitGroup.Done()
			}()
			rl.waitGroup.Add(1)
			next(w, r)
		default:
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	}
}

// Sample handler responding with JSON
func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Request successful!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	rate := 5 // Allow a maximum of 5 requests per second
	rl := NewRateLimiter(rate)

	http.HandleFunc("/api", rl.LimitRequest(apiHandler))
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
