package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	numGoroutines := 5
	workerCount := 1000

	// Create a shared channel
	results := make(chan int, numGoroutines*workerCount)

	// Start the worker Goroutines using WaitGroup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(i, workerCount, results, &wg)
	}

	// Close the channel to signal worker Goroutines to stop
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process the results from the channel
	totalSum := 0
	for result := range results {
		totalSum += result
	}

	fmt.Println("Total sum of results:", totalSum)
}

func worker(id int, workCount int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < workCount; i++ {
		// Simulate some work
		time.Sleep(time.Duration(10) * time.Millisecond)
		// Send the result to the channel
		results <- i + 100*id
	}
}
