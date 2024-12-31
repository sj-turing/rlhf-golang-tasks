package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const (
	numWords   = 10000
	wordLength = 10
)

func generateWords() []string {
	words := make([]string, numWords)
	for i := range words {
		b := make([]byte, wordLength)
		for j := 0; j < wordLength; j++ {
			b[j] = byte('a' + rand.Intn(26))
		}
		words[i] = string(b)
	}
	return words
}

// Benchmark Join
func BenchmarkStringsJoin(b *testing.B) {
	words := generateWords()
	delimiter := " "
	for n := 0; n < b.N; n++ {
		strings.Join(words, delimiter)
	}
}

// Benchmark fmt.Sprintf
func BenchmarkSprintf(b *testing.B) {
	words := generateWords()
	delimiter := " "
	for n := 0; n < b.N; n++ {
		fmt.Sprintf(strings.Repeat("%s%s", len(words)-1), words...) + delimiter
	}
}

// Benchmark bytes.Buffer
func BenchmarkBytesBuffer(b *testing.B) {
	words := generateWords()
	delimiter := " "
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		for i, word := range words {
			buf.WriteString(word)
			if i < len(words)-1 {
				buf.WriteString(delimiter)
			}
		}
	}
}

// Benchmark String Concatenation
func BenchmarkStringConcatenation(b *testing.B) {
	words := generateWords()
	delimiter := " "
	for n := 0; n < b.N; n++ {
		result := ""
		for i, word := range words {
			result += word
			if i < len(words)-1 {
				result += delimiter
			}
		}
	}
}

func main() {
	testing.Main()
}
