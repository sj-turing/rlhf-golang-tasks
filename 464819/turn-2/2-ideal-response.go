package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	Notify(ctx context.Context, event Event)
}

// eventPlannerChan is used to communicate event registrations and notifications.
type eventPlannerChan chan Observer

// EventPlanner is the main structure that handles event registration and notification
type EventPlanner struct {
	register eventPlannerChan
}

// NewEventPlanner creates a new EventPlanner with two buffered channels for registration and notification.
func NewEventPlanner() *EventPlanner {
	ep := &EventPlanner{
		register: make(eventPlannerChan, 10),
	}
	return ep
}

// RegisterObserver registers an observer in a separate goroutine to handle it asynchronously.
func (ep *EventPlanner) RegisterObserver(observer Observer) {
	ep.register <- observer
}

// StartEventPlanner goroutine runs continuously, listening for registrations and notifications.
func (ep *EventPlanner) NotifyObservers(ctx context.Context, event Event) {
	for {
		select {
		case person := <-ep.register:
			person.Notify(ctx, event)
		case <-ctx.Done():
			fmt.Println("Cancelling the context")
			return
		}
	}
}

// Close gracefully shuts down the EventPlanner and its worker goroutine.
func (ep *EventPlanner) Close() {
	close(ep.register)
}

// Person represents a person who is an observer of the event planner
type Person struct {
	name string
}

// NewPerson returns a pointer of Person struct
func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

// Notify implements the Observer interface and notifies the person about the event.
func (p *Person) Notify(ctx context.Context, event Event) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	time.Sleep(300 * time.Millisecond) // Simulate a task that takes time
	fmt.Printf("%s is notified about the event: %s\n", p.name, event.Name())

}

func main() {

	// creating a context to deal with asynchronous processing
	ctx, cancel := context.WithCancel(context.Background())

	// buffered sigChan to listening on interrupt
	sigChan := make(chan os.Signal, 1)

	// notifies if receives syscall.SIGINT
	signal.Notify(sigChan, syscall.SIGINT)

	// creating a new event planner
	planner := NewEventPlanner()

	// persons to send invitations
	// this can be database data
	persons := []string{"Alice", "Bob", "Charlie"}

	// registering person asynchronously
	for _, name := range persons {
		go planner.RegisterObserver(NewPerson(name))
	}

	// creating an event
	event := NewSimpleEvent("Birthday Party")

	// send notification to asynchronously
	go planner.NotifyObservers(ctx, event)

	// listening on signal chan
	// if receives terms signal then stop
	// the execution of program
	<-sigChan
	cancel()

	// closing the channel
	planner.Close()

	fmt.Println("Shutting down event planner")
}
