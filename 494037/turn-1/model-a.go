package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated data aggregation task
func aggregateData(id int, data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, value := range data {
		sum += value
	}
	fmt.Printf("Task %d: Aggregated sum = %d\n", id, sum)
}

// Simulate event handling
func handleEvent(data [][]int) {
	var wg sync.WaitGroup

	// Launch parallel aggregation tasks
	for id, subData := range data {
		wg.Add(1)
		go aggregateData(id, subData, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()
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
	handleEvent(data)

	// Simulate another event after a delay
	time.Sleep(2 * time.Second)
	handleEvent(data)
}
