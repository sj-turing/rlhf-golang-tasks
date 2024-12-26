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
	dataCh := make(chan []int, numRoutines)
	wg.Add(numRoutines)

	go func() {
		for i := 0; i < numRoutines; i++ {
			data := make([]int, dataSize) // Allocate memory once
			for j := 0; j < dataSize; j++ {
				data[j] = j
			}
			dataCh <- data
		}
		close(dataCh)
	}()

	for i := 0; i < numRoutines; i++ {
		go func(idx int) {
			for data := range dataCh {
				results[idx] = processData(data)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Results:", results)
}
