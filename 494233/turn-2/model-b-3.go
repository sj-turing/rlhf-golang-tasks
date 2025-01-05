package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	transposed := make([][]int, 3)
	for i := range matrix {
		for j := range matrix[i] {
			transposed[j] = append(transposed[j], matrix[i][j])
		}
	}
	fmt.Println("Transposed matrix:")
	for _, row := range transposed {
		fmt.Println(row)
	}
	// Output:
	// Transposed matrix:
	// [1 4 7]
	// [2 5 8]
	// [3 6 9]
}
