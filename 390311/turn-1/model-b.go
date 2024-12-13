package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// FetchData fetches data from a given URL and returns the content
func FetchData(url string, wg *sync.WaitGroup) (string, error) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Simulate network delay

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// ProcessData processes the fetched data
func ProcessData(data string) string {
	// Simulate processing
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return fmt.Sprintf("Processed: %s", data)
}

func main() {
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
		"https://api.example.com/data4",
	}

	var wg sync.WaitGroup
	results := make([]string, len(urls))

	// Fan-Out: Start Go routines to fetch data
	for i, url := range urls {
		wg.Add(1)
		go func(index int, u string) {
			defer wg.Done()
			data, err := FetchData(u, &wg)
			if err != nil {
				fmt.Printf("Error fetching data from %s: %v\n", u, err)
				results[index] = ""
				return
			}
			results[index] = data
		}(i, url)
	}

	// Fan-In: Wait for all fetches to complete and process the results
	wg.Wait()
	fmt.Println("Fetching completed. Processing data...")

	for i, data := range results {
		if data != "" {
			processedData := ProcessData(data)
			fmt.Printf("Result %d: %s\n", i+1, processedData)
		}
	}
}
