package main

import (
	"fmt"
	"net/http"
	"time"
)

// maxRetries specifies the maximum number of retries.
const maxRetries = 5

// backoffFactor is the base value for the exponential backoff.
const backoffFactor = 200 * time.Millisecond // 200ms

// get mit api mit retry logic und exponential backoff
func getWithRetry(url string) (resp *http.Response, err error) {
	for attempts := 0; attempts < maxRetries; attempts++ {
		resp, err = http.Get(url)

		// Check if the response is successful or if we should retry
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		// If there's an error or a failed status code, determine the wait time.
		backoffTime := time.Duration(float64(backoffFactor) * (1 << uint(attempts)))
		fmt.Printf("Attempt %d failed; retrying in %v...\n", attempts+1, backoffTime)

		// Wait before retrying
		time.Sleep(backoffTime)
	}

	// Return the last error encountered
	return nil, fmt.Errorf("failed to get the response after %d attempts: %w", maxRetries, err)
}

func main() {
	url := "https://api.example.com/data"

	resp, err := getWithRetry(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	// Process the response body as needed ...
}
