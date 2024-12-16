package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines = 10000
	iterations    = 100000
)

// Function to simulate work in a goroutine
func work(wg *sync.WaitGroup, data *[]int) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		// Simulate work by creating and discarding slices
		_ = make([]int, 1000)
	}
}

func main() {
	var wg sync.WaitGroup
	data := make([]int, 0)

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go work(&wg, &data)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print initial memory stats
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("Initial Memory Stats:")
	printMemStats(&ms)

	// Optimize by reusing slices
	optimizedWork(&wg)

	// Print optimized memory stats
	runtime.ReadMemStats(&ms)
	fmt.Println("\nOptimized Memory Stats:")
	printMemStats(&ms)
}

// Optimized function to simulate work in a goroutine
func optimizedWork(wg *sync.WaitGroup) {
	var slice []int
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		// Reuse the same slice
		slice = slice[:0]
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}
}

// Helper function to print memory stats
func printMemStats(ms *runtime.MemStats) {
	fmt.Printf("Alloc: %d\n", ms.Alloc)
	fmt.Printf("TotalAlloc: %d\n", ms.TotalAlloc)
	fmt.Printf("Sys: %d\n", ms.Sys)
	fmt.Printf("NumGC: %d\n", ms.NumGC)
	fmt.Printf("PauseTotalNs: %d\n", ms.PauseTotalNs)
}
