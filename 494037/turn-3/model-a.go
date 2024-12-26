package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Simulate data from a database or API
func fetchData() []int {
	// In a real application, this would be a database query or API call
	data := make([]int, 100)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	return data
}

// Simulated data aggregation task
func aggregateData(id int, data []int, aggData *AggregatedData, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	count := 0

	for _, value := range data {
		sum += value
		count++
	}

	// Update the aggregated data structure using atomic operations to prevent race conditions
	atomic.AddInt64(&aggData.Count, int64(count))
	atomic.AddInt64(&aggData.Sum, int64(sum))

	fmt.Printf("Task %d: Aggregated count = %d, sum = %d\n", id, count, sum)
}

type AggregatedData struct {
	Count int64
	Sum   int64
}

func main() {
	const numTasks = 5
	var wg sync.WaitGroup
	aggregatedData := &AggregatedData{}

	// Launch parallel data fetching and aggregation tasks
	for id := 0; id < numTasks; id++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			// Fetch data (simulated)
			data := fetchData()
			// Aggregate data in a separate goroutine
			aggregateData(taskID, data, aggregatedData, &wg)
		}(id)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Calculate and print the final aggregated result
	fmt.Println("--- Final Aggregated Result ---")
	fmt.Println("Final Count:", aggregatedData.Count)
	fmt.Println("Final Sum:  ", aggregatedData.Sum)
}
