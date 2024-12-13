package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

// Function simulating work that could fail
func doWork(ctx context.Context) {
	defer cleanup() // Ensure cleanup is called

	// Simulated "work" that might panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Received cancellation signal, exiting work.")
			return
		default:
			// Simulating a potential panic situation
			if someConditionCausesPanic() {
				panic("something went wrong!")
			}
			log.Println("Doing work...")
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}
}

// Cleanup function to release resources
func cleanup() {
	log.Println("Cleaning up resources...")
	// Place cleanup code here (e.g., closing files, network connections, etc.)
}

// Function simulating a condition that might cause a panic
func someConditionCausesPanic() bool {
	// Randomly return true or false for demonstration purposes
	return false // Change to true to simulate a panic
}

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// Start work in a Goroutine
	go doWork(ctx)

	// Wait for a termination signal
	<-sigs // Block until we receive a signal

	// Cancel the context to signal Goroutines to clean up before exit
	log.Println("Received interrupt signal, cancelling context.")
	cancel()

	// Give some time for cleanup
	time.Sleep(2 * time.Second) // Wait for cleaning up before exiting
	log.Println("Exiting program.")
}
