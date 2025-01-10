package main

import (
	"fmt"
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

	if entry, ok := c.entries[key]; ok {
		if entry.expires.Before(time.Now()) {
			c.delete(key)
			return nil, false
		}

		// Update the expiration time
		entry.expires = time.Now().Add(entry.config.Period)
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

func main() {
	cache := NewCache(time.Second * 10)

	// Set a configuration for a key
	cache.set("key1", &Config{Limit: 10, Period: time.Second * 5})

	// Retrieve the configuration
	config, exists := cache.get("key1")
	if exists {
		fmt.Printf("Retrieved config for key1: %+v\n", config)
	} else {
		fmt.Printf("Config for key1 not found or expired\n")
	}

	// Simulate time passing
	time.Sleep(time.Second * 6)

	// Retrieve the configuration again
	config, exists = cache.get("key1")
	if exists {
		fmt.Printf("Retrieved config for key1: %+v\n", config)
	} else {
		fmt.Printf("Config for key1 not found or expired\n")
	}
}
