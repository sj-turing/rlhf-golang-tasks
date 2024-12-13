package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Simulates a long-running operation
func slowOperation(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	defer w.Close()

	// Check if the context is cancelled before starting the operation
	select {
	case <-ctx.Done():
		// Operation cancelled, return early
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Operation cancelled: %v\n", ctx.Err())
		return
	default:
	}

	// Simulate a 5-second operation
	log.Printf("Starting slow operation...")
	time.Sleep(5 * time.Second)
	log.Printf("Slow operation completed.")

	fmt.Fprintf(w, "Operation completed successfully.\n")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a context with a 2-second timeout
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// Start the slow operation in a new goroutine
		go slowOperation(ctx, w, r)

		// Wait for the operation to complete or the context to be cancelled
		select {
		case <-ctx.Done():
			// Context was cancelled or expired
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Request timed out: %v\n", ctx.Err())
		default:
		}
	})

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
