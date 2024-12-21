package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmtSprintf(b *testing.B) {
	var num int64 = 1234567890
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", num)
	}
}

func BenchmarkStrconvFormatInt(b *testing.B) {
	var num int64 = 1234567890
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(num, 10)
	}
}

func main() {
	testing.Main()
}
