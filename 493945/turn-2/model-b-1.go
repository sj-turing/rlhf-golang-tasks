package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	n := 100
	start := time.Now()
	result := factorial(n)
	elapsed := time.Since(start)
	fmt.Printf("Factorial of %d is: %d\n", n, result)
	fmt.Printf("Calculation took: %s\n", elapsed)
}
