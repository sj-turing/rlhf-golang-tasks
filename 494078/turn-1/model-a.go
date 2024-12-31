package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Simulated database operation
func simulateDatabaseOperation(wg *sync.WaitGroup, dbChannel chan bool) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Simulate successful database operation
	fmt.Println("Database operation completed")

	// Signal that the database operation is done
	dbChannel <- true
}

// HTTP handler to process requests
func processRequest(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	dbChannel := make(chan bool, 10) // Limit to 10 concurrent database operations

	start := time.Now()

	// Spawn multiple goroutines to simulate concurrent database operations
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go simulateDatabaseOperation(wg, dbChannel)
	}

	// Wait for all database operations to complete
	wg.Wait()

	end := time.Now()

	// Respond to the HTTP request
	fmt.Fprintf(w, "All operations completed in %v\n", end.Sub(start))
}

func main() {
	http.HandleFunc("/process", processRequest)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
