package main

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	bigNumber = 1_000_000_000
	repeat    = 10_000_000
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", bigNumber)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(bigNumber, 10)
	}
}

func main() {
}
