package main

import (
	"fmt"
	"sync"
)

// Event represents an event that observers are interested in
type Event interface {
	Name() string
}

// SimpleEvent is a simple implementation of the Event interface
type SimpleEvent struct {
	name string
}

// NewSimpleEvent take a name and returns SimpleEvent
func NewSimpleEvent(name string) *SimpleEvent {
	return &SimpleEvent{
		name: name,
	}
}

// Name returns a name of an event
func (se *SimpleEvent) Name() string {
	return se.name
}

// Observer represents an observer of events
type Observer interface {
	Notify(event Event)
}

// EventPlanner is the main structure that handles event registration and notification
type EventPlanner struct {
	observers []Observer
	mu        sync.Mutex
}

// RegisterObserver adds an observer to the event planner
func (ep *EventPlanner) RegisterObserver(observer Observer) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.observers = append(ep.observers, observer)
}

// NotifyObservers notifies all observers about the specified event
func (ep *EventPlanner) NotifyObservers(event Event) {
	// Creating a WaitGroup to send notification concurrently
	var wg sync.WaitGroup

	// total number of observers
	wg.Add(len(ep.observers))

	for _, observer := range ep.observers {
		// concurrently sending notifactions
		go func(event Event) {
			defer wg.Done()
			observer.Notify(event)
		}(event)
	}

	// Waiting till all notification sent
	wg.Wait()
}

// Person represents a person who is an observer of the event planner
type Person struct {
	name string
}

// Notify implements the Observer interface and notifies the person about the event
func (p *Person) Notify(event Event) {
	fmt.Printf("%s is notified about the event: %s\n", p.name, event.Name())
}

func main() {
	planner := EventPlanner{}

	// Register observers
	planner.RegisterObserver(&Person{name: "Alice"})
	planner.RegisterObserver(&Person{name: "Bob"})
	planner.RegisterObserver(&Person{name: "Charlie"})

	// Plan an event
	event := NewSimpleEvent("Birthday Party")

	// Notify all observers about the event
	planner.NotifyObservers(event)

	// Output:
	// Alice is notified about the event: Birthday Party
	// Bob is notified about the event: Birthday Party
	// Charlie is notified about the event: Birthday Party
}
