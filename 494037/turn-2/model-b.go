package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// A simple structure to hold aggregated data
type AggregatedData struct {
	Count int
	Sum   int
}

// Simulated data aggregation task
func aggregateData(id int, data []int, aggData *AggregatedData, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	count := 0

	for _, value := range data {
		// Simulate some processing time
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		sum += value
		count++
	}

	// Update the aggregated data structure using a mutex to prevent race conditions
	mutex.Lock()
	aggData.Count += count
	aggData.Sum += sum
	mutex.Unlock()

	fmt.Printf("Task %d: Aggregated count = %d, sum = %d\n", id, count, sum)
}

var mutex sync.Mutex // mutex to protect shared aggregatedData

func main() {
	const numTasks = 5
	const numElementsPerTask = 100

	var wg sync.WaitGroup
	aggregatedData := &AggregatedData{}

	// Simulate different data sets for each task
	dataSets := make([][]int, numTasks)
	for i := 0; i < numTasks; i++ {
		dataSets[i] = generateRandomData(numElementsPerTask)
	}

	// Launch parallel aggregation tasks
	for id, subData := range dataSets {
		wg.Add(1)
		go aggregateData(id, subData, aggregatedData, &wg)
	}

	// Wait for all tasks to complete before accessing the aggregatedData
	wg.Wait()

	// Calculate and print the final aggregated result
	fmt.Println("--- Final Aggregated Result ---")
	fmt.Println("Final Count:", aggregatedData.Count)
	fmt.Println("Final Sum:  ", aggregatedData.Sum)
}

// Function to generate random data for testing
func generateRandomData(count int) []int {
	data := make([]int, count)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	return data
}
