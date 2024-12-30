package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the Event struct to hold the event data
type Event struct {
	Timestamp time.Time
	Data      []byte
}

// Define the Logger struct to manage the event buffer
type Logger struct {
	buffer     []*Event // Use a slice of pointers to Event to reuse memory
	bufferSize int
	freeList   *sync.Pool // Use a sync.Pool to reuse event structs
	mu         sync.Mutex
	cond       *sync.Cond
	full       bool
	empty      bool
}

// NewLogger creates a new Logger with an initial buffer size
func NewLogger(bufferSize int) *Logger {
	l := &Logger{
		buffer:     make([]*Event, 0, bufferSize),
		bufferSize: bufferSize,
		freeList:   &sync.Pool{New: func() interface{} { return &Event{} }},
	}
	l.cond = sync.NewCond(&l.mu)
	return l
}

// LogEvent appends an event to the buffer or waits if the buffer is full
func (l *Logger) LogEvent(event *Event) {
	l.mu.Lock()
	for l.full {
		l.cond.Wait() // Block until the buffer is not full
	}
	l.buffer = append(l.buffer, event)
	l.empty = false
	if len(l.buffer) == l.bufferSize {
		l.full = true
	}
	l.cond.Signal() // Signal to consumers that there is an event
	l.mu.Unlock()
}

// FlushEvents writes the buffer events to the console and clears the buffer
func (l *Logger) FlushEvents() {
	l.mu.Lock()
	for l.empty {
		l.cond.Wait() // Block until the buffer is not empty
	}
	fmt.Println("Flushing events:")
	for _, event := range l.buffer {
		fmt.Printf("[%s] %s\n", event.Timestamp.Format("2006-01-02 15:04:05"), string(event.Data))
		l.freeList.Put(event) // Put the event back into the free list for reuse
	}
	l.buffer = l.buffer[:0] // Clear the buffer
	l.full = false
	l.cond.Signal() // Signal to producers that there is space in the buffer
	l.mu.Unlock()
}

func main() {
	bufferSize := 1000 // Adjust the buffer size as needed
	logger := NewLogger(bufferSize)

	// Generate and log events continuously
	go func() {
		for {
			time.Sleep(time.Millisecond * 10) // Reduce event generation rate to keep up with flushing
			eventData := []byte(fmt.Sprintf("Event at %s", time.Now().Format("2006-01-02 15:04:05")))
			event := logger.freeList.Get().(*Event) // Get an event from the free list
			event.Timestamp = time.Now()
			event.Data = eventData
			logger.LogEvent(event)
		}
	}()

	// Flush events periodically
	go func() {
		for {
			time.Sleep(time.Second * 5)
			logger.FlushEvents()
		}
	}()
	select {}
}
