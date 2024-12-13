package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Handler function that simulates an HTTP call processing
func requestHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context that will timeout after 2 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel() // Important to ensure resources are cleaned up

	// Simulate processing, for example, making an HTTP call
	select {
	case <-time.After(3 * time.Second): // Simulating a long-running operation
		fmt.Fprintf(w, "Data processed successfully")
	case <-ctx.Done(): // This will trigger if the processing takes too long
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		return // End early so we don't write a response
	}
}

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
