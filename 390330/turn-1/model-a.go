package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Simulate a long-running operation
func longRunningOperation(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // Simulate a long operation
		fmt.Println("Operation completed successfully")
	case <-ctx.Done(): // Listen for cancellation
		fmt.Println("Operation cancelled:", ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request

	// Start the long-running operation
	go longRunningOperation(ctx)

	// Simulate a delay in responding to the client
	time.Sleep(3 * time.Second)

	// Respond to the client
	w.Write([]byte("Request processed"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
