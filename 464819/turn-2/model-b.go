package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents an event that observers are interested in
type Event interface {
	Name() string
}

// SimpleEvent is a simple implementation of the Event interface
type SimpleEvent struct {
	name string
}

// NewSimpleEvent creates a new event with the given name.
func NewSimpleEvent(name string) *SimpleEvent {
	return &SimpleEvent{name}
}

// Name returns the name of the event.
func (se *SimpleEvent) Name() string {
	return se.name
}

// Observer represents an observer of events.
type Observer interface {
	Notify(event Event)
}

// eventPlannerChan is used to communicate event registrations and notifications.
type eventPlannerChan chan interface{}

// EventPlanner is the main structure that handles event registration and notification
type EventPlanner struct {
	observers []Observer
	register  eventPlannerChan
	notify    eventPlannerChan
	mu        sync.Mutex
	wg        sync.WaitGroup
}

// NewEventPlanner creates a new EventPlanner with two buffered channels for registration and notification.
func NewEventPlanner() *EventPlanner {
	ep := &EventPlanner{
		register:  make(eventPlannerChan, 100),
		notify:    make(eventPlannerChan, 100),
		observers: make([]Observer, 0),
	}
	ep.wg.Add(1)
	go ep.startEventPlanner()
	return ep
}

// RegisterObserver registers an observer in a separate goroutine to handle it asynchronously.
func (ep *EventPlanner) RegisterObserver(observer Observer) {
	ep.register <- observer
}

// NotifyObservers notifies all observers about the specified event asynchronously.
func (ep *EventPlanner) NotifyObservers(event Event) {
	ep.notify <- event
}

// StartEventPlanner goroutine runs continuously, listening for registrations and notifications.
func (ep *EventPlanner) startEventPlanner() {
	defer ep.wg.Done()
	for {
		select {
		case msg := <-ep.register:
			ep.handleRegistration(msg)
		case msg := <-ep.notify:
			ep.handleNotification(msg)
		}
	}
}

// HandleRegistration handles observer registrations.
func (ep *EventPlanner) handleRegistration(msg interface{}) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	observer, ok := msg.(Observer)
	if !ok {
		// Handle invalid registration message
		return
	}
	ep.observers = append(ep.observers, observer)
}

// HandleNotification handles event notifications by broadcasting them to all observers.
func (ep *EventPlanner) handleNotification(msg interface{}) {
	event, ok := msg.(Event)
	if !ok {
		// Handle invalid notification message
		return
	}
	// Notify all observers in a goroutine for concurrency
	for _, observer := range ep.observers {
		go observer.Notify(event)
	}
}

// Close gracefully shuts down the EventPlanner and its worker goroutine.
func (ep *EventPlanner) Close() {
	close(ep.register)
	close(ep.notify)
	ep.wg.Wait()
}

// Person represents a person who is an observer of the event planner
type Person struct {
	name string
}

// Notify implements the Observer interface and notifies the person about the event.
func (p *Person) Notify(event Event) {
	time.Sleep(2 * time.Second) // Simulate a task that takes time
	fmt.Printf("%s is notified about the event: %s\n", p.name, event.Name())
}
