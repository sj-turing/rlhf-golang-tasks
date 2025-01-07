package main

import (
	"fmt"
	"testing"
)

// Factorial calculates the factorial of a number.
// It's a simple function for demonstration purposes.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// BenchmarkFactorial benchmarks the Factorial function.
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(12)
	}
}

// ExampleFactorial shows an example usage of the Factorial function.
func ExampleFactorial() {
	result := Factorial(5)
	fmt.Println("Factorial of 5 is:", result)
	// Output: Factorial of 5 is: 120
}

func main() {
	fmt.Println("Hello, World!")
}
