package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Function to print memory usage
func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, Sys = %v MiB, NumGC = %v\n", bToMb(m.Alloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Worker function with optimized memory allocation
func optimizedWorker(id int, buffer *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	// Using a pre-allocated buffer to avoid repeated allocations
	for i := 0; i < 100000; i++ {
		for j := range *buffer {
			(*buffer)[j] = byte(i)
		}
	}
}

// Worker function with unoptimized memory allocation
func unoptimizedWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulating work and frequent memory allocation
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 64) // Allocate 64 bytes
	}
}

func main() {
	var wg sync.WaitGroup

	// Start monitoring memory usage
	printMemoryUsage()

	// Optimize: Use a pre-allocated buffer for workers
	buffer := make([]byte, 64)

	numWorkers := 10
	wg.Add(numWorkers)

	fmt.Println("Starting optimized workers...")
	for i := 0; i < numWorkers; i++ {
		go optimizedWorker(i, &buffer, &wg)
	}

	// Wait for all optimized Goroutines to finish
	wg.Wait()
	time.Sleep(1 * time.Second) // Wait for GC to settle

	printMemoryUsage()
	fmt.Println("\nOptimized Workers Complete.")

	// Unoptimized: Simulate frequent allocations
	numWorkers = 10
	wg.Add(numWorkers)

	fmt.Println("Starting unoptimized workers...")
	for i := 0; i < numWorkers; i++ {
		go unoptimizedWorker(i, &wg)
	}

	// Wait for all unoptimized Goroutines to finish
	wg.Wait()
	time.Sleep(1 * time.Second) // Wait for GC to settle

	printMemoryUsage()
	fmt.Println("\nUnoptimized Workers Complete.")
}
