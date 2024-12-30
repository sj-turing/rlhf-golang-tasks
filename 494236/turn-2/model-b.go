package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

func main() {
	taskCh := make(chan Task, 100)
	resultCh := make(chan int, 100)

	// Method 1: Plain Range Loop
	go func() {
		for task := range taskCh {
			enrichedData := enrichData(task.Data)
			resultCh <- len(enrichedData)
		}
	}()

	// Method 2: Using Goroutines
	const concurrencyLevel = 10
	var wg sync.WaitGroup
	wg.Add(concurrencyLevel)

	for i := 0; i < concurrencyLevel; i++ {
		go func() {
			for task := range taskCh {
				enrichedData := enrichData(task.Data)
				resultCh <- len(enrichedData)
			}
			wg.Done()
		}()
	}

	// Method 3: Streaming with Go Channels
	go func() {
		for {
			task := Task{ID: len(tasks), Data: generateRandomData()}
			taskCh <- task
			tasks = append(tasks, task)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	var results []int
	for result := range resultCh {
		results = append(results, result)
		if len(results) == len(tasks) {
			break
		}
	}

	// Consolidate results
	for _, result := range results {
		totalEnrichedCharacters += result
	}

	fmt.Printf("Total enriched characters: %d\n", totalEnrichedCharacters)
	close(taskCh)
	wg.Wait()
}

func enrichData(data string) string {
	// Some expensive enrichment logic goes here
	time.Sleep(100 * time.Millisecond)
	return data + " - Enriched!"
}

func generateRandomData() string {
	// Function to generate random data
	return ""
}
