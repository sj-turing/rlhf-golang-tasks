// filename: main.go
package main

import (
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
	m := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		key := generateRandomWord(50)
		m[key] = struct{}{}
	}
}

func BenchmarkMapStructKey(b *testing.B) {
	m := make(map[keyStruct]struct{})
	for i := 0; i < b.N; i++ {
		key := keyStruct{Word: generateRandomWord(50)}
		m[key] = struct{}{}
	}
}

func BenchmarkMapInterfaceKey(b *testing.B) {
	m := make(map[any]struct{})
	for i := 0; i < b.N; i++ {
		key := generateRandomWord(50)
		m[key] = struct{}{}
	}
}
