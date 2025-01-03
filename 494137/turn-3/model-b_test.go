package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func countWords(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		wordCount[word]++
	}

	return wordCount, scanner.Err()
}

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = countWords("testdata/words.txt")
	}
}

// Test using map[string]int
func BenchmarkCountWordsStringInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = countWordsStringInt("testdata/words.txt")
	}
}

// Test using map[string]uint64
func BenchmarkCountWordsStringUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = countWordsStringUint64("testdata/words.txt")
	}
}

// Test using map[interface{}]int
func BenchmarkCountWordsInterfaceInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = countWordsInterfaceInt("testdata/words.txt")
	}
}

// Implementation of countWords using map[string]int
func countWordsStringInt(fileName string) (map[string]int, error) {
	// ... (Same code as before)
}

// Implementation of countWords using map[string]uint64
func countWordsStringUint64(fileName string) (map[string]uint64, error) {
	wordCount := make(map[string]uint64)
	// ... (Same code as before, but increment count using uint64)
}

// Implementation of countWords using map[interface{}]int
func countWordsInterfaceInt(fileName string) (map[interface{}]int, error) {
	wordCount := make(map[interface{}]int)
	// ... (Same code as before, but use interface{} as key)
}

func main() {
	wordCount, err := countWords("testdata/words.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Word Count:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
