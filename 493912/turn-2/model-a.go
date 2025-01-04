package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RateLimiter struct {
	maxRequests int
	cacheExpiry time.Duration
	cache       map[string]time.Time
	callback    func() int
	lock        sync.RWMutex
}

func NewRateLimiter(maxRequests int, cacheExpiry time.Duration, callback func() int) *RateLimiter {
	rand.Seed(time.Now().UnixNano())
	return &RateLimiter{
		maxRequests: maxRequests,
		cacheExpiry: cacheExpiry,
		cache:       make(map[string]time.Time),
		callback:    callback,
		lock:        sync.RWMutex{},
	}
}

func (rl *RateLimiter) Allow(clientID string) bool {
	rl.lock.RLock()
	defer rl.lock.RUnlock()

	now := time.Now()

	// Check if client entry is expired
	if t, ok := rl.cache[clientID]; ok && now.Sub(t) < rl.cacheExpiry {
		// Deny if over max requests
		if len(rl.cache) >= rl.maxRequests {
			return false
		}
	}

	rl.cache[clientID] = now
	return true
}

func (rl *RateLimiter) AdjustRateLimit() {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	newMaxRequests := rl.callback()
	rl.maxRequests = newMaxRequests
	fmt.Printf("Adjusted rate limit: %d\n", rl.maxRequests)
}

func dynamicRateLimitCallback() int {
	// Example callback logic that dynamically adjusts rate limit
	currentTime := time.Now()
	currentHour := currentTime.Hour()

	if currentHour >= 16 && currentHour < 18 {
		// Peak traffic hour, lower the rate limit
		return 5
	} else if currentHour < 7 {
		// Overnight, lower the rate limit
		return 3
	} else {
		// Default rate limit
		return 10
	}
}

func main() {
	rateLimiter := NewRateLimiter(10, time.Second, dynamicRateLimitCallback)

	go func() {
		for i := 0; ; i++ {
			clientID := fmt.Sprintf("client_%d", rand.Intn(10))
			if rateLimiter.Allow(clientID) {
				fmt.Println("Request allowed from", clientID)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			} else {
				fmt.Println("Request denied from", clientID)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	// Simulate dynamic rate limit adjustment based on time
	rateLimiter.AdjustRateLimit()

	time.Sleep(5 * time.Second)
}
