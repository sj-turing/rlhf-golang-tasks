package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	maxRequests int
	window      time.Duration
	cache       map[string]time.Time
	callback    func(currentRate, newRate int) int
	mutex       sync.RWMutex
}

func NewRateLimiter(maxRequests int, window time.Duration, callback func(currentRate, newRate int) int) *RateLimiter {
	return &RateLimiter{
		maxRequests: maxRequests,
		window:      window,
		cache:       make(map[string]time.Time),
		callback:    callback,
		mutex:       sync.RWMutex{},
	}
}

func (rl *RateLimiter) Allow(clientID string) bool {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()

	now := time.Now()

	// Remove expired entries from the cache
	for key := range rl.cache {
		if now.Sub(rl.cache[key]) > rl.window {
			delete(rl.cache, key)
		}
	}

	// Check if the client has reached the request limit
	if len(rl.cache) >= rl.maxRequests {
		return false
	}

	rl.cache[clientID] = now
	return true
}

func (rl *RateLimiter) AdjustRateLimit() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	// Get the current rate limit
	currentRate := rl.maxRequests
	// Suggest a new rate limit based on an internal condition (for illustration purposes)
	suggestedRate := currentRate * 2
	// Use the callback to adjust the rate limit
	rl.maxRequests = rl.callback(currentRate, suggestedRate)
	fmt.Printf("Adjusted rate limit to: %d\n", rl.maxRequests)
}

func dynamicRateLimitCallback(currentRate, newRate int) int {
	// Example callback logic
	if newRate > currentRate {
		return newRate
	} else {
		// Optionally apply specific logic to determine a valid new rate limit
		return newRate
	}
}

func main() {
	rateLimiter := NewRateLimiter(10, time.Second, dynamicRateLimitCallback)

	// Simulate handling requests
	go func() {
		for i := 0; i < 50; i++ {
			clientID := fmt.Sprintf("client_%d", i%10)
			if rateLimiter.Allow(clientID) {
				fmt.Printf("Request allowed from %s\n", clientID)
				time.Sleep(time.Millisecond * time.Duration(1000-i*10))
			} else {
				fmt.Printf("Request denied from %s\n", clientID)
			}
		}
	}()

	// Simulate a dynamic rate limit adjustment after 2 seconds
	time.Sleep(2 * time.Second)
	rateLimiter.AdjustRateLimit()

	// Continue handling more requests to demonstrate the adjusted limit
	time.Sleep(5 * time.Second)
}
