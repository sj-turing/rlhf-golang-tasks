package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// rateLimiter represents the rate limiter for a single client.
type rateLimiter struct {
	rate     int           // Number of requests allowed per duration.
	duration time.Duration // Duration for rate limit.
	window   []time.Time   // Sliding window of request timestamps.
	mutex    sync.Mutex    // Mutex for concurrent access.
}

// newRateLimiter creates a new rateLimiter with the specified rate and duration.
func newRateLimiter(rate int, duration time.Duration) *rateLimiter {
	return &rateLimiter{
		rate:     rate,
		duration: duration,
		window:   make([]time.Time, 0, rate),
	}
}

// allow checks if a request is allowed based on the rate limit.
func (rl *rateLimiter) allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	// Calculate the duration since the earliest request in the window.
	if len(rl.window) == rl.rate {
		elapsed := time.Since(rl.window[0])
		if elapsed < rl.duration {
			return false
		}
		rl.window = rl.window[1:]
	}

	// Add the current timestamp to the window.
	rl.window = append(rl.window, time.Now())
	return true
}

// rateLimitSyncMap stores the rate limiters for each client.
var rateLimitSyncMap = sync.Map{}

// rateLimitMutex protects concurrent access to rateLimitMap.
var rateLimitMutex sync.Mutex

// updateRateLimit updates the rate limit for the specified client.
func updateRateLimit(client string, rate int, duration time.Duration) {

	rl, ok := rateLimitSyncMap.Load(client)
	if !ok {
		rl = newRateLimiter(rate, duration)
		rateLimitSyncMap.Store(client, rl)
	} else {
		rl, ok := rl.(*rateLimiter)
		if !ok {
			// if failed to get typecast object you may take necessary actions here
			return
		}
		rateLimitMutex.Lock()
		rl.rate = rate
		rl.duration = duration
		rateLimitMutex.Unlock()
	}
}

// evictStaleData removes rate limiters for clients that have not made any requests in the past duration.
func evictStaleData(duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		totalClients := 0
		deletedClients := 0
		rateLimitSyncMap.Range(func(key, value any) bool {
			client := key.(string)
			rl := value.(*rateLimiter)

			rateLimitMutex.Lock()
			totalClients++
			if len(rl.window) == 0 || time.Since(rl.window[0]) >= rl.duration {
				deletedClients++
				rateLimitSyncMap.Delete(client)
			}
			rateLimitMutex.Unlock()

			return true
		})

		fmt.Println("Total Clients: ", totalClients)
		fmt.Println("Total Clients after removing inactive clients:", totalClients-deletedClients)
	}
}

// rateLimitHandler is a middleware that handles rate limiting for HTTP requests.
func rateLimitHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := r.RemoteAddr // Replace this with a more appropriate way to identify the client.

		// Update the rate limit for the client if needed (e.g., from a configuration file or database).
		// For simplicity, we assume a fixed rate limit for all clients.
		updateRateLimit(client, 5, time.Second*10) // 5 requests per 10 seconds

		rl, ok := rateLimitSyncMap.Load(client)
		if !ok {
			// Client does not have a rate limit, proceed with the request.
			next.ServeHTTP(w, r)
			return
		}

		rateLimiterClient, ok := rl.(*rateLimiter)

		if !rateLimiterClient.allow() {
			// Request rate exceeded, respond with a 429 Too Many Requests error.
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		// Request allowed, proceed with the request.
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Start the rate limit data eviction routine.
	go evictStaleData(time.Second * 20)

	// Define the routes for your web application.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start the web server with rate limiting enabled.
	log.Println("Starting web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", rateLimitHandler(http.DefaultServeMux)))
}
