package main

import (
	"fmt"
	"time"
)

func main() {
	data := []int{1, 2, 3, 4, 5} // You can also use an array with `var data [5]int = {1, 2, 3, 4, 5}`
	for index, value := range data {
		fmt.Printf("Element %d: %d\n", index, value)
	}
	// Output:
	// Element 0: 1
	// Element 1: 2
	// Element 2: 3
	// Element 3: 4
	// Element 4: 5
}
