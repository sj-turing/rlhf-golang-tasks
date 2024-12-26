package main

import (
	"fmt"
	"sync"
)

func processData(data []int, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	total := 0
	for _, value := range data {
		total += value
	}
	result <- total
}

func main() {
	const numRoutines = 4
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Example data slice
	result := make(chan int, numRoutines)
	var wg sync.WaitGroup

	// Divide the data into chunks for each routine
	chunkSize := len(data) / numRoutines
	if len(data)%numRoutines != 0 {
		chunkSize++
	}

	// Start worker routines
	for i := 0; i < numRoutines; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, len(data))
		wg.Add(1)
		go processData(data[start:end], &wg, result)
	}

	// Wait for all routines to finish
	go func() {
		wg.Wait()
		close(result)
	}()

	// Aggregate results from all routines
	totalSum := 0
	for sum := range result {
		totalSum += sum
	}

	fmt.Println("Total sum:", totalSum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
