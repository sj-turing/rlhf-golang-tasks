package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Callback context to pass additional information
type CallbackContext struct {
	ID      int    // Unique identifier for this callback
	Desc    string // Description or label
	Timeout time.Duration
}

// A function that accepts a callback
func processTask(ctx *CallbackContext, callback func(string, *CallbackContext), wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Starting task %d with desc: %s, timeout: %s\n", ctx.ID, ctx.Desc, ctx.Timeout.String())

	// Simulate some work with a timeout
	select {
	case <-time.After(ctx.Timeout):
		log.Printf("Task %d timed out after %s\n", ctx.ID, ctx.Timeout.String())
		callback("Timed out", ctx)
	default:
		// Simulate work
		time.Sleep(time.Second)
		result := fmt.Sprintf("Task %d completed", ctx.ID)
		log.Printf("Task %d finished: %s\n", ctx.ID, result)
		callback(result, ctx)
	}
}

// Another function that accepts a callback
func processAsyncTask(ctx *CallbackContext, callback func(string, *CallbackContext), wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Starting async task %d with desc: %s, timeout: %s\n", ctx.ID, ctx.Desc, ctx.Timeout.String())

	time.Sleep(2 * time.Second) // Simulate longer work

	result := fmt.Sprintf("Async task %d completed", ctx.ID)
	log.Printf("Async task %d finished: %s\n", ctx.ID, result)
	callback(result, ctx)
}

// A simple callback function
func handleResult(result string, ctx *CallbackContext) {
	log.Printf("Callback result: %s for task %d with desc: %s\n", result, ctx.ID, ctx.Desc)
}

func main() {
	// Create a wait group to wait for all goroutines
	var wg sync.WaitGroup

	// Callback contexts
	ctx1 := &CallbackContext{ID: 1, Desc: "Normal task", Timeout: 3 * time.Second}
	ctx2 := &CallbackContext{ID: 2, Desc: "Short task", Timeout: 1 * time.Second}
	ctx3 := &CallbackContext{ID: 3, Desc: "Async task", Timeout: 5 * time.Second}

	// Start goroutines for processing tasks
	wg.Add(1)
	go processTask(ctx1, handleResult, &wg)

	wg.Add(1)
	go processTask(ctx2, handleResult, &wg)

	wg.Add(1)
	go processAsyncTask(ctx3, handleResult, &wg)

	// Wait for all goroutines to complete
	wg.Wait()
}
