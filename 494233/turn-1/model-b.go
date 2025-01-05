package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a large slice of random integers
	slice := make([]int, 1000000)
	for i := range slice {
		slice[i] = rand.Intn(1000000)
	}

	// Original implementation with range loop
	start := time.Now()
	sum := 0
	for _, num := range slice {
		sum += num
	}
	fmt.Println("Original implementation:", time.Since(start))

	// Optimized implementation using preallocation and indexing
	start = time.Now()
	sum = 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	fmt.Println("Optimized implementation:", time.Since(start))
}
