package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func fetchURLFanOut(url string) <-chan string {
	response := make(chan string)

	go func() {
		defer close(response)
		// Simulate HTTP request
		time.Sleep(time.Duration(randomDuration()) * time.Second)
		// Prepare a sample response for the example
		response <- fmt.Sprintf("Fetched %s", url)
	}()

	return response
}

func randomDuration() int {
	return int(rand.Intn(5) + 1) // Random duration between 1 and 5 seconds
}

func main() {
	// Fan-Out: Start goroutines for each URL
	url1 := fetchURLFanOut("url-1")
	url2 := fetchURLFanOut("url-2")
	url3 := fetchURLFanOut("url-3")

	// Fan-In: Collect results
	results := getResponseFanIn(url1, url2, url3)

	for response := range results {
		log.Println(response)
	}

	log.Println("All URLs fetched.")
}

func getResponseFanIn(inputs ...<-chan string) <-chan string {
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

				out <- value
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
