package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Shared data structure
type Counter struct {
	count int
}

// Operation represents a command to operate on the counter
type Operation string

const (
	IncrementOp Operation = "increment"
	GetCountOp  Operation = "get_count"
)

// Process a single operation on the counter
func (c *Counter) processOperation(op Operation) {
	switch op {
	case IncrementOp:
		c.count++
	case GetCountOp:
	default:
		fmt.Println("Invalid operation")
	}
}

// Worker goroutine that processes operations from the channel
func counterWorker(c *Counter, ops chan Operation) {
	for op := range ops {
		c.processOperation(op)
	}
	close(ops)
}

// Global counter instance
var counter Counter
var opsChannel = make(chan Operation, 100)
var wg sync.WaitGroup

// HTTP handler to increment the counter
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	opsChannel <- IncrementOp
	fmt.Fprintf(w, "Counter increment operation sent. Current count: %d\n", counter.count)
}

// HTTP handler to get the current count
func getCountHandler(w http.ResponseWriter, r *http.Request) {
	opsChannel <- GetCountOp
	fmt.Fprintf(w, "Count get operation sent. Current count: %d\n", counter.count)
}

func main() {
	// Start worker goroutine
	wg.Add(1)
	go func() {
		counterWorker(&counter, opsChannel)
		wg.Done()
	}()

	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/count", getCountHandler)

	// Start the HTTP server
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Failed to start server:", err)
		}
	}()

	// Simulate concurrent requests
	for i := 0; i < 10; i++ {
		go func() {
			http.Get("http://localhost:8080/increment")
			http.Get("http://localhost:8080/count")
		}()
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to finish

	// Shut down the worker goroutine
	close(opsChannel)
	wg.Wait()

	// Get the final count
	fmt.Println("Final count:", counter.count)
}
