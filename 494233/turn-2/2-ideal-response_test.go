package main

import (
	"bytes"
	"testing"
)

func BenchmarkIterateArray(b *testing.B) {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	for i := 0; i < b.N; i++ {
		iterateArray(buf)
	}
}

func BenchmarkIterateSlice(b *testing.B) {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	for i := 0; i < b.N; i++ {
		iterateSlice(buf)
	}
}
