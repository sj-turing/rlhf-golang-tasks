package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents a time-based event
type Event struct {
	ID        string
	Timestamp time.Time
	Processed bool
}

// processEvent processes the event if it's within the specified time window
func processEvent(event Event, timeWindow time.Duration) bool {
	now := time.Now()
	if now.Sub(event.Timestamp) <= timeWindow {
		event.Processed = true
		fmt.Printf("Event %s processed at %s\n", event.ID, now)
		return true
	}
	fmt.Printf("Event %s expired at %s\n", event.ID, now)
	return false
}

func main() {
	var wg sync.WaitGroup
	var events []Event

	// Simulate creating events
	for i := 0; i < 5; i++ {
		// Create an event with a random timestamp
		event := Event{
			ID:        fmt.Sprintf("event-%d", i),
			Timestamp: time.Now().Add(-time.Duration(10*(i+1)) * time.Second),
			Processed: false,
		}
		events = append(events, event)
	}

	// Time window for processing events (e.g., 20 seconds)
	timeWindow := 20 * time.Second

	// Process each event in a separate goroutine
	for _, event := range events {
		wg.Add(1)
		go func(ev Event) {
			defer wg.Done()
			processEvent(ev, timeWindow)
		}(event)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
