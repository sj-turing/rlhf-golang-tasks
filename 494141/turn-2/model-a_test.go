package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkStandardMapConcurrency(b *testing.B) {
	const numGoroutines = 10
	const numIterations = 1000000

	var myMap map[int]int = make(map[int]int)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < numIterations; i++ {
				myMap[i] = i
			}
		}
	})
}

func BenchmarkSyncMapConcurrency(b *testing.B) {
	const numGoroutines = 10
	const numIterations = 1000000

	var myMap sync.Map

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < numIterations; i++ {
				myMap.Store(i, i)
			}
		}
	})
}

func main() {
	testing.Init()
	testing.Run()
}
