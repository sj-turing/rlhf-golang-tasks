package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Define a shared variable
	var counter int = 0

	// Define the function that will be executed in a Goroutine
	increment := func() {
		defer wg.Done() // Signal that this Goroutine is done
		for i := 0; i < 10000; i++ {
			counter++
		}
	}

	// Start 10 Goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go increment()
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Print the final counter value
	fmt.Println("Final counter value:", counter)
}
