package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asyncTask(done chan bool, errChan chan error) {
	// Simulate a long-running task with potential error
	time.Sleep(2 * time.Second)
	if err := someSimulatedError(); err != nil {
		errChan <- err
		return
	}
	fmt.Println("Async task completed without error.")
	done <- true
}

func someSimulatedError() error {
	if random := randInt(0, 10); random < 5 {
		return fmt.Errorf("Simulated error: %d", random)
	}
	return nil
}

func randInt(min, max int) int {
	return min + int(rand.Intn(max-min+1))
}

func main() {
	done := make(chan bool)
	errChan := make(chan error)

	fmt.Println("Starting async task...")
	go asyncTask(done, errChan)

	// Do other tasks while waiting for the async task to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Doing other tasks...")

	// Wait for async task to complete
	<-done

	// Handle potential errors
	if err := <-errChan; err != nil {
		fmt.Println("Error from async task:", err)
	} else {
		fmt.Println("Main thread terminated without error.")
	}
}
