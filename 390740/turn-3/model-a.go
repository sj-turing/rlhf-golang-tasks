package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Resource that needs to be cleaned up
type Resource struct {
	// example field
	ID int
}

func (r *Resource) Cleanup() {
	// Simulating resource cleanup
	log.Printf("Cleaning up resource with ID: %d", r.ID)
	// Here we would typically close connections, files, etc.
}

// Centralized panic handler for HTTP servers
func panicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
				log.Printf("Recovered from panic: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Example HTTP handler that may cause a panic
func riskyHandler(w http.ResponseWriter, r *http.Request) {
	// Simulated panic for demonstration
	panic("Simulated panic for demonstration")
}

func main() {
	// Create a new HTTP server
	mux := http.NewServeMux()
	mux.Handle("/", panicHandler(http.HandlerFunc(riskyHandler)))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel to listen for OS interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	resource := &Resource{ID: 1} // Example resource to clean up
	defer resource.Cleanup()     // Ensure resource cleanup on exit

	// Graceful shutdown handling
	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %s", sig)
		log.Println("Shutting down server...")

		// Create a deadline for the shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt to gracefully shut down the server
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
		log.Println("Server exiting gracefully.")
	}()

	log.Println("Starting HTTP server on :8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server listen error: %v", err)
	}
}

