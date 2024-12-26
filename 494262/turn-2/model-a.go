package main

/*

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// A simple structure to hold data that will be processed concurrently.
type Data struct {
	ID        int
	Value     string
	Processed bool
}

// A slice of Data structures to simulate a collection of data.
var data = []Data{
	{1, "data1", false},
	{2, "data2", false},
	{3, "data3", false},
	{4, "data4", false},
	{5, "data5", false},
	// Add more data as needed
}

// Function to process a single Data item.
func processData(d *Data, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate processing
	time.Sleep(time.Duration(d.ID) * time.Millisecond)

	// Mark data as processed
	atomic.StoreBool(&d.Processed, true)
}

func main() {
	var wg sync.WaitGroup

	// Iterate over the data slice
	for _, d := range data {
		// Create a copy of the loop variable to avoid closure issues
		dataCopy := d
		wg.Add(1)
		go func() {
			processData(&dataCopy, &wg)
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Check and print processed data
	fmt.Println("Processed data:")
	for _, d := range data {
		if d.Processed {
			fmt.Printf("ID: %d, Value: %s, Processed: %t\n", d.ID, d.Value, d.Processed)
		}
	}
}
*/
