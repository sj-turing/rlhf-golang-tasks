package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// Simulate a running resource (e.g., a worker goroutine)
func simulateWork(id int, shutdown chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-shutdown:
			log.Printf("Goroutine %d received shutdown signal, exiting.", id)
			return
		default:
			log.Printf("Goroutine %d working...", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Define waitgroup to manage goroutines
	var wg sync.WaitGroup

	// Channel for graceful shutdown of goroutines
	shutdown := make(chan<- bool)

	// Create and start worker goroutines
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go simulateWork(i, shutdown, &wg)
	}

	// Create a new HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", panicHandler(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a panic
		panic("Unexpected condition occurred!")
	}))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Listen for panics
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic: %v", err)
			// Ensure goroutines are stopped on panic
			close(shutdown)
			wg.Wait()  // Wait for all goroutines to complete
			os.Exit(1) // Exit the application
		}
	}()

	log.Println("Starting HTTP server on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}

// Panicking handler wrapper
func panicHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
				log.Printf("Recovered from panic: %v", err)
			}
		}()
		f(w, r)
	}
}
