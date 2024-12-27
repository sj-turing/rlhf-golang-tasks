package main

import (
	"runtime"
	"testing"
	"time"
)

var globalCounter int

func TestGarbageCollectionImpact(t *testing.T) {
	t.Run("IncreaseGlobalCounter", func(t *testing.T) {
		increaseGlobalCounter()
		// Verify that the global counter is increased
		if globalCounter != 1 {
			t.Errorf("Global counter should be 1 after increase, got: %d", globalCounter)
		}
	})

	t.Run("ForceGarbageCollectionAndCheckCounter", func(t *testing.T) {
		// Force garbage collection to occur
		runtime.GC()
		// Sleep for a little while to give GC time to run
		time.Sleep(time.Millisecond * 100)
		// Verify that the global counter is still 1 after GC
		if globalCounter != 1 {
			t.Errorf("Global counter should be 1 after GC, got: %d", globalCounter)
		}
	})
}

func increaseGlobalCounter() {
	globalCounter++
}
