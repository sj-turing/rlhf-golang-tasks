package main

import (
	"fmt"
	"sync"
	"time"
)

// FetchOptions defines the options for the fetch function
type FetchOptions struct {
	// SleepDuration defines how long the goroutine should sleep before executing
	SleepDuration time.Duration

	// Logger defines a function to log messages
	Logger func(string)
}

// Fetch performs a simulated fetch operation with optional configuration
func Fetch(url string, options *FetchOptions) {
	// Default options
	if options == nil {
		options = &FetchOptions{}
	}

	// Log the start of the fetch
	if options.Logger != nil {
		options.Logger("Fetching URL: " + url)
	}

	// Simulate fetch operation with a sleep
	time.Sleep(options.SleepDuration)

	// Simulate a successful fetch
	result := "Successfully fetched " + url

	// Log the result
	if options.Logger != nil {
		options.Logger(result)
	}
}

// Main demonstrates concurrent use of the Fetch function with options
func main() {
	var wg sync.WaitGroup
	var urls = []string{"https://example.com", "https://google.com", "https://github.com"}

	// Create a logger function for demonstration
	logger := func(message string) {
		fmt.Println("Log:", message)
	}

	for _, url := range urls {
		// Increment the wait group counter
		wg.Add(1)

		// Create options with custom sleep duration and logger
		options := &FetchOptions{
			SleepDuration: time.Second,
			Logger:        logger,
		}

		// Launch a goroutine to fetch the URL
		go func(url string) {
			defer wg.Done()
			Fetch(url, options)
		}(url)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All fetch operations completed.")
}
