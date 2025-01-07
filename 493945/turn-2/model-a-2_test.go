package main

import (
	"testing"
)

// IterativeFactorial calculates the factorial of a number using iteration
func IterativeFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// BenchmarkIterativeFactorial benchmarks the iterative factorial function
func BenchmarkIterativeFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativeFactorial(10) // Use a large number for significant benchmarking
	}
}

func main() {
	testing.Main()
}
