package main

import (
	"fmt"
	"testing"
)

func appendSliceWithoutResize(b *testing.B) {
	s := make([]int, 0)
	for n := 0; n < b.N; n++ {
		s = append(s, n)
	}
}

func appendSliceWithResize(b *testing.B) {
	s := make([]int, 0, b.N)
	for n := 0; n < b.N; n++ {
		s = append(s, n)
	}
}

func main() {
	b.Run("AppendWithoutResize", appendSliceWithoutResize)
	b.Run("AppendWithResize", appendSliceWithResize)
}
