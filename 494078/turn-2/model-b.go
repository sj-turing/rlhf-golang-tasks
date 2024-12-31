package main

/*

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Simulated database operation
func simulateDatabaseOperation(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Database operation completed")
}

// RateLimiter middleware
func RateLimiter(limiter *rate.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := limiter.Wait(r.Context()); err != nil {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// HTTP handler to process requests
func processRequest(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	numOperations := 5 // Number of concurrent database operations

	start := time.Now()

	// Spawn multiple goroutines to simulate concurrent database operations
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go simulateDatabaseOperation(wg)
	}

	// Wait for all database operations to complete
	wg.Wait()

	end := time.Now()

	// Respond to the HTTP request
	fmt.Fprintf(w, "All %d operations completed in %v\n", numOperations, end.Sub(start))
}

func main() {
	// Create a Rate Limiter with a burst limit of 10 requests per second
	limiter := rate.NewLimiter(rate.Every(time.Second), 10)

	http.HandleFunc("/process", RateLimiter(limiter)(processRequest))

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
*/
