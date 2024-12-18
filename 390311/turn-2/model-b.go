package main

import (
	"fmt"
	"sync"
	"time"
)

// DataFetcher represents the fetched data
type DataFetcher interface {
	Fetch() (string, error) // Fetch data and return it along with an error
}

// SimpleFetcher implements DataFetcher for demonstration
type SimpleFetcher struct {
	id int
}

func (f *SimpleFetcher) Fetch() (string, error) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Simulate fetch delay
	return fmt.Sprintf("Data from source %d", f.id), nil
}

// DataProcessor processes the fetched data
type DataProcessor interface {
	Process(data string)
}

// SimpleProcessor implements DataProcessor for demonstration
type SimpleProcessor struct {
	mu sync.Mutex // Lock for safe concurrent printing
}

func (p *SimpleProcessor) Process(data string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Println("Processing:", data)
}

// FetchData fetches data asynchronously
func FetchData(fetcher DataFetcher, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := fetcher.Fetch()
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	result <- data // Send the fetched data to the result channel
}

// ProcessData processes data concurrently
func ProcessData(processor DataProcessor, result chan string) {
	for data := range result {
		processor.Process(data) // Process the data
	}
}

func main() {
	var wg sync.WaitGroup
	result := make(chan string) // Unbuffered channel for data processing

	// Number of fetchers and processors
	numFetchers := 10
	numProcessors := 3

	// Create fetchers
	fetchers := make([]DataFetcher, numFetchers)
	for i := 0; i < numFetchers; i++ {
		fetchers[i] = &SimpleFetcher{id: i + 1}
	}

	// Create processors
	processors := make([]*SimpleProcessor, numProcessors)
	for i := 0; i < numProcessors; i++ {
		processors[i] = &SimpleProcessor{}
		go ProcessData(processors[i], result) // Start processor goroutines
	}

	// Start fetching data
	for _, fetcher := range fetchers {
		wg.Add(1)
		go FetchData(fetcher, result, &wg) // Start fetcher goroutines
	}

	// Wait for all fetch operations to complete
	wg.Wait()
	close(result) // Close the result channel to signal processors to stop
}
