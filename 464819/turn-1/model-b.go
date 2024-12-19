package main

import (
	"fmt"
	"sync"
)

// Event represents an event in the event planner
type Event struct {
	Name string
}

// Observer represents an observer of events
type Observer interface {
	Notify(event *Event)
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
func (ep *EventPlanner) NotifyObservers(event *Event) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	for _, observer := range ep.observers {
		go observer.Notify(event)
	}
}

// Person represents a person who is an observer of the event planner
type Person struct {
	name string
}

// Notify implements the Observer interface and notifies the person about the event
func (p *Person) Notify(event *Event) {
	fmt.Printf("%s is notified about the event: %s\n", p.name, event.Name)
}

func main() {
	planner := EventPlanner{}

	// Register observers
	planner.RegisterObserver(&Person{name: "Alice"})
	planner.RegisterObserver(&Person{name: "Bob"})
	planner.RegisterObserver(&Person{name: "Charlie"})

	// Plan an event
	event := &Event{Name: "Birthday Party"}

	// Notify all observers about the event
	planner.NotifyObservers(event)

	// Output:
	// Alice is notified about the event: Birthday Party
	// Bob is notified about the event: Birthday Party
	// Charlie is notified about the event: Birthday Party
}
