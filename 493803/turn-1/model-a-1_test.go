package main

import (
	"testing"
	"time"
)

var globalVar int = 0

func TestGarbageCollectionImpact(t *testing.T) {
	for i := 0; i < 10000; i++ {
		globalVar++
		if i%1000 == 0 {
			// Simulate some work
			time.Sleep(100 * time.Millisecond)
		}
	}
}
