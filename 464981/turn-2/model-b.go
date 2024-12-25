package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Simple function that accepts an asynchronous callback
func processNumberAsync(num int, callback func(int, error)) {
	go func() {
		// Simulate some work
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		// Call the callback
		callback(num*num, nil)
	}()
}

// Example callback function that uses atomic counters
func logResultWithAtomicCounter(result int, err error) {
	var activeCallbacks int32
	// Increment the counter before handling the callback
	atomic.AddInt32(&activeCallbacks, 1)
	defer func() {
		// Decrement the counter after handling the callback
		atomic.AddInt32(&activeCallbacks, -1)
	}()
	if err != nil {
		log.Printf("Callback Error: %v", err)
		return
	}
	log.Printf("Callback Result: %d (Active Callbacks: %d)", result, activeCallbacks)
}
func main() {
	fmt.Println("Starting program")
	// Initialize a wait group to wait for all callbacks to complete
	var wg sync.WaitGroup
	// Set the number of concurrent callbacks
	const numCallbacks = 10
	wg.Add(numCallbacks)
	for i := 0; i < numCallbacks; i++ {
		processNumberAsync(i+1, func(result int, err error) {
			logResultWithAtomicCounter(result, err)
			wg.Done()
		})
	}
	// Wait for all callbacks to finish
	wg.Wait()
	fmt.Println("Program finished")
}
