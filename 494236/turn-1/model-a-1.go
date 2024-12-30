package main

import (
	"fmt"
	"time"
)

func processBatch(batch []int) {
	for _, value := range batch {
		fmt.Println("Processing value:", value)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	data := make([]int, 100)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	batchSize := 10
	for i := 0; i < len(data); i += batchSize {
		batch := data[i : i+batchSize]
		processBatch(batch)
	}
}
