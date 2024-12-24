package main

import (
	"464944/turn-3/model-a/dispose"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Example of an external resource that needs cleanup
type FileCloser struct {
	*os.File
}

func (fc *FileCloser) Close() error {
	return fc.File.Close()
}

type DatabaseConnection struct {
	// Simulate database connection state
	isConnected bool
}

func (dc *DatabaseConnection) Close() error {
	dc.isConnected = false
	return nil
}

func (dc *DatabaseConnection) IsConnected() bool {
	return dc.isConnected
}

// Example usage of the dispose package
func main() {
	// Create a dispose instance
	disposer := dispose.New()

	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	fc := &FileCloser{file}
	disposer.Register(fc)

	// Connect to a database
	db := &DatabaseConnection{isConnected: true}
	disposer.Register(db)

	// Register a custom resource for cleanup
	type CustomResource struct{}
	customResource := &CustomResource{}
	disposer.Register(customResource)

	// Simulate some work
	fmt.Println("Working with resources...")
	time.Sleep(2 * time.Second)

	// Signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("Received shutdown signal, disposing resources...")

		// Dispose all registered resources
		if err := disposer.Dispose(); err != nil {
			fmt.Println("Error disposing resources:", err)
		} else {
			fmt.Println("Resources disposed successfully.")
		}
		os.Exit(0)
	}()

	// Simulate more work
	time.Sleep(5 * time.Second)

	// If no signal is received, dispose resources manually
	fmt.Println("Shutting down manually, disposing resources...")
	if err := disposer.Dispose(); err != nil {
		fmt.Println("Error disposing resources:", err)
	} else {
		fmt.Println("Resources disposed successfully.")
	}
}
