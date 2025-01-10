package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// UserConfig represents the rate limiter configuration
type UserConfig struct {
	MaxRequestAllowed int32
	ExpiryWindow      time.Time
}

// Cache is the main structure for managing rate limiter configurations
type Cache struct {
	mu      sync.RWMutex
	records map[string]*UserConfig
}

// NewCache initializes a new Cache with a default expiration check interval
func NewCache(expireInterval time.Duration) *Cache {
	cache := &Cache{
		records: make(map[string]*UserConfig),
	}

	go cache.expireEntries(expireInterval)
	return cache
}

// get retrieves the configuration for a given key, updating the expiration time and hit count
func (c *Cache) get(key string) (*UserConfig, error) {
	c.mu.RLock()
	record, ok := c.records[key]
	c.mu.RUnlock()

	if ok {
		if record.MaxRequestAllowed <= 1 || record.ExpiryWindow.Before(time.Now()) {
			return nil, fmt.Errorf("Limit exhuasted")
		}

		record.ExpiryWindow = time.Now().Add(5 * time.Second)
		atomic.AddInt32(&record.MaxRequestAllowed, -1)
		c.set(key, record)
		return record, nil

	}

	newRecord := &UserConfig{
		MaxRequestAllowed: 3,
		ExpiryWindow:      time.Now().Add(5 * time.Second),
	}
	c.set(key, newRecord)

	return record, nil
}

// set stores or updates the configuration for a given key
func (c *Cache) set(key string, userConfig *UserConfig) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.records[key] = userConfig
}

// expireEntries periodically checks for expired entries and removes them
func (c *Cache) expireEntries(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("cleaning up the cache")
		c.mu.Lock()
		for key, entry := range c.records {
			if entry.ExpiryWindow.Before(time.Now()) {
				delete(c.records, key)
			}
		}
		c.mu.Unlock()
	}
}

func main() {
	cache := NewCache(time.Second * 10) // Expire after 10 seconds of inactivity

	rateLimitMiddleware := func(next http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Extract API key from the request header (e.g., 'Authorization: Bearer <key>')
			apiKey := r.Header.Get("Authorization")
			if apiKey == "" {
				http.Error(w, "Invalid or missing API key", http.StatusUnauthorized)
				return
			}

			if err := checkRateLimit(cache, apiKey); err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusTooManyRequests)
				return
			}

			// Rate limit not exceeded, proceed with the request
			next.ServeHTTP(w, r)
		}
	}

	http.HandleFunc("/", rateLimitMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Welcome to the API!")
		fmt.Fprintf(w, "Welcome to the API!")
	})))

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func checkRateLimit(cache *Cache, apiKey string) error {
	// API key not found in the cache
	_, err := cache.get(apiKey)
	return err
}
