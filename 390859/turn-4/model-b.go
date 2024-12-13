package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

type RateLimiter struct {
	mu             sync.Mutex
	hourlyRequests map[string][]time.Time
	dailyRequests  map[string][]time.Time
	hourlyLimit    int
	dailyLimit     int
}

func NewRateLimiter(hourlyLimit, dailyLimit int) *RateLimiter {
	return &RateLimiter{
		hourlyRequests: make(map[string][]time.Time),
		dailyRequests:  make(map[string][]time.Time),
		hourlyLimit:    hourlyLimit,
		dailyLimit:     dailyLimit,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// Trim old timestamps
	rl.trimRequests(now, &rl.hourlyRequests[userID], time.Hour)
	rl.trimRequests(now, &rl.dailyRequests[userID], 24*time.Hour)

	// Check hourly limit
	if len(rl.hourlyRequests[userID]) >= rl.hourlyLimit {
		return false
	}

	// Check daily limit
	if len(rl.dailyRequests[userID]) >= rl.dailyLimit {
		return false
	}

	// Add new request timestamp
	rl.hourlyRequests[userID] = append(rl.hourlyRequests[userID], now)
	rl.dailyRequests[userID] = append(rl.dailyRequests[userID], now)

	return true
}

func (rl *RateLimiter) trimRequests(now time.Time, requests *[]time.Time, window time.Duration) {
	for len(*requests) > 0 && (*requests)[0].Add(window).Before(now) {
		*requests = (*requests)[1:]
	}
	sort.Times(*requests)
}

func main() {
	rateLimiter := NewRateLimiter(3, 10) // 3 OTPs per hour, 10 OTPs per day

	ctx := context.Background()
	userID := "user123"

	for i := 0; i < 15; i++ { // Test with 15 requests
		if rateLimiter.Allow(ctx, userID) {
			fmt.Printf("OTP sent to %s at %v\n", userID, time.Now())
		} else {
			fmt.Printf("Rate limit exceeded for %s\n", userID)
		}
		time.Sleep(time.Second * 2) // Simulate time between requests
	}
}
