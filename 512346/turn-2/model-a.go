package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const lokiURL = "http://your-loki-url:3100/loki/api/v1/push"

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	events := make(chan fsnotify.Event, 1000) // Buffered channel for handling events

	var wg sync.WaitGroup

	// Start a dispatcher goroutine to handle log export to Loki
	go func() {
		for event := range events {
			wg.Add(1)
			go func(e fsnotify.Event) {
				defer wg.Done()
				// Process file changes
				fmt.Printf("Processing file change: %s\n", e.Name)
				processFileChange(e.Name)
			}(event)
		}
	}()

	// Walk the log directory and add watchers for each file
	logDir := "/path/to/log/dir"
	err = filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Start watching for file system events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					events <- event
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	<-done

	close(events) // Close channel when done
	wg.Wait()     // Wait for all goroutines to finish
}

// processFileChange reads changes from a file and sends them to Loki
func processFileChange(fileName string) {
	// Logic to read log changes from the file
	// Example: tail the file or read new lines since last update

	// Send data to Loki
	sendToLoki(fileName)
}

// sendToLoki sends log data to the specified Loki URL
func sendToLoki(fileName string) {
	// Read new lines from the file or changes
	// Prepare the payload for Loki

	// Example:
	fmt.Printf("Sending log data from %s to Loki at %s\n", fileName, lokiURL)
	// Make HTTP POST request to Loki API
}
