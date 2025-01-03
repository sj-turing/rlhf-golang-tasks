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
