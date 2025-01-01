package main

import (
	"fmt"
	"sync"
	"testing"
)

const (
	numGoroutines = 1000
	numOps        = 100000
)

func BenchmarkStandardMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myMap := make(map[int]int)
		var wg sync.WaitGroup

		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func() {
				for j := 0; j < numOps; j++ {
					myMap[j] = j
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myMap := new(sync.Map)
		var wg sync.WaitGroup

		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func() {
				for j := 0; j < numOps; j++ {
					myMap.Store(j, j)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
