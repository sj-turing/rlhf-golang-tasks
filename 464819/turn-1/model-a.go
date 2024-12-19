package main

import (
	"fmt"
)

// Event represents an event that observers are interested in
type Event interface {
	Name() string
}

// SimpleEvent is a simple implementation of the Event interface
type SimpleEvent struct {
	Name string
}

func (se *SimpleEvent) Name() string {
	return se.Name
}

// Observer defines a contract for any object that can be notified of an event
type Observer interface {
	Notify(event Event)
}

// EventPlanner is the center of the Observer Pattern
type EventPlanner struct {
	observers []Observer
}

// RegisterObserver adds an observer to the list of observers
func (ep *EventPlanner) RegisterObserver(observer Observer) {
	ep.observers = append(ep.observers, observer)
}

// NotifyObservers sends a notification to all observers about an event
func (ep *EventPlanner) NotifyObservers(event Event) {
	for _, observer := range ep.observers {
		go observer.Notify(event)
	}
}

// Person represents a simple observer that responds to events
type Person struct {
	Name string
}

func (p *Person) Notify(event Event) {
	fmt.Println(p.Name, "received an event:", event.Name())
}

func main() {
	// Create an event planner
	ep := EventPlanner{}

	// Register some observers
	ep.RegisterObserver(&Person{Name: "Alice"})
	ep.RegisterObserver(&Person{Name: "Bob"})

	// Create an event
	event := &SimpleEvent{Name: "Party"}

	// Notify observers
	ep.NotifyObservers(event)

	// Note: Notifications may be sent concurrently due to the use of goroutines
}
