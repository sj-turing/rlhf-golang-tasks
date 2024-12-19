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
	// Range loop
	sum2 := 0
	for _, value := range largeArray {
		sum2 += value
	}
	runtime.ReadMemStats(&m)
	fmt.Println("Memory usage after range loop:", m.Alloc)
}
