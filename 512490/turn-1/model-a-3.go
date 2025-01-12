package main

import (
	"fmt"
	"strings"
)

func processString(input string, result chan string) {
	// Simulate some string processing
	var sb strings.Builder
	for _, char := range input {
		sb.WriteString(string(char) + "*")
	}
	result <- sb.String()
}

func main() {
	strings := []string{"Hello, World!", "Go Programming", "Concurrency", "Goroutines"}
	results := make(chan string, len(strings))

	for _, input := range strings {
		go processString(input, results)
	}

	for _, result := range strings {
		fmt.Println(<-results)
	}
}
