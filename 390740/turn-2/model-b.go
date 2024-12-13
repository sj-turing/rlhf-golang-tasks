package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Simulated resource (e.g., a database connection or a file)
type Resource struct {
	id int
}

func newResource() *Resource {
	return &Resource{id: 1}
}

func (r *Resource) Close() error {
	log.Printf("Closing resource with ID: %d\n", r.id)
	return nil
}

// Handler function that might panic
func panicHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request...")
	// Simulating a panic
	if someConditionCausesPanic() {
		panic("something went wrong!")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

// Function simulating a condition that might cause a panic
func someConditionCausesPanic() bool {
	// Randomly return true or false for demonstration purposes
	return false // Change to true to simulate a panic
}

// Middleware to catch panics
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v\n", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a resource
	resource := newResource()
	defer resource.Close() // Ensure resource is closed on program exit

	// Create a server and router
	router := http.NewServeMux()
	router.Handle("/", recoveryMiddleware(http.HandlerFunc(panicHandler)))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown setup
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		log.Println("Shutting down server...")

		// Shutdown the server gracefully
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shutdown server: %v\n", err)
		}
		log.Println("Server shut down successfully.")
	}()

	// Handle OS signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	<-sigs

	// Cancel the context to trigger the shutdown
	cancel()

	// Wait for the server to shut down
	time.Sleep(2 * time.Second)
	log.Println("Exiting program.")
}
