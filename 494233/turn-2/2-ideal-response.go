package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func iterateArray(output io.Writer) {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	now := time.Now()
	for _, value := range arr {
		fmt.Fprintln(output, value)
	}

	fmt.Fprintf(output, "Time taken by range loop for iterating array: %v\n", time.Since(now))
}

func iterateSlice(output io.Writer) {
	slice := make([]int, 10)

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}

	now := time.Now()
	for _, value := range slice {
		fmt.Fprintln(output, value)
	}
	fmt.Fprintf(output, "Time taken by range loop for iterating slice: %v\n", time.Since(now))
}

func main() {
	iterateArray(os.Stdout)
	iterateSlice(os.Stdout)
}
