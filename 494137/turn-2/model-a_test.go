package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Struct to test map performance with struct keys
type keyStruct struct {
	Word string
}

// Generate a random long word
func generateRandomWord(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func BenchmarkMapStringKey(b *testing.B) {
	m := make(map[string]int)
	for i := 0; i < b.N; i++ {
		key := generateRandomWord(50)
		m[key] = i
	}
}

func BenchmarkMapByteSliceKey(b *testing.B) {
	m := make(map[[]byte]int)
	for i := 0; i < b.N; i++ {
		key := []byte(generateRandomWord(50))
		m[key] = i
	}
}

func BenchmarkMapStructKey(b *testing.B) {
	m := make(map[keyStruct]int)
	for i := 0; i < b.N; i++ {
		key := keyStruct{Word: generateRandomWord(50)}
		m[key] = i
	}
}

func main() {
	rand.Seed(42)
	fmt.Println("Running benchmarks...")
	testing.Benchmark(BenchmarkMapStringKey)
	testing.Benchmark(BenchmarkMapByteSliceKey)
	testing.Benchmark(BenchmarkMapStructKey)
}
