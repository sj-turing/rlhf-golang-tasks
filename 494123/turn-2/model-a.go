package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Define the rateLimiter struct to hold the rate limiting data for a country.
type rateLimiter struct {
	countries   sync.Map // Map of IP to time of last request
	lastUpdated time.Time
	rate        int // Maximum number of requests per time window
	duration    time.Duration
}

// Define a configuration for the rate limiting.
var config = map[string]*rateLimiter{
	"US": &rateLimiter{rate: 5, duration: time.Second * 10},
	"UK": &rateLimiter{rate: 10, duration: time.Second * 5},
}

// determineCountry determines the country for the request based on IP address.
// This is a placeholder function; in a real implementation, you would use an
// IP-to-country database or API.
func determineCountry(ip string) string {
	if ip[:2] == "12" {
		return "US"
	} else if ip[:2] == "13" {
		return "UK"
	} else {
		return "US"
	}
}

// checkRateLimit checks if the request from the given IP within the specified country is allowed.
func checkRateLimit(w http.ResponseWriter, r *http.Request) bool {
	ip := r.RemoteAddr.Split(":")[0] // Extract the IP address
	country := determineCountry(ip)
	rl := config[country]
	if rl == nil {
		log.Printf("No rate limit configuration for country: %s\n", country)
		return false
	}

	// Update the last update time and clear stale entries
	now := time.Now()
	rl.lastUpdated = now
	rl.countries.Range(func(k, v any) bool {
		if now.Sub(v.(time.Time)) > rl.duration {
			rl.countries.Delete(k)
		}
		return true
	})

	// Check if request is allowed
	if rl.countries.Load(ip) == nil {
		rl.countries.Store(ip, now)
		return true
	} else {
		return now.Sub(rl.countries.Load(ip).(time.Time)) > rl.duration
	}
}

// rateLimitHandler is the middleware to enforce rate limiting.
func rateLimitHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkRateLimit(w, r) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Define the routes for your web application.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start the web server with rate limiting enabled.
	log.Println("Starting web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", rateLimitHandler(http.DefaultServeMux)))
}
