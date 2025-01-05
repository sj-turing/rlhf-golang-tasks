package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RateLimiter struct
type RateLimiter struct {
	rateLimiter *Limiter   // Underlying rate limiter
	callback    func() int // Callback function to adjust rate limits dynamically
}

// NewRateLimiter creates a new RateLimiter with an initial rate limit and a callback function.
func NewRateLimiter(initialRate int, callback func() int) *RateLimiter {
	return &RateLimiter{
		rateLimiter: NewLimiter(initialRate, time.Second), // Create a new Limiter with the initial rate and 1-second burst size
		callback:    callback,
	}
}

// Allow checks if a request is allowed based on the current rate limit.
func (rl *RateLimiter) Allow() bool {
	if rl.rateLimiter.AllowN(time.Now(), 1) {
		return true
	}

	// If request is denied, call the callback to adjust the rate limit
	newRate := rl.callback()
	rl.rateLimiter.SetRate(bucket.Rate(newRate))
	return rl.rateLimiter.Allow()
}

// Limiter is an implementation of a token bucket rate limiter.
type Limiter struct {
	rate     bucket.Rate
	last     time.Time
	count    int
	capacity int
	lock     sync.Mutex
}

// NewLimiter creates a new Limiter with the specified rate and capacity.
func NewLimiter(rate int, burst time.Duration) *Limiter {
	capacity := int(rate * burst.Seconds())
	return &Limiter{
		rate:     bucket.Rate(rate),
		last:     time.Now(),
		capacity: capacity,
		lock:     sync.Mutex{},
	}
}

// AllowN allows n tokens to be taken from the bucket.
func (l *Limiter) AllowN(now time.Time, n int) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	// Calculate the number of tokens to fill the bucket in this interval.
	added := int((now.Sub(l.last)) * l.rate)
	if added > l.capacity {
		added = l.capacity
	}

	// Update the token count and bucket time.
	l.count += added
	l.last = now

	// Allow the request only if there are enough tokens available.
	if l.count >= n {
		l.count -= n
		return true
	}
	return false
}

// SetRate sets the new rate of the bucket.
func (l *Limiter) SetRate(rate bucket.Rate) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.rate = rate
}

// dynamicRateLimitCallback demonstrates how to adapt the rate limit dynamically.
func dynamicRateLimitCallback() int {
	// Example logic: Adjust rate limit based on current hour
	currentTime := time.Now()
	currentHour := currentTime.Hour()

	if currentHour >= 16 && currentHour < 18 {
		// Peak traffic hour: Increase rate limit
		return 20
	} else if currentHour < 7 {
		// Overnight: Decrease rate limit
		return 5
	} else {
		// Default rate limit
		return 10
	}
}

func main() {
	// Initialize the rate limiter with an initial rate of 10 requests per second
	rateLimiter := NewRateLimiter(10, dynamicRateLimitCallback)

	go func() {
		for i := 0; ; i++ {
			if rateLimiter.Allow() {
				fmt.Println("Request allowed.")
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			} else {
				fmt.Println("Request denied.")
			}
		}
	}()

	// Let the goroutine run for some time and simulate traffic changes by calling the callback function after 2 seconds.
	time.Sleep(2 * time.Second)
	fmt.Println("Simulating traffic changes...")
}
