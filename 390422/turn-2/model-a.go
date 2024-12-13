package main

import (
	"fmt"
	"sync"
	"time"
)

// Custom error types
type ErrNilPointer struct {
	message string
}

func (e *ErrNilPointer) Error() string {
	return e.message
}

type ErrOutOfRange struct {
	index int
}

func (e *ErrOutOfRange) Error() string {
	return fmt.Sprintf("index %d out of range", e.index)
}

// Function to simulate work that may cause different errors
func simulateWork(id int, results chan<- interface{}) {
	defer func() {
		if r := recover(); r != nil {
			results <- fmt.Errorf("panic recovered in Goroutine %d: %v", id, r)
		}
	}()

	// Simulate a nil pointer dereference for id == 1
	if id == 1 {
		var ptr *int
		_ = *ptr // This will cause a nil pointer dereference
	}

	// Simulate an out of range error for id == 2
	if id == 2 {
		arr := []int{1, 2, 3}
		results <- &ErrOutOfRange{index: 5} // Simulating out of range error
		return
	}

	// Simulate a timeout or a network error for id == 3
	if id == 3 {
		time.Sleep(2 * time.Second) // Simulating a long operation
		results <- fmt.Errorf("timeout error in Goroutine %d", id)
		return
	}

	// Simulate successful work
	results <- fmt.Sprintf("Goroutine %d completed successfully", id)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan interface{}, 10) // Buffered channel for results

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			simulateWork(id, results)
		}(i)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	close(results) // Close the channel after all Goroutines are done

	// Handle collected results
	for result := range results {
		switch err := result.(type) {
		case error:
			fmt.Println("Error:", err)
		case string:
			fmt.Println("Result:", result)
		default:
			fmt.Println("Unknown type:", result)
		}
	}
}
