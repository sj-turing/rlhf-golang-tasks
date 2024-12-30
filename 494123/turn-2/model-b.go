package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/oschwald/geoip2-golang"
)

// rateLimiter represents the rate limiter for a single client and country.
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

// geoIPDatabase is the GeoIP database used to determine client country.
var geoIPDatabase *geoip2.Reader

// initGeoIPDatabase initializes the GeoIP database for country determination.
func initGeoIPDatabase() error {
	var err error
	geoIPDatabase, err = geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		return fmt.Errorf("failed to open GeoIP2 database: %w", err)
	}
	return nil
}

// getClientCountry retrieves the country of the client based on its IP address.
func getClientCountry(client string) (string, error) {
	ip := net.ParseIP(client)
	if ip == nil {
		return "", fmt.Errorf("invalid client IP address: %q", client)
	}

	record, err := geoIPDatabase.Country(ip)
	if err != nil {
		return "", fmt.Errorf("failed to get GeoIP2 record for IP: %w", err)
	}

	return record.Country.ISO_Code, nil
}

// updateRateLimit updates the rate limit for the specified client and country.
func updateRateLimit(client string, country string, rate int, duration time.Duration) {
	key := client + "+" + country

	rl, ok := rateLimitSyncMap.Load(key)
	if !ok {
		rl = newRateLimiter(rate, duration)
		rateLimitSyncMap.Store(key, rl)
	} else {
		rl, ok := rl.(*rateLimiter)
		if !ok {
			// If failed to get typecast object you may take necessary actions here
			return
		}
		rateLimitMutex.Lock()
		rl.rate = rate
		rl.duration = duration
		rateLimitMutex.Unlock()
	}
}

// evictStaleData removes rate limiters for clients that have not made any requests in the past duration.
