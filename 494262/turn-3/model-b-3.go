package main

import (
	"fmt"
	"sync"
)

func processData(data []int) int {
	total := 0
	for _, value := range data {
		total += value
	}
	return total
}

func main() {
	const numRoutines = 100
	dataSize := 1000000 // Large data size
	var wg sync.WaitGroup

	results := make([]int, numRoutines)
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		data := make([]int, dataSize) // Excessive memory allocation for each loop
		for j := 0; j < dataSize; j++ {
			data[j] = j
		}
		go func(idx int) {

			results[idx] = processData(data)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Results:", results)
}
