package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataSet struct {
	ID    int
	Value int
}

// Simulate fetching data from the database for each task
func fetchData(taskID int, dataCh chan<- DataSet) {
	// Simulate database query latency and fetch random data
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	numElements := rand.Intn(10) + 1

	for i := 0; i < numElements; i++ {
		dataCh <- DataSet{
			ID:    taskID*100 + i,
			Value: rand.Intn(100),
		}
	}
}

// Aggregate the fetched data concurrently
func aggregateData(dataCh <-chan DataSet, aggData *AggregatedData) {
	for data := range dataCh {
		// Simulate data aggregation process
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		// Update the aggregated data
		mutex.Lock()
		aggData.Count += 1
		aggData.Sum += data.Value
		mutex.Unlock()
	}
}

type AggregatedData struct {
	Count int
	Sum   int
}

var mutex sync.Mutex

func main() {
	const numTasks = 5

	// Channel to pass data between fetchData and aggregateData
	aggregatedData := &AggregatedData{}
	dataCh := make(chan DataSet)

	var wg sync.WaitGroup
	wg.Add(numTasks)

	// Launch data fetching tasks concurrently
	for i := 0; i < numTasks; i++ {
		go func() {
			defer wg.Done()
			fetchData(i, dataCh)
		}()
	}

	// Launch aggregation goroutine
	go aggregateData(dataCh, aggregatedData)

	// Wait for all aggregation to complete
	wg.Wait()

	// Close the channel to indicate no more data will be sent
	close(dataCh)

	// Print the final aggregated result
	fmt.Println("--- Final Aggregated Result ---")
	fmt.Println("Final Count:", aggregatedData.Count)
	fmt.Println("Final Sum:  ", aggregatedData.Sum)
}
