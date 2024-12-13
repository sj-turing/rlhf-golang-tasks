package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulating data fetching from a source
func fetchData(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	// Simulate a random delay for data fetching
	delay := rand.Intn(5)
	time.Sleep(time.Duration(delay) * time.Second)
	result := fmt.Sprintf("Data from source %d", id)
	results <- result // Send the result to the channel
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 5) // Buffered channel to hold results

	// Launch multiple goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)                     // Increment the WaitGroup counter
		go fetchData(i, &wg, results) // Start the goroutine
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(results) // Close the channel once all results are sent
	}()

	// Receive results from the channel
	for result := range results {
		fmt.Println(result) // Process the results
	}

	fmt.Println("All data fetched!")
}
