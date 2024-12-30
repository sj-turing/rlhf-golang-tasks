package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	maxRequestsPerMinute = 60 // Example rate limit
	requestDuration      = time.Minute
)

var (
	clients         map[string]*clientData
	clientsLock     sync.RWMutex
	rateLimiterCh   chan struct{} // Channel to control concurrency
	rateLimiterSize = 100         // Example concurrency limit
)

func main() {
	rateLimiterCh = make(chan struct{}, rateLimiterSize)

	http.HandleFunc("/", rateLimitedHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func rateLimitedHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve client IP
	clientIP := r.RemoteAddr

	// Control concurrency using the rate limiter channel
	<-rateLimiterCh
	defer func() {
		rateLimiterCh <- struct{}{}
	}()

	// Acquire lock for reading and writing to clients map
	clientsLock.RLock()
	clientData, exists := clients[clientIP]
	clientsLock.RUnlock()

	// If client not in map, add with initial data
	if !exists {
		clientData = newClientData()
		clientsLock.Lock()
		clients[clientIP] = clientData
		clientsLock.Unlock()
	}

	// Check if client has exceeded rate limit
	if !clientData.canMakeRequest(maxRequestsPerMinute) {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	// Update request count
	clientData.updateRequestCount()

	// Send response
	fmt.Fprintf(w, "Hello, %s! You've made %d requests this minute.", clientIP, clientData.requestCount)
}

func newClientData() *clientData {
	return &clientData{
		lastRequestTime: time.Now(),
		requestCount:    0,
	}
}

type clientData struct {
	lastRequestTime time.Time
	requestCount    int
}

func (cd *clientData) canMakeRequest(maxRequestsPerMinute int) bool {
	// Calculate window
	windowDuration := time.Since(cd.lastRequestTime)

	// If window has expired, reset
	if windowDuration >= requestDuration {
		cd.lastRequestTime = time.Now()
		cd.requestCount = 0
		return true
	}

	// Check if request count is within limit
	return cd.requestCount < maxRequestsPerMinute
}

func (cd *clientData) updateRequestCount() {
	cd.requestCount++
	cd.lastRequestTime = time.Now()
}

// Function to handle stale data
func evictStaleData() {
	for {
		clientsLock.RLock()
		for clientIP, clientData := range clients {
			if time.Since(clientData.lastRequestTime) > requestDuration {
				clientsLock.RUnlock()
				clientsLock.Lock()
				delete(clients, clientIP)
				clientsLock.Unlock()
				break
			}
		}
		clientsLock.RUnlock()

		time.Sleep(10 * time.Second) // Check for stale data every 10 seconds
	}
}

func init() {
	// Start a goroutine to evict stale data
	go evictStaleData()
}
