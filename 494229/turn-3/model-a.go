package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Define the Event struct to hold the event data
type Event struct {
	Timestamp time.Time
	Data      []byte
}

// Channel buffer to hold events before logging
type eventChannel struct {
	buffer     []Event // Use a slice of Events
	bufferSize int
	nextIndex  int32 // Atomic index to append and fetch events
	mu         sync.Mutex
	condition  *sync.Cond
}

func newEventChannel(bufferSize int) *eventChannel {
	return &eventChannel{
		buffer:     make([]Event, 0, bufferSize),
		bufferSize: bufferSize,
		nextIndex:  0,
		mu:         sync.Mutex{},
		condition:  sync.NewCond(&sync.Mutex{}),
	}
}

func (ch *eventChannel) producer(event *Event) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	i := atomic.AddInt32(&ch.nextIndex, 1) - 1
	ch.buffer = append(ch.buffer, *event)
	if len(ch.buffer) > ch.bufferSize {
		ch.buffer = ch.buffer[i+1:] // Circular buffer logic
	}
	ch.condition.Signal()
}

func (ch *eventChannel) consumer() {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	for {
		for i := len(ch.buffer); i == 0; i = len(ch.buffer) {
			ch.condition.Wait()
		}

		for len(ch.buffer) > 0 {
			i := atomic.AddInt32(&ch.nextIndex, 1) - 1
			event := ch.buffer[i]
			ch.buffer = ch.buffer[:i] + ch.buffer[i+1:] // Remove the event from the buffer

			fmt.Printf("[%s] %s\n", event.Timestamp.Format("2006-01-02 15:04:05"), string(event.Data))
		}
	}
}

func main() {
	// Initialize the event channel
	eventChannel := newEventChannel(100)

	// Generate and log events continuously (producers)
	go func() {
		for {
			time.Sleep(time.Millisecond * 50)
			eventData := []byte(fmt.Sprintf("Event at %s", time.Now().Format("2006-01-02 15:04:05")))
			event := Event{Timestamp: time.Now(), Data: eventData}
			eventChannel.producer(&event)
		}
	}()

	// Flush events (consumers)
	go func() {
		for {
			time.Sleep(time.Second)
			eventChannel.consumer()
		}
	}()

	select {}
}
