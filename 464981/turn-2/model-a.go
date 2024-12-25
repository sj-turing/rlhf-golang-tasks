package main

/*

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// An asynchronous function that accepts a callback
func asyncProcess(input chan string, callback chan<- string, wg *sync.WaitGroup) {
	for val := range input {
		log.Println("Received input:", val)
		time.Sleep(time.Second) // Simulate work
		processed := val + "-processed"
		callback <- processed
		log.Println("Sending callback with:", processed)
		wg.Done()
	}
}

// A simple callback function
func printResult(result chan<- string) {
	for res := range result {
		fmt.Println("Callback result:", res)
	}
}

func main() {
	// Create a channel for input and a callback channel
	input := make(chan string)
	callback := make(chan string)

	// Create a wait group to wait for all goroutines
	var wg sync.WaitGroup

	// Start goroutine for asynchronous processing
	wg.Add(1)
	go asyncProcess(input, callback, &wg)

	// Start goroutine to handle callback results
	wg.Add(1)
	go printResult(callback)

	// Send inputs
	for i := 1; i <= 3; i++ {
		input <- fmt.Sprintf("Item %d", i)
	}

	// Close the input channel when done
	close(input)

	// Wait for all goroutines to complete
	wg.Wait()
}
*/
