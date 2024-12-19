package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
)

func sliceSumTraditionalFor(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}

func sliceSumRange(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func BenchmarkSliceSumTraditionalFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100000)
		for j := 0; j < len(slice); j++ {
			slice[j] = rand.Intn(1000)
		}
		sliceSumTraditionalFor(slice)
	}
}

func BenchmarkSliceSumRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100000)
		for j := 0; j < len(slice); j++ {
			slice[j] = rand.Intn(1000)
		}
		sliceSumRange(slice)
	}
}

func main() {
	var m runtime.MemStats

	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(1000)
	}

	// Compare memory usage before and after calculating the sum using the traditional for loop
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage before traditional for loop:", m.Alloc)
	sliceSumTraditionalFor(slice)
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage after traditional for loop: ", m.Alloc)

	// Compare memory usage before and after calculating the sum using the range loop
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage before range loop:", m.Alloc)
	sliceSumRange(slice)
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage after range loop: ", m.Alloc)

	// Run the benchmarks
	testing.Benchmark(BenchmarkSliceSumTraditionalFor)
	testing.Benchmark(BenchmarkSliceSumRange)
}
