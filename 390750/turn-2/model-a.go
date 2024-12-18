package main

/*

import (
	"fmt"
	"strings"
	"sync"
)

// Function to convert a word to uppercase.
func toUpperCase(word string) string {
	return strings.ToUpper(word)
}

// Function to reverse a string.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Fan-out function to convert multiple words to uppercase concurrently.
func fanOutToUpper(words []string) <-chan string {
	upperCaseChannel := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(words))

	for _, word := range words {
		go func(w string) {
			defer wg.Done()
			upperCaseChannel <- toUpperCase(w)
		}(word)
	}

	go func() {
		wg.Wait()
		close(upperCaseChannel)
	}()

	return upperCaseChannel
}

// Fan-in function to reverse words collected from multiple goroutines.
func fanInReverse(upperCaseChannel <-chan string) <-chan string {
	reversedChannel := make(chan string)

	go func() {
		for word := range upperCaseChannel {
			reversedChannel <- reverseString(word)
		}
		close(reversedChannel)
	}()

	return reversedChannel
}

func main() {
	words := []string{"hello", "world", "golang", "concurrency"}

	// Step 1: Fan-out to convert words to uppercase.
	upperCaseChannel := fanOutToUpper(words)

	// Step 2: Fan-in to reverse the uppercase words.
	reversedChannel := fanInReverse(upperCaseChannel)

	// Print the reversed words.
	for reversedWord := range reversedChannel {
		fmt.Println(reversedWord)
	}
}
*/
