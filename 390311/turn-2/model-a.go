package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulating a data fetching function
func fetchData(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // Mark this goroutine as done

	// Simulate a random delay for fetching data
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	result := fmt.Sprintf("Result from source %d", id)
	results <- result // Send the result back to the results channel
}

func main() {
	const numFetchers = 5

	var wg sync.WaitGroup
	results := make(chan string, numFetchers) // Buffered channel to hold results

	// Start multiple goroutines to fetch data
	for i := 1; i <= numFetchers; i++ {
		wg.Add(1)                     // Increment the WaitGroup counter
		go fetchData(i, &wg, results) // Launch the fetchData goroutine
	}

	// Close the results channel once all fetchers are done
	go func() {
		wg.Wait()      // Wait for all fetchers to complete
		close(results) // Close the results channel
	}()

	// Process the results as they come in
	for result := range results {
		fmt.Println(result) // Process each result (in this case, just print it)
	}

	fmt.Println("All data fetched and processed!")
}
