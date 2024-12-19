package main

import "fmt"

func main() {
	largeArray := make([]int, 1000000)
	for i := 0; i < len(largeArray); i++ {
		largeArray[i] = i
	}
	// Traditional for loop
	sum1 := 0
	for i := 0; i < len(largeArray); i++ {
		sum1 += largeArray[i]
	}
	fmt.Println("Sum using traditional for loop:", sum1)
	// Range loop
	sum2 := 0
	for _, value := range largeArray {
		sum2 += value
	}
	fmt.Println("Sum using range loop:", sum2)
}
