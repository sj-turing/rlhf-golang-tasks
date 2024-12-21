package main

import (
	"fmt"
	"math"
)

func nextPowerOfTwo(n uint) uint {
	if n == 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n + 1
}

func resizeSlice(s []int, newCap uint) []int {
	if newCap > uint(cap(s)) {
		newCap = nextPowerOfTwo(newCap)
		newS := make([]int, newCap)
		copy(newS, s)
		return newS
	}
	return s
}

func main() {
	s := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		s = resizeSlice(s, uint(len(s))+1)
		s = append(s, i)
	}
	fmt.Println(s)
}
