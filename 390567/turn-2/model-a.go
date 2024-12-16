package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Function to print current memory usage
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

// Worker function simulating heavy workload
func worker(id int, wg *sync.WaitGroup, pool *sync.Pool) {
	defer wg.Done()

	// Acquire object from pool or create a new one if none are available
	buf := pool.Get().([]byte)[:0] // Reuse the slice to reduce allocations
	defer func() {
		pool.Put(buf) // Return the buffer to the pool after use
	}()

	// Simulating work
	for i := 0; i < 100000; i++ {
		buf = append(buf, byte(id)) // Modify the buffer
	}
}

func main() {
	var wg sync.WaitGroup

	// Initialize a sync.Pool to manage reusable memory
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 64) // Pre-allocate a slice with length 0 and capacity of 64
		},
	}

	// Monitor initial memory usage
	fmt.Println("Initial Memory Usage:")
	printMemoryUsage()

	numWorkers := 10
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg, pool)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Print memory usage after Goroutines finish and before GC
	fmt.Println("Memory Usage After Workers Complete:")
	printMemoryUsage()

	// Force garbage collection for analysis
	runtime.GC()
	fmt.Println("Memory Usage After GC:")
	printMemoryUsage()
}
