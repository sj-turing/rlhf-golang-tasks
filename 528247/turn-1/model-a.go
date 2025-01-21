package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// ConnectToDatabase establishes a new database connection
func ConnectToDatabase() (*sql.DB, error) {
	connStr := "user=username dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func fetchURL(client *http.Client, url string, wg *sync.WaitGroup, responses chan<- string) {
	defer wg.Done() // Ensures the WaitGroup counter decrements when the goroutine completes
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to fetch %s: %v", url, err)
		return
	}
	defer resp.Body.Close() // Ensures response body is closed to free network resources

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response from %s: %v", url, err)
		return
	}
	responses <- string(body)
}

func main() {
	// Database connection
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure that the database connection is always closed

	// URLs to fetch concurrently
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
	}

	// HTTP client with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	var wg sync.WaitGroup
	responses := make(chan string, len(urls))

	// Start fetching URLs concurrently
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(client, url, &wg, responses)
	}

	// Closing the responses channel once all goroutines are done
	go func() {
		wg.Wait()
		close(responses)
	}()

	// Insert responses into the database
	for response := range responses {
		if _, err := db.Exec("INSERT INTO responses (body) VALUES ($1)", response); err != nil {
			log.Printf("Failed to insert response: %v", err)
		}
	}

	fmt.Println("Done fetching and storing URL responses.")
}
