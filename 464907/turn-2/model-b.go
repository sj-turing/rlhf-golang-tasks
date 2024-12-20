package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a custom error type for task failures
type taskError struct {
	message string
}

func (e *taskError) Error() string {
	return fmt.Sprintf("Task failed: %s", e.message)
}

func asyncTask(done chan<- error) {
	// Simulate a long-running task that may fail
	time.Sleep(2 * time.Second)
	if randomBool() {
		done <- nil // Task completed successfully
	} else {
		done <- &taskError{"Async task failed!"}
	}
}

func asyncTaskWithWaitGroup(wg *sync.WaitGroup, done chan<- error) {
	defer wg.Done()
	asyncTask(done)
}

func syncTask() error {
	// Simulate a long-running task that may fail
	time.Sleep(2 * time.Second)
	if randomBool() {
		return nil // Task completed successfully
	} else {
		return &taskError{"Sync task failed!"}
	}
}

func randomBool() bool {
	return time.Now().UnixNano()%2 == 0
}

func main() {
	// Scenario: Asynchronous callback - Preferred when tasks can be non-blocking
	// and you want to handle multiple tasks concurrently.
	done := make(chan error, 1)
	fmt.Println("Starting async task...")
	go asyncTask(done)

	// Do other tasks while waiting for the async task to complete
	fmt.Println("Doing other tasks...")

	// Wait for the asynchronous task to finish and handle the error
	err := <-done
	if err != nil {
		fmt.Println("Error in async task:", err)
	} else {
		fmt.Println("Async task completed successfully.")
	}

	// Scenario: Synchronous callback - Preferred when tasks are dependent
	// on the previous one's output, and blocking the main thread is acceptable.
	fmt.Println("Starting sync task...")
	err = syncTask()
	if err != nil {
		fmt.Println("Error in sync task:", err)
	} else {
		fmt.Println("Sync task completed successfully.")
	}

	// Scenario: Using WaitGroup for asynchronous callbacks - useful when you want to wait for multiple
	// asynchronous tasks to finish before proceeding with the main thread.
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go asyncTaskWithWaitGroup(wg, done)
	go asyncTaskWithWaitGroup(wg, done)

	wg.Wait() // Wait for both tasks to complete
	fmt.Println("All asynchronous tasks completed.")
}
