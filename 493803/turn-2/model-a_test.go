package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Example of a global variable
var globalCounter int

func incrementCounter() {
	globalCounter++
}

func TestGCImpactOnGlobalVariables(t *testing.T) {
	// Initialize GC stats
	var gcStats runtime.MemStats
	runtime.ReadMemStats(&gcStats)
	t.Logf("Initial GC Stats: %+v\n", gcStats)

	// Simulate long-running work
	for i := 0; i < 100000; i++ {
		incrementCounter()
		if i%10000 == 0 {
			// Trigger a GC manually
			runtime.GC()
			runtime.ReadMemStats(&gcStats)
			t.Logf("GC Stats after %d iterations: %+v\n", i, gcStats)
		}
	}

	// Verify final counter value
	expected := 100000
	if globalCounter != expected {
		t.Errorf("Expected %d, got %d", expected, globalCounter)
	}
}
