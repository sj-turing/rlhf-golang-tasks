package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Function to monitor memory usage
func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Worker function simulating work
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulating work and memory allocation
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 64) // Allocate 64 bytes
	}
}

func main() {
	var wg sync.WaitGroup

	// Start monitoring memory usage
	printMemoryUsage()

	numWorkers := 10
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Print memory usage after Goroutines finish
	printMemoryUsage()

	// Force garbage collection for analysis
	runtime.GC()
	printMemoryUsage()
}
