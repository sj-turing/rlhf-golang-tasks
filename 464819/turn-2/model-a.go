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

// NewSimpleEvent takes a name and returns a SimpleEvent
func NewSimpleEvent(name string) *SimpleEvent {
	return &SimpleEvent{
		name: name,
	}
}

// Name returns the name of an event
func (se *SimpleEvent) Name() string {
	return se.name
}

// Observer represents an observer of events
type Observer interface {
	Notify(event Event)
}

// EventPlanner manages event registration and notification
type EventPlanner struct {
	mu            sync.Mutex
	notifications chan Event
	observers     []Observer
}

// NewEventPlanner returns a new instance of EventPlanner
func NewEventPlanner() *EventPlanner {
	return &EventPlanner{
		notifications: make(chan Event),
	}
}

// RegisterObserver adds an observer to the event planner
func (ep *EventPlanner) RegisterObserver(observer Observer) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.observers = append(ep.observers, observer)
}

// NotifyObservers notifies all observers about the specified event
func (ep *EventPlanner) NotifyObservers(event Event) {
	ep.notifications <- event
}

// Start listener which processes notifications
func (ep *EventPlanner) Start() {
	go func() {
		for event := range ep.notifications {
			// Creating a WaitGroup to send notification concurrently
			var wg sync.WaitGroup

			// total number of observers
			wg.Add(len(ep.observers))

			for _, observer := range ep.observers {
				// concurrently sending notifications
				go func(event Event, observer Observer) {
					defer wg.Done()
					observer.Notify(event)
				}(event, observer)
			}

			// Waiting till all notification sent
			wg.Wait()
		}
	}()
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
	planner := NewEventPlanner()

	// Register observers
	planner.RegisterObserver(&Person{name: "Alice"})
	planner.RegisterObserver(&Person{name: "Bob"})
	planner.RegisterObserver(&Person{name: "Charlie"})

	// Start notification listener
	planner.Start()

	// Plan an event asynchronously
	go func() {
		event := NewSimpleEvent("Birthday Party")
		planner.NotifyObservers(event)
		fmt.Println("Event published asynchronously.")
	}()

	// Keep the main program running
	select {}
}
