package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Config represents the rate limiter configuration
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
	hits    int
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

// get retrieves the configuration for a given key, updating the expiration time and hit count
func (c *Cache) get(key string) (*Config, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if entry, ok := c.entries[key]; ok {
		if entry.expires.Before(time.Now()) {
			c.delete(key)
			return nil, false
		}

		// Update the expiration time and hit count
		entry.expires = time.Now().Add(entry.config.Period)
		entry.hits++
		if entry.hits > entry.config.Limit {
			return nil, false // Exceeded limit
		}

		return &entry.config, true
	}
	return nil, false
}

// set stores or updates the configuration for a given key
func (c *Cache) set(key string, config *Config) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := &entry{
		config:  *config,
		expires: time.Now().Add(config.Period),
		hits:    1,
	}

	c.entries[key] = entry
	select {
	case c.expireChannel <- struct{}{}:
	default:
	}
}

// delete removes the configuration for a given key
func (c *Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.entries, key)
}

// expireEntries periodically checks for expired entries and removes them
func (c *Cache) expireEntries(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.expires.Before(time.Now()) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

// RateLimitHandler checks the cache for rate limit and serves a request
func RateLimitHandler(w http.ResponseWriter, r *http.Request, cache *Cache) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	config, limitExceeded := cache.get(userID)
	if limitExceeded {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	// Simulate processing the request
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request processed successfully!"))
}

func main() {
	cache := NewCache(time.Second * 5)

	// Set a configuration for a user
	cache.set("user1", &Config{Limit: 3, Period: time.Second * 60})

	http.HandleFunc("/api/rate-limited", func(w http.ResponseWriter, r *http.Request) {
		RateLimitHandler(w, r, cache)
	})

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
