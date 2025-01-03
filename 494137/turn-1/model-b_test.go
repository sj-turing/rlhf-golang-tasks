package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Struct to test map performance with struct keys
type keyStruct struct {
	A int
	B string
}

func BenchmarkMapStringKey(b *testing.B) {
	m := make(map[string]int)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		m[key] = i
	}
}

func BenchmarkMapIntKey(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkMapStructKey(b *testing.B) {
	m := make(map[keyStruct]int)
	for i := 0; i < b.N; i++ {
		key := keyStruct{A: i, B: fmt.Sprintf("key-%d", i)}
		m[key] = i
	}
}

func main() {
	testing.Benchmark(BenchmarkMapStringKey)
	testing.Benchmark(BenchmarkMapIntKey)
	testing.Benchmark(BenchmarkMapStructKey)
}
