package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// Generate random strings of length 1000 characters
func generateRandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Define the map types we want to test
type mapString string
type mapInt int
type mapByte []byte

func BenchmarkMapStringKey(b *testing.B) {
	m := make(map[mapString]int)
	for i := 0; i < b.N; i++ {
		key := mapString(generateRandomString(1000))
		m[key] = i
	}
}

func BenchmarkMapIntKey(b *testing.B) {
	m := make(map[mapInt]int)
	for i := 0; i < b.N; i++ {
		r := rand.Intn(1000000)
		key := mapInt(r)
		m[key] = i
	}
}

func BenchmarkMapByteKey(b *testing.B) {
	m := make(map[mapByte]int)
	for i := 0; i < b.N; i++ {
		key := mapByte(generateRandomString(1000))
		m[key] = i
	}
}

func main() {
	// Run the benchmark tests
	testing.Benchmark(BenchmarkMapStringKey)
	testing.Benchmark(BenchmarkMapIntKey)
	testing.Benchmark(BenchmarkMapByteKey)
}
