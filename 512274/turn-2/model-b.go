package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Config represents the rate limiter configuration for each API key
type Config struct {
	Limit  int           `json:"limit"`
	Period time.Duration `json:"period"`
}

// Cache is the main structure for managing rate limiter configurations
type Cache struct {
	mu            sync.RWMutex
	entries       map[string]*entry
	expireChannel chan struct{} // channel to trigger expiration checks
}

// entry holds the configuration and its expiration time
type entry struct {
	config  Config
	expires time.Time
}

func main() {
	cache := NewCache(time.Second * 5) // Expire after 5 seconds of inactivity

	rateLimitMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract API key from the request header (e.g., 'Authorization: Bearer <key>')
			apiKey := r.Header.Get("Authorization")
			if apiKey == "" {
				http.Error(w, "Invalid or missing API key", http.StatusUnauthorized)
				return
			}

			if err := checkRateLimit(cache, apiKey); err != nil {
				http.Error(w, err.Error(), http.StatusTooManyRequests)
				return
			}

			// Rate limit not exceeded, proceed with the request
			next.ServeHTTP(w, r)
		})
	}

	http.HandleFunc("/", rateLimitMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the API!")
	})))

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func checkRateLimit(cache *Cache, apiKey string) error {
	config, exists := cache.get(apiKey)
	if !exists {
		// API key not found in the cache
		return fmt.Errorf("invalid API key")
	}

	// Simulate a request and update the rate limit counter (you would use an actual implementation here)
	// For demonstration purposes, let's increment a counter in the cache entry.
	// In a real application, you would use an in-memory store like Redis or distributed counters.
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cache.entries[apiKey]
	if entry.expires.Before(time.Now()) {
		// Rate limit period has expired, reset the counter
		entry.expires = time.Now().Add(config.Period)
	}
	cache.entries[apiKey] = entry

	// Implement your actual rate limit logic here, based on your use case
	// For this basic example, let's just check if the counter has exceeded the limit
	// In a real implementation, you might consider time-based or sliding-window rate limiting.
	if entry.config.Limit < 1 {
		return fmt.Errorf("rate limit exceeded")
	}

	return nil
}

// NewCache initializes a new Cache with a default expiration check interval
func NewCache(expireInterval time.Duration) *Cache {
	cache := &Cache{
		entries:       make(map[string]*entry),
		expireChannel: make(chan struct{}),
	}

	go cache.expireEntries(expireInterval)
	return cache
}

// get retrieves the configuration for a given key, updating the expiration time
func (c *Cache) get(key string) (*Config, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// ... (existing code)
}

// set stores or updates the configuration for a given key
func (c *Cache) set(key string, config *Config) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// ... (existing code)
}

// delete removes the configuration for a given key
func (c *Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// ... (existing code)
}
