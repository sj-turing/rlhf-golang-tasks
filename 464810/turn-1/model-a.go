package main

import "fmt"

func main() {
	// Slice example
	slice := make([]int, 1000000)
	fmt.Println("Slice Allocated:", cap(slice)) // Captured for reference

	// Traditional for loop
	for i := range slice {
	}

	// Range loop
	for _, _ = range slice {
	}

	// Map example
	m := make(map[int]int, 1000000)
	for i := 0; i < 1000000; i++ {
		m[i] = i
	}
	fmt.Println("Map Allocated:", len(m)) // Captured for reference

	// Traditional for loop
	for i := range m {
	}

	// Range loop
	for _, _ = range m {
	}
}
