package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"sync"
	"time"
)

const (
	lokiURL       = "http://loki:3100" // Replace with your Loki URL
	numWorkers    = 10                 // Number of workers for concurrent processing
	batchSize     = 100                // Size of log entry batches
	channelBuffer = 1000               // Buffer size for log entry channels
)

// Worker for processing file change events
func worker(wg *sync.WaitGroup, logChan <-chan fsnotify.Event, errChan chan<- error) {
	defer wg.Done()
	for event := range logChan {
		fmt.Println("event:", event)
		if event.Op&fsnotify.Write == fsnotify.Write {
			fmt.Println("modified file:", event.Name)
			// Read file changes and send to Loki
			if err := sendToLoki(event.Name); err != nil {
				errChan <- err
			}
		}
	}
}

// Function to send updates to Loki
func sendToLoki(fileName string) error {
	// Logic to read the file and send new logs to Loki
	// For simplicity, we'll just print the file name here
	fmt.Println("Sending log data to Loki for:", fileName)
	time.Sleep(time.Second) // Simulate some processing time
	return nil
}

func main() {
	logChan := make(chan fsnotify.Event, channelBuffer)
	errChan := make(chan error)
	var wg sync.WaitGroup

	// Start workers for concurrent processing
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, logChan, errChan)
	}

	// Monitor log files in batches of 100
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logChan <- event
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				errChan <- err
			}
		}
	}()

	// Add directories containing log files to the watcher
	logDir := "/path/to/log/dir"
	if err := watcher.Add(logDir); err != nil {
		log.Fatal(err)
	}

	// Wait for all workers to complete and handle errors
	wg.Wait()
	close(logChan)
	for err := range errChan {
		log.Println("Error:", err)
	}
}
