package main

import (
	"sync"
	"testing"
)

func TestLocalVarConcurrency(t *testing.T) {
	const numGoroutines = 100
	const numIncrements = 10000

	var wg = &sync.WaitGroup{}
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			localVar := 0
			for i := 0; i < numIncrements; i++ {
				localVar++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	// Assertions to verify the correctness of localVar values
}
