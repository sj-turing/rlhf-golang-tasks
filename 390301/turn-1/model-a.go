package main

/*

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Shared data structure
type Counter struct {
	mu    sync.Mutex
	count int
}

// Method to increment the counter safely
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Method to get the current count safely
func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// Global counter instance
var counter Counter

// HTTP handler to increment the counter
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	counter.Increment()
	fmt.Fprintf(w, "Counter incremented. Current count: %d\n", counter.GetCount())
}

// HTTP handler to get the current count
func getCountHandler(w http.ResponseWriter, r *http.Request) {
	count := counter.GetCount()
	fmt.Fprintf(w, "Current count: %d\n", count)
}

func main() {
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/count", getCountHandler)

	// Start the HTTP server
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Failed to start server:", err)
		}
	}()

	// Simulate concurrent requests
	for i := 0; i < 10; i++ {
		go func() {
			http.Get("http://localhost:8080/increment")
		}()
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to finish

	// Get the final count
	resp, err := http.Get("http://localhost:8080/count")
	if err != nil {
		fmt.Println("Error getting count:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Final count response status:", resp.Status)
}
*/
