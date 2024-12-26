package main

import (
	"fmt"
	"sync"
	"time"
)

// Aggregator represents the event-driven data aggregator
type Aggregator struct {
	Data      map[string]int
	Mutex     sync.Mutex
	WaitGroup sync.WaitGroup
}

// NewAggregator returns an initialized Aggregator instance
func NewAggregator() *Aggregator {
	return &Aggregator{
		Data: make(map[string]int),
	}
}

// AddData is a method to add data to the aggregator
func (a *Aggregator) AddData(key string, value int) {
	a.Mutex.Lock()
	a.Data[key] += value
	a.Mutex.Unlock()
}

// StartAggregation initiates parallel data aggregation tasks
func (a *Aggregator) StartAggregation(dataSources map[string]func() int) {
	a.WaitGroup.Add(len(dataSources))
	for sourceName, dataFunc := range dataSources {
		go a.aggregateData(sourceName, dataFunc)
	}
}

func (a *Aggregator) aggregateData(sourceName string, dataFunc func() int) {
	defer a.WaitGroup.Done()
	data := dataFunc()
	a.AddData(sourceName, data)
}

// WaitForCompletion blocks the main goroutine until all data aggregation tasks are finished
func (a *Aggregator) WaitForCompletion() {
	a.WaitGroup.Wait()
}

// Example data sources
func dataSource1() int {
	time.Sleep(time.Second * 2) // Simulate data fetching
	return 100
}

func dataSource2() int {
	time.Sleep(time.Second * 1) // Simulate data fetching
	return 50
}

func dataSource3() int {
	time.Sleep(time.Second * 3) // Simulate data fetching
	return 150
}

func main() {
	aggregator := NewAggregator()
	dataSources := map[string]func() int{
		"Source1": dataSource1,
		"Source2": dataSource2,
		"Source3": dataSource3,
	}

	aggregator.StartAggregation(dataSources)
	aggregator.WaitForCompletion()

	// Display the aggregated results
	fmt.Println("Aggregated Data:")
	for source, value := range aggregator.Data {
		fmt.Printf("%s: %d\n", source, value)
	}
}
