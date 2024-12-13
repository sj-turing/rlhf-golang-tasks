package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// LeakyBucket implements a leaky bucket rate limiter
type LeakyBucket struct {
	bucket     int // Current number of requests in the bucket
	capacity   int // Maximum capacity of the bucket
	rate       int // Rate of requests per second
	mutex      sync.Mutex
	stopTicker context.CancelFunc
}

// NewLeakyBucket initializes a new leaky bucket rate limiter
func NewLeakyBucket(rate int, capacity int) *LeakyBucket {
	lb := &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		bucket:   capacity, // Start full
	}
	ticker := time.NewTicker(time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	lb.stopTicker = cancel

	go func() {
		for {
			select {
			case <-ticker.C:
				lb.mutex.Lock()
				lb.bucket += lb.rate
				if lb.bucket > lb.capacity {
					lb.bucket = lb.capacity
				}
				lb.mutex.Unlock()
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()

	return lb
}

// LimitRequest handles incoming requests with rate limiting
func (lb *LeakyBucket) LimitRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lb.mutex.Lock()
		if lb.bucket > 0 {
			lb.bucket--
			lb.mutex.Unlock()
			next(w, r)
		} else {
			lb.mutex.Unlock()
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	}
}

// Sample JSON handler
func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Request successful!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	rate := 1     // 1 request per second
	capacity := 5 // Maximum of 5 requests can be queued
	lb := NewLeakyBucket(rate, capacity)

	http.HandleFunc("/api", lb.LimitRequest(apiHandler))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
