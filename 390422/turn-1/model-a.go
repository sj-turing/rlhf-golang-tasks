package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

// simulateWork simulates a task that can fail
func simulateWork(id int) error {
	if id%2 == 0 {
		// Simulate an error for even IDs
		return fmt.Errorf("work failed for id: %d", id)
	}
	// Simulate successful work
	return nil
}

func main() {
	var wg sync.WaitGroup
	errorsChannel := make(chan error, 10) // Buffered channel for error collection

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := simulateWork(id)
			if err != nil {
				// Wrap the error with context and send it to the channel
				wrappedErr := errors.Wrap(err, fmt.Sprintf("Goroutine %d", id))
				errorsChannel <- wrappedErr
			}
		}(i)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	close(errorsChannel) // Close the channel after all Goroutines are done

	// Handle collected errors
	for err := range errorsChannel {
		fmt.Println(err)
	}
}
