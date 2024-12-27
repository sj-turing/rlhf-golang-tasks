package main

import (
	"runtime"
	"sync"
	"testing"
)

var globalCounter int
var globalMap = make(map[int]int)

// TestGlobalCounterWithGC verifies that globalCounter is not affected by GC.
func TestGlobalCounterWithGC(t *testing.T) {
	for i := 0; i < 10000; i++ {
		globalCounter++
		if i%1000 == 0 {
			runtime.GC()
			if globalCounter != i+1 {
				t.Errorf("Expected %d, got %d after GC at iteration %d", i+1, globalCounter, i)
			}
		}
	}
}

// TestGlobalMapWithGC verifies that globalMap is not affected by GC.
func TestGlobalMapWithGC(t *testing.T) {
	for i := 0; i < 10000; i++ {
		globalMap[i] = i
		if i%1000 == 0 {
			runtime.GC()
			if len(globalMap) != i+1 {
				t.Errorf("Expected %d keys, got %d after GC at iteration %d", i+1, len(globalMap), i)
			}
		}
	}
}

// TestGlobalCounterWithConcurrency verifies globalCounter is safe for concurrent use.
func TestGlobalCounterWithConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	const numGoroutines = 100
	const numIncrements = 10000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numIncrements; j++ {
				globalCounter++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	expected := numGoroutines * numIncrements
	if globalCounter != expected {
		t.Errorf("Expected %d, got %d", expected, globalCounter)
	}
}

// TestGlobalMapWithConcurrency verifies globalMap is safe for concurrent use.
func TestGlobalMapWithConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	const numGoroutines = 100
	const numIncrements = 10000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numIncrements; j++ {
				globalMap[j] = j
			}
			wg.Done()
		}()
	}

	wg.Wait()
	expected := numGoroutines * numIncrements
	if len(globalMap) != expected {
		t.Errorf("Expected %d keys, got %d", expected, len(globalMap))
	}
}

// TestGCStatistics verifies that GC is triggered and performs as expected.
func TestGCStatistics(t *testing.T) {
	var gcStats runtime.MemStats
	runtime.ReadMemStats(&gcStats)
	initialAlloc := gcStats.TotalAlloc

	// Simulate heavy memory usage
	for i := 0; i < 1000000; i++ {
		globalMap[i] = i
	}

	runtime.GC()

	runtime.ReadMemStats(&gcStats)
	finalAlloc := gcStats.TotalAlloc

	if finalAlloc >= initialAlloc {
		t.Errorf("Total memory allocation should decrease after GC, got %d before and %d after", initialAlloc, finalAlloc)
	}
}

func BenchmarkGlobalCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalCounter++
	}
}

func BenchmarkGlobalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalMap[i] = i
	}
}
