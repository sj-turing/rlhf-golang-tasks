package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RateLimiter struct {
	max    int
	window time.Duration
	limit  int
	count  int
	last   time.Time
	mtx    sync.Mutex
}

func NewRateLimiter(max int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		max:    max,
		window: window,
		limit:  max,
		count:  0,
		last:   time.Now(),
		mtx:    sync.Mutex{},
	}
}

func (rl *RateLimiter) Limit() {
	rl.mtx.Lock()
	defer rl.mtx.Unlock()

	elapsed := time.Since(rl.last)

	if elapsed >= rl.window {
		rl.last = time.Now()
		rl.count = 0
	}

	if rl.count >= rl.limit {
		since := time.Since(rl.last)
		sleep := rl.window - since
		time.Sleep(sleep)
		rl.last = time.Now()
	}

	rl.count++
}

// Simulated database operation
func simulateDatabaseOperation(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Database operation completed")
}

// Rate limiter middleware
func rateLimitMiddleware(next http.Handler, rateLimiter *RateLimiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateLimiter.Limit()
		next.ServeHTTP(w, r)
	})
}

// HTTP handler to process requests
func processRequest(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	maxRequests := 10             // Define the rate limit
	windowTime := 5 * time.Second // Define the rate limit window

	rl := NewRateLimiter(maxRequests, windowTime)

	start := time.Now()

	// Allow 5 concurrent requests
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go simulateDatabaseOperation(wg)
	}

	wg.Wait()

	end := time.Now()

	fmt.Fprintf(w, "All operations completed in %v\n", end.Sub(start))
}

func main() {
	maxRequests := 10
	windowTime := 5 * time.Second
	rl := NewRateLimiter(maxRequests, windowTime)

	http.Handle("/process", rateLimitMiddleware(http.HandlerFunc(processRequest), rl))

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
