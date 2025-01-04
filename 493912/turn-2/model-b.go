package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	rateLimit int
	window    []time.Time
	callback  func(current int) int
}

func NewRateLimiter(rateLimit int, windowDuration time.Duration, callback func(current int) int) *RateLimiter {
	return &RateLimiter{
		rateLimit: rateLimit,
		window:    make([]time.Time, 0, rateLimit),
		callback:  callback,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	// Remove expired requests from the sliding window
	rl.window = removeExpired(rl.window, now)

	if len(rl.window) >= rl.rateLimit {
		return false
	}

	// Add the current request to the window
	rl.window = append(rl.window, now)

	// Adjust the rate limit dynamically
	rl.rateLimit = rl.callback(len(rl.window))

	return true
}

func removeExpired(requests []time.Time, now time.Time) []time.Time {
	result := make([]time.Time, 0, len(requests))
	for _, t := range requests {
		if now.Sub(t) < time.Second {
			result = append(result, t)
		}
	}
	return result
}

func dynamicRateLimitCallback(currentCount int) int {
	suggestedLimit := int(math.Max(float64(5), math.Min(float64(currentCount*2), float64(20))))
	return suggestedLimit
}

func main() {
	// Create a rate limiter with a default rate limit of 10 requests per second
	rateLimiter := NewRateLimiter(10, time.Second, dynamicRateLimitCallback)

	// Simulate request handling in multiple goroutines
	go func() {
		for {
			if rateLimiter.Allow() {
				fmt.Println("Request allowed.")
			} else {
				fmt.Println("Request denied.")
			}
			time.Sleep(time.Millisecond * 50)
		}
	}()

	go func() {
		for {
			if rateLimiter.Allow() {
				fmt.Println("Request allowed.")
			} else {
				fmt.Println("Request denied.")
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	time.Sleep(10 * time.Second) // Simulate request handling for 10 seconds
}
