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
	buffer []Event
	mu     sync.Mutex
}

// NewLogger creates a new Logger with an initial buffer size
func NewLogger(bufferSize int) *Logger {
	return &Logger{
		buffer: make([]Event, 0, bufferSize),
	}
}

// LogEvent appends an event to the buffer
func (l *Logger) LogEvent(event Event) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buffer = append(l.buffer, event)
}

// FlushEvents writes the buffer events to the console and clears the buffer
func (l *Logger) FlushEvents() {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Println("Flushing events:")
	for _, event := range l.buffer {
		fmt.Printf("[%s] %s\n", event.Timestamp.Format("2006-01-02 15:04:05"), string(event.Data))
	}
	l.buffer = l.buffer[:0] // Clear the buffer
}

func main() {
	logger := NewLogger(100) // Create a logger with a buffer size of 100
	// Generate and log events continuously
	go func() {
		for {
			time.Sleep(time.Second)
			eventData := []byte(fmt.Sprintf("Event at %s", time.Now().Format("2006-01-02 15:04:05")))
			logger.LogEvent(Event{Timestamp: time.Now(), Data: eventData})
		}
	}()
	// Flush events periodically
	for {
		time.Sleep(5 * time.Second)
		logger.FlushEvents()
	}
}
