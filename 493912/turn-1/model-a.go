package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter struct handles rate limiting logic
type RateLimiter struct {
	mu          sync.Mutex
	rateLimit   int
	cache       map[string]int // IP address or client ID -> request count
	cacheExpiry time.Duration
	callback    func(current, suggested int) int
}

// NewRateLimiter creates a new RateLimiter with a default rate limit and cache expiry
func NewRateLimiter(rateLimit int, cacheExpiry time.Duration, callback func(current, suggested int) int) *RateLimiter {
	return &RateLimiter{
		rateLimit:   rateLimit,
		cache:       make(map[string]int),
		cacheExpiry: cacheExpiry,
		callback:    callback,
	}
}

// Allow checks if a request should be allowed based on the current rate limit
func (rl *RateLimiter) Allow(clientID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Remove expired entries from the cache
	now := time.Now()
	for key, value := range rl.cache {
		if now.Sub(time.Unix(int64(value), 0)) > rl.cacheExpiry {
			delete(rl.cache, key)
		}
	}

	// Check if the request is allowed
	if count, ok := rl.cache[clientID]; ok {
		if count >= rl.rateLimit {
			return false
		}
	}

	// Record the request
	rl.cache[clientID] = time.Now().Unix()

	return true
}

// SetRateLimit dynamically sets or adjusts the rate limit using the callback
func (rl *RateLimiter) SetRateLimit(suggested int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Use the callback to adjust the rate limit
	rl.rateLimit = rl.callback(rl.rateLimit, suggested)
}

// DynamicRateLimitCallback is a sample callback function
func DynamicRateLimitCallback(current, suggested int) int {
	// Example logic: if suggested rate limit is lower than 5, set it to 5
	if suggested < 5 {
		return 5
	}
	return suggested
}

func main() {
	// Create a new rate limiter with a default rate limit of 10 requests per second
	rateLimiter := NewRateLimiter(10, time.Second, DynamicRateLimitCallback)

	// Simulate request handling
	go func() {
		for i := 0; i < 20; i++ {
			clientID := fmt.Sprintf("client_%d", i)
			if rateLimiter.Allow(clientID) {
				fmt.Println("Request allowed from", clientID)
				time.Sleep(time.Duration(100-i) * time.Millisecond)
			} else {
				fmt.Println("Request denied from", clientID)
			}
		}
	}()

	// Simulate a dynamic adjustment of the rate limit
	time.Sleep(2 * time.Second)
	rateLimiter.SetRateLimit(3) // Suggest a new rate limit of 3

	// Continue handling requests
	go func() {
		for i := 20; i < 40; i++ {
			clientID := fmt.Sprintf("client_%d", i)
			if rateLimiter.Allow(clientID) {
				fmt.Println("Request allowed from", clientID)
				time.Sleep(time.Duration(100-i) * time.Millisecond)
			} else {
				fmt.Println("Request denied from", clientID)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
