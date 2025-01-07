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
		result = schoolbookMultiply(result, i)
	}
	return result
}
func schoolbookMultiply(a, b int) int {
	c := 0
	result := 0
	for ; b != 0; b-- {
		c += a
		a += a
		result += c
	}
	return result
}
func main() {
	n := 1000
	start := time.Now()
	result := factorial(n)
	elapsed := time.Since(start)
	fmt.Printf("Factorial of %d is: %d\n", n, result)
	fmt.Printf("Calculation took: %s\n", elapsed)
}
