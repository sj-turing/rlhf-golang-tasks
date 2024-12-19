package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	SensorID  string
	Timestamp time.Time
	Data      float64
}

// eventProcessor represents a worker that processes events based on their timestamps
type eventProcessor struct {
	eventChannel chan Event
	wg           sync.WaitGroup
}

func newEventProcessor(numWorkers int) *eventProcessor {
	ep := &eventProcessor{
		eventChannel: make(chan Event, 1000), // Buffered channel for better performance
	}

	// Start the specified number of workers
	for i := 0; i < numWorkers; i++ {
		ep.wg.Add(1)
		go ep.worker()
	}

	return ep
}

func (ep *eventProcessor) worker() {
	defer ep.wg.Done()
	for event := range ep.eventChannel {
		// Process the event here
		fmt.Printf("Processing event: %v, Timestamp: %v\n", event.SensorID, event.Timestamp)
		// Simulate some work
		time.Sleep(time.Millisecond * 50)
	}
}

func (ep *eventProcessor) Stop() {
	close(ep.eventChannel)
	ep.wg.Wait()
}

func main() {
	numWorkers := 4 // Number of worker goroutines to process events
	ep := newEventProcessor(numWorkers)

	// Generate and send events with timestamps
	for i := 0; i < 100; i++ {
		event := Event{
			SensorID:  fmt.Sprintf("Sensor-%d", i),
			Timestamp: time.Now().Add(time.Duration(i) * time.Second),
			Data:      float64(i),
		}
		ep.eventChannel <- event
	}

	// Wait for all events to be processed
	ep.Stop()
	fmt.Println("All events processed.")
}
