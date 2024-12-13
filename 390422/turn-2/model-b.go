package main

/*

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

// Simulate a work function that may encounter various errors
func simulateWork(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate an I/O error (network error)
	select {
	case <-time.After(time.Second):
		// Simulate successful operation
		fmt.Printf("Work completed for id: %d\n", id)
	default:
		// Simulate network timeout
		err := fmt.Errorf("network timeout for id: %d", id)
		handleError(ctx, err)
	}
}

// Handle different types of errors
func handleError(ctx context.Context, err error) {
	select {
	case <-ctx.Done():
		fmt.Printf("Error handling cancelled: %v\n", err)
	default:
		switch err {
		case context.Canceled:
			fmt.Println("Context was cancelled")
		case net.ErrClosed:
			fmt.Println("Network connection closed")
		case time.ErrTimeout:
			fmt.Println("Operation timed out")
		default:
			fmt.Printf("An error occurred: %v\n", err)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go simulateWork(ctx, i, &wg)
	}

	wg.Wait()
}
*/
