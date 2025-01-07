package main

import (
	"testing"
)

// BenchmarkTestFunction benchmarks the performance of testFunction
func BenchmarkTestFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFunction()
	}
}

func testFunction() {
	// Your function logic here
}

func main() {
	testing.Main()
}
