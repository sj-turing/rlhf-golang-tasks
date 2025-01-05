// Example: Sum of squares using range loop
package main

import (
	"fmt"
	"time"
)

func sumSquaresRangeLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func main() {
	numbers := make([]int, 1_000_000)
	for i := range numbers {
		numbers[i] = i + 1
	}

	start := time.Now()
	sum := sumSquaresRangeLoop(numbers)
	duration := time.Since(start)

	fmt.Println("Sum of squares using range loop:", sum)
	fmt.Println("Execution time:", duration)
}
