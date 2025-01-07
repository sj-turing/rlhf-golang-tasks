package main

import (
	"testing"
)

// Factorial calculates the factorial of a number using recursion
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// BenchmarkFactorial benchmarks the recursive factorial function
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10) // Use a large number for significant benchmarking
	}
}

func main() {
	testing.Main()
}
