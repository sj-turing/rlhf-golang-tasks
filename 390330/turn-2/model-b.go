package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func fetch(ctx context.Context, url string) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Return the error if the context is done
	default:
	}

	// Simulate a slow HTTP request
	time.Sleep(2 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Fetched %s\n", url)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	urls := []string{
		"https://example.com",
		"https://example.com",
		"https://example.com",
	}

	// Launch goroutines to fetch URLs
	for _, url := range urls {
		go func(u string) {
			if err := fetch(ctx, u); err != nil {
				log.Printf("Fetching %s failed: %v\n", u, err)
			}
		}(url)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Canceling requests...")
	cancel() // Cancel the context, which will cancel all ongoing fetches

	time.Sleep(1 * time.Second) // Wait for goroutines to check the context
}
