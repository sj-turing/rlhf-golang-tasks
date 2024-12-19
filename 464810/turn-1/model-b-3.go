package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	largeArray := make([]int, 1000000)
	for i := 0; i < len(largeArray); i++ {
		largeArray[i] = i
	}
	// Traditional for loop
	sum1 := 0
	for i := 0; i < len(largeArray); i++ {
		sum1 += largeArray[i]
	}
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage after traditional for loop:", m.Alloc)
	// Bounded range loop
	sum2 := 0
	for i := 0; i < len(largeArray); i += 10 {
		for j := 0; j < 10; j++ {
			if i+j < len(largeArray) {
				sum2 += largeArray[i+j]
			}
		}
	}
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage after bounded range loop:", m.Alloc)
}
