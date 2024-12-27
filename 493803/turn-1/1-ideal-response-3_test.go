package main

import (
	"sync"
	"testing"
)

var globalVar int = 0

func incrementGlobalVar(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		globalVar++
	}
}

func TestGarbageCollectionImpactConcurrency(t *testing.T) {
	const numGoroutines = 100
	const numIncrements = 10000

	var wg = &sync.WaitGroup{}

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go incrementGlobalVar(numIncrements, wg)
	}

	wg.Wait()

	expected := numGoroutines * numIncrements
	if globalVar != expected {
		t.Errorf("Expected %d, got %d", expected, globalVar)
	}
}
