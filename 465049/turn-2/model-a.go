package main

/*

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Define a shared resource
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}
	processedItems := make([]string, 0)

	// Start worker Goroutines
	for _, item := range items {
		wg.Add(1) // Increment the WaitGroup counter for each Goroutine
		go processItem(item, &processedItems, &wg)
	}

	// Wait for all Goroutines to finish processing
	wg.Wait()

	// Print the processed items
	fmt.Println("Processed items:", processedItems)
}

func processItem(item string, processedItems *[]string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this Goroutine is done
	fmt.Printf("Processing %s...\n", item)

	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	// Add the processed item to the shared slice (ensuring thread safety)
	processedItemsMutex.Lock()
	defer processedItemsMutex.Unlock()

	*processedItems = append(*processedItems, item)
	fmt.Printf("%s processed.\n", item)
}
*/
