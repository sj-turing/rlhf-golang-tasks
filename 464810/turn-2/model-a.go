package main

/*

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkForLoop(b *testing.B) {
	var m runtime.MemStats
	largeSlice := make([]int, 1000000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for j := 0; j < len(largeSlice); j++ {
			sum += largeSlice[j]
		}
		// Prevent optimization by using the result
		b.SetBytes(int64(len(largeSlice)))
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("For loop: Alloc = %d, Sys = %d\n", m.Alloc, m.Sys)
}

func BenchmarkRangeLoop(b *testing.B) {
	var m runtime.MemStats
	largeSlice := make([]int, 1000000)
	for i := 0; i < len(largeSlice); i++ {
		largeSlice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for _, value := range largeSlice {
			sum += value
		}
		// Prevent optimization by using the result
		b.SetBytes(int64(len(largeSlice)))
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("Range loop: Alloc = %d, Sys = %d\n", m.Alloc, m.Sys)
}

func main() {
	testing.Main()
}
*/
