package main

/*

import (
	"fmt"
	"net/http"
)

// Shared data structure
type Counter struct {
	value int
}

// Channel for sending increment requests
type Request struct {
	increment bool
	response  chan int
}

// Function that processes requests to the counter
func (c *Counter) ProcessRequests(requests chan Request) {
	for req := range requests {
		if req.increment {
			c.value++
		}
		// Send the current value back through the response channel
		req.response <- c.value
	}
}

// HTTP handler to increment the counter
func incrementHandler(requests chan Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseChan := make(chan int)
		req := Request{increment: true, response: responseChan}
		requests <- req // Send request to the processing goroutine

		// Wait for the response
		newValue := <-responseChan
		fmt.Fprintf(w, "Counter incremented. Current count: %d\n", newValue)
	}
}

// HTTP handler to get the current count
func getCountHandler(requests chan Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseChan := make(chan int)
		req := Request{increment: false, response: responseChan}
		requests <- req // Send request to the processing goroutine

		// Wait for the response
		currentValue := <-responseChan
		fmt.Fprintf(w, "Current count: %d\n", currentValue)
	}
}

func main() {
	// Create a channel for requests
	requests := make(chan Request)

	// Create a counter instance
	counter := Counter{}

	// Start a goroutine to process requests
	go counter.ProcessRequests(requests)

	// Set up HTTP handlers
	http.HandleFunc("/increment", incrementHandler(requests))
	http.HandleFunc("/count", getCountHandler(requests))

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
*/
