package main

import (
	"fmt"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	sumOfLengths := 0
	for _, name := range names {
		sumOfLengths += len(name)
	}
	fmt.Println("Sum of lengths:", sumOfLengths) // Output: Sum of lengths: 15
}
