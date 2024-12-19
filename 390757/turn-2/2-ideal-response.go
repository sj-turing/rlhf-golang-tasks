package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type event struct {
	id        int
	timestamp time.Time
	action    func()
}

var (
	events  []event
	eventMu sync.Mutex
)

// addEvent adds a new event to the priority queue
func addEvent(timestamp time.Time, action func()) {
	eventMu.Lock()
	defer eventMu.Unlock()
	events = append(events, event{
		id:        len(events),
		timestamp: timestamp,
		action:    action,
	})
	// Sort events based on timestamp
	sort.Slice(events, func(i, j int) bool {
		return events[i].timestamp.Before(events[j].timestamp)
	})
}

// processEvents processes events from the priority queue
func processEvents() {
	for {
		if len(events) == 0 {
			time.Sleep(1 * time.Second) // Wait for new events
			continue
		}
		eventMu.Lock()
		if events[0].timestamp.Before(time.Now()) {
			event := events[0]
			events = events[1:]
			eventMu.Unlock()
			event.action()
		} else {
			eventMu.Unlock()
			time.Sleep(events[0].timestamp.Sub(time.Now()))
		}
	}
}

func main() {
	go processEvents()

	addEvent(time.Now().Add(1*time.Second), func() { fmt.Println("Event 1 triggered at", time.Now()) })
	addEvent(time.Now().Add(3*time.Second), func() { fmt.Println("Event 2 triggered at", time.Now()) })
	addEvent(time.Now().Add(2*time.Second), func() { fmt.Println("Event 3 triggered at", time.Now()) })

	select {} // Block the main thread
}
