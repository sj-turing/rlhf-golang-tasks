package main

import (
	"fmt"
	"sync"
	"time"
)

// SharedCounter is a struct to hold a counter and a mutex to protect it
type SharedCounter struct {
	count int
	mutex sync.RWMutex
}

// Increment increments the counter safely by acquiring the mutex
func (sc *SharedCounter) Increment() {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	sc.count++
}

// Decrement decrements the counter safely by acquiring the mutex
func (sc *SharedCounter) Decrement() {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	sc.count--
}

// Get returns the current count
func (sc *SharedCounter) Get() int {
	sc.mutex.RLock() // Use RLock since we're only reading
	defer sc.mutex.RUnlock()
	return sc.count
}

func main() {
	counter := &SharedCounter{}

	// Create multiple goroutines to increment and decrement the counter concurrently
	numGoroutines := 100
	numOperations := 10000

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numOperations; j++ {
				if j%2 == 0 {
					counter.Increment()
				} else {
					counter.Decrement()
				}
			}
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(2 * time.Second)

	// Print the final count
	fmt.Printf("Final count: %d\n", counter.Get())
}
