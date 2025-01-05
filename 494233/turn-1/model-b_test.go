package main

import (
	"testing"
)

func BenchmarkRangeLoop(b *testing.B) {
	// Generate a large slice of random integers
	slice := make([]int, 1000000)
	for i := range slice {
		slice[i] = rand.Intn(1000000)
	}

	// Original implementation with range loop
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, num := range slice {
			sum += num
		}
	}
}

func BenchmarkIndexing(b *testing.B) {
	// Generate a large slice of random integers
	slice := make([]int, 1000000)
	for i := range slice {
		slice[i] = rand.Intn(1000000)
	}

	// Optimized implementation using preallocation and indexing
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < len(slice); i++ {
			sum += slice[i]
		}
	}
}
