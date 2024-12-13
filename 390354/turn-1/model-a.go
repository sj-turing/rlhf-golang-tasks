package main

/*

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

// requestHandler handles incoming TCP requests
func requestHandler(conn net.Conn, timeout time.Duration) {
	defer conn.Close()

	// Create a context with a timeout for this request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Simulate processing the request
	done := make(chan struct{})

	go func() {
		// Simulate some work (e.g., database call, API request)
		time.Sleep(2 * time.Second) // Simulate work taking 2 seconds
		fmt.Fprintln(conn, "Request processed successfully")
		close(done)
	}()

	// Wait for processing to finish or context timeout
	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Fprintln(conn, "Request timed out")
		}
	case <-done:
		// Request completed before timeout
	}
}

func main() {
	// Set up TCP listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	// Handle OS interrupts gracefully
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	fmt.Println("Server listening on :8080")
	go func() {
		<-sigs
		fmt.Println("Shutting down server...")
		listener.Close()
	}()

	// Accept connections continuously
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each request in a new goroutine
		go requestHandler(conn, 1*time.Second) // 1 second timeout for each request
	}
}

*/
