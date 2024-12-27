package main

import (
	"runtime"
	"testing"
	"time"
)

var globalVar int = 0

func TestGCStats(t *testing.T) {
	var gcStats runtime.MemStats
	// Before starting work
	runtime.ReadMemStats(&gcStats)
	t.Logf("GC Stats: %+v\n", gcStats.TotalAlloc)
	for i := 0; i < 10000; i++ {
		globalVar++
		if i%1000 == 0 {
			// Simulate some work
			time.Sleep(100 * time.Millisecond)
		}
	}
	// After work
	runtime.ReadMemStats(&gcStats)
	t.Logf("GC Stats: %+v\n", gcStats.TotalAlloc)
}
