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
func fetchData(taskID int, wg *sync.WaitGroup, dataCh chan<- []DataSet) {
	defer wg.Done()
	// Simulate database query latency and fetch random data
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	numElements := rand.Intn(10) + 1
	var data []DataSet
	for i := 0; i < numElements; i++ {
		data = append(data, DataSet{
			ID:    taskID*100 + i,
			Value: rand.Intn(100),
		})
	}
	dataCh <- data
}

// Aggregate the fetched data concurrently
func aggregateData(dataCh <-chan []DataSet, aggData *AggregatedData, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		for _, value := range data {
			// Simulate data aggregation process
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			// Update the aggregated data
			mutex.Lock()
			aggData.Count += 1
			aggData.Sum += value.Value
			mutex.Unlock()
		}
	}
}

type AggregatedData struct {
	Count int
	Sum   int
}

var mutex sync.Mutex

func main() {
	const numTasks = 5
	var wg sync.WaitGroup

	// Channel to pass data between fetchData and aggregateData
	dataCh := make(chan []DataSet)
	aggregatedData := &AggregatedData{}

	// Launch data fetching tasks concurrently
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go fetchData(i, &wg, dataCh)
	}

	// Launch aggregation goroutine
	wg.Add(1)
	go aggregateData(dataCh, aggregatedData, &wg)

	// Close the channel to indicate no more data will be sent
	go func() {
		wg.Wait()
		close(dataCh)
	}()

	// Wait for all aggregation to complete
	wg.Wait()

	// Print the final aggregated result
	fmt.Println("--- Final Aggregated Result ---")
	fmt.Println("Final Count:", aggregatedData.Count)
	fmt.Println("Final Sum:  ", aggregatedData.Sum)
}
