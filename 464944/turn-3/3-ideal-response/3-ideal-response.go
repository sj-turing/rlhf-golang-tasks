package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"464944/turn-3/3-ideal-response/dispose"
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

// Register a custom resource for cleanup
type CustomResource struct{}

func (cr *CustomResource) Close() error {
	fmt.Println("Closing the customResource")
	return nil
}

type HttpServer struct {
	*http.Server
}

func (s *HttpServer) Close() error {
	return s.Server.Shutdown(nil)
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

	customResource := &CustomResource{}
	disposer.Register(customResource)

	// Simulate some work
	fmt.Println("Working with resources...")
	time.Sleep(200 * time.Millisecond)

	ticker := time.NewTicker(5 * time.Second)

	// Signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start a simple HTTP server for example
	server := &HttpServer{&http.Server{Addr: ":8080"}}
	go func() {
		fmt.Println("HTTP server started on :8080")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()
	disposer.Register(server)

	select {
	case <-sigChan:
		fmt.Println("Received shutdown signal, disposing resources...")
	case <-ticker.C:
		fmt.Println("Received ticker signal, disposing resources...")
	}

	if err := disposer.Dispose(); err != nil {
		fmt.Println("Error disposing resources:", err)
	} else {
		fmt.Println("Resources disposed successfully.")
	}
	os.Exit(0)
}
