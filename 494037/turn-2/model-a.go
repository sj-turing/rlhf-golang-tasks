package main

/*

import (
	"fmt"
	"sync"
	"time"
)

// Simulated data aggregation task
func aggregateData(id int, data []int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal the completion of this goroutine to the WaitGroup
	sum := 0
	for _, value := range data {
		sum += value
	}
	fmt.Printf("Task %d: Aggregated sum = %d\n", id, sum)
}

// Simulate event handling
func handleEvent(data [][]int, numWorkers int) {
	var wg sync.WaitGroup

	// Create a channel to collect results (optional, for synchronization)
	results := make(chan int)

	// Launch parallel aggregation tasks
	for id, subData := range data {
		wg.Add(1) // Increment the WaitGroup count for this task
		go func(id int, subData []int) {
			defer wg.Done() // Decrement the WaitGroup count when this task finishes
			sum := 0
			for _, value := range subData {
				sum += value
			}
			// Optionally send result through the channel
			results <- sum
			fmt.Printf("Task %d: Aggregated sum = %d\n", id, sum)
		}(id, subData)
	}

	// Wait for all tasks to complete
	wg.Wait()      // Main thread waits here until all WaitGroup counts are zero
	close(results) // Close the channel to signal no more results

	// Optionally process results from the channel
	for result := range results {
		fmt.Printf("Processing result: %d\n", result)
	}

	fmt.Println("All aggregation tasks completed.")
}

func main() {
	// Simulated data
	data := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	// Simulate an event that triggers data aggregation
	handleEvent(data, 3) // You can specify the number of worker goroutines

	// Simulate another event after a delay
	time.Sleep(2 * time.Second)
	handleEvent(data, 3)
}
*/
