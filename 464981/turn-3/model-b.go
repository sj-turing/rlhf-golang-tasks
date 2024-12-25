package main

import (
	"log"
	"sync"
	"time"

	"github.com/rs/xid"
)

// Unique callback identifier
type CallbackID string

// Define a type for the callback function
type Callback func(CallbackID, int)

// Simulates an asynchronous operation with a callback
func processData(data int, callback Callback) {
	// Generate a new callback ID
	cbID := CallbackID(xid.New().String())
	log.Printf("[%s] Started processing data: %d", cbID, data)

	// Simulate work with a delay
	time.Sleep(2 * time.Second)

	result := data * data
	callback(cbID, result)
	log.Printf("[%s] Finished processing data, result: %d", cbID, result)
}

// Another callback function
func handleProcessingResult(cbID CallbackID, result int) {
	log.Printf("[%s] Received processing result: %d", cbID, result)
	if result == 0 {
		log.Printf("[%s] Processing result is 0, something might have gone wrong!", cbID)
	}
	// Do further processing or cleanup
}

func main() {
	// Create a wait group to synchronize execution
	var wg sync.WaitGroup

	// Use unique callback function names
	processingCallback := Callback(handleProcessingResult)

	// Start multiple asynchronous processes
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(data int) {
			processData(data, processingCallback)
			wg.Done()
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	log.Println("All processing completed.")
}
