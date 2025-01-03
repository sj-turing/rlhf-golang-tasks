// filename: main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"rand"
	"strings"
	"testing"
)

// Generate random words for testing
func generateRandomWords(n int) []string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	words := make([]string, n)
	for i := range words {
		length := rand.Intn(20) + 1
		word := make([]rune, length)
		for j := range word {
			word[j] = letters[rand.Intn(len(letters))]
		}
		words[i] = string(word)
	}
	return words
}

// Function to count occurrences of words using a map with string keys
func countWordsWithStringKeys(words []string) map[string]int {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

// Function to count occurrences of words using a map with struct keys
func countWordsWithStructKeys(words []string) map[keyStruct]int {
	m := make(map[keyStruct]int)
	for _, word := range words {
		m[keyStruct{Word: word}]++
	}
	return m
}

// Function to count occurrences of words using a map with slice keys
func countWordsWithSliceKeys(words []string) map[[]rune]int {
	m := make(map[[]rune]int)
	for _, word := range words {
		runeWord := []rune(word)
		m[runeWord]++
	}
	return m
}

type keyStruct struct {
	Word string
}

func BenchmarkCountWordsWithStringKeys(b *testing.B) {
	words := generateRandomWords(b.N)
	for i := 0; i < b.N; i++ {
		countWordsWithStringKeys(words)
	}
}

func BenchmarkCountWordsWithStructKeys(b *testing.B) {
	words := generateRandomWords(b.N)
	for i := 0; i < b.N; i++ {
		countWordsWithStructKeys(words)
	}
}

func BenchmarkCountWordsWithSliceKeys(b *testing.B) {
	words := generateRandomWords(b.N)
	for i := 0; i < b.N; i++ {
		countWordsWithSliceKeys(words)
	}
}

func main() {
	// Example usage
	file, err := os.Open("text.txt") // Change "text.txt" to your file path
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Print the top 10 most frequent words
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
