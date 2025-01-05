package main

import (
	"fmt"
)

func main() {
	keys := []string{"key1", "key2", "key3"}
	values := []int{1, 2, 3} // Oops! Mismatched length
	for i, key := range keys {
		fmt.Printf("%s: %d\n", key, values[i]) // Panic: index out of range
	}
}
