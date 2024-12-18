package main

/*

import (
	"fmt"
	"sync"
	"strings"
)

// ProcessWord converts a word to uppercase and returns it
func processWordFanOut(word string) <-chan string {
	response := make(chan string)

	go func() {
		defer close(response)
		// Convert to uppercase
		upperWord := strings.ToUpper(word)
		// Simulate some work
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		response <- upperWord
	}()

	return response
}

// ReverseAndCollect collects results from all goroutines, reverses each word, and returns them
func reverseAndCollectFanIn(inputs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan string) {
			for {
				value, ok := <-ch

				if !ok {
					wg.Done()
					break
				}

				// Reverse the uppercase word
				reversed := string([]byte(value)[::-1])
				out <- reversed
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	words := []string{"hello", "world", "golang", "concurrency"}

	// Fan-Out: Convert each word to uppercase
	upperChannels := make([]<-chan string, len(words))
	for i, word := range words {
		upperChannels[i] = processWordFanOut(word)
	}

	// Fan-In: Collect, reverse, and output results
	results := reverseAndCollectFanIn(upperChannels...)

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All words processed.")
}
*/
