package main

/*

import (
	"fmt"
	"net/http"
	"time"
)

// RateLimiter struct to limit the number of concurrent requests
type RateLimiter struct {
	maxConcurrent int
	ch            chan struct{}
}

func NewRateLimiter(maxConcurrent int) *RateLimiter {
	return &RateLimiter{
		maxConcurrent: maxConcurrent,
		ch:            make(chan struct{}, maxConcurrent),
	}
}

// Acquire blocks until a slot is available
func (rl *RateLimiter) Acquire() {
	rl.ch <- struct{}{}
}

// Release frees a slot
func (rl *RateLimiter) Release() {
	<-rl.ch
}

func handler(w http.ResponseWriter, r *http.Request, limiter *RateLimiter) {
	limiter.Acquire()
	defer limiter.Release()

	fmt.Fprintln(w, "Handling request...")

	// Simulate some work
	time.Sleep(time.Millisecond * 200)
}

func main() {
	limiter := NewRateLimiter(2) // Allow only 2 concurrent requests

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, limiter)
	})

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

*/
