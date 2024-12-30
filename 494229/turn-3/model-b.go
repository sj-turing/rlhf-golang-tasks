package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	eventBufferSize = 1000
	maxEventSize    = 1024 // Adjust the maximum event size as needed
)

// Define the Event struct to hold the event data
type Event struct {
	Timestamp time.Time
	Data      []byte
}

// Define the Logger struct to manage the event buffer
type Logger struct {
	buffer     []Event
	bufferSize int
	mu         sync.Mutex
	wg         sync.WaitGroup
	full       bool
	dropped    int
	eventChan  chan Event
}

// NewLogger creates a new Logger with an initial buffer size and starts a goroutine to handle events
func NewLogger() *Logger {
	l := &Logger{
		buffer:     make([]Event, 0, eventBufferSize),
		bufferSize: eventBufferSize,
		eventChan:  make(chan Event),
	}
	l.wg.Add(1)
	go l.handleEvents()
	return l
}

// LogEvent appends an event to the buffer or drops it if the buffer is full
func (l *Logger) LogEvent(event Event) {
	select {
	case l.eventChan <- event: // Non-blocking send to the event channel
	default:
		// Event buffer is full, drop the event
		l.mu.Lock()
		l.dropped++
		l.mu.Unlock()
	}
}

// handleEvents is a goroutine that reads events from the channel and processes them
func (l *Logger) handleEvents() {
	defer l.wg.Done()
	for event := range l.eventChan {
		l.mu.Lock()
		if len(l.buffer) < l.bufferSize {
			l.buffer = append(l.buffer, event)
		} else {
			// Buffer is full, drop the event
			l.dropped++
		}
		l.mu.Unlock()
	}
}

// FlushEvents writes the buffer events to the console and clears the buffer
func (l *Logger) FlushEvents() {
	l.mu.Lock()
	fmt.Println("Flushing events:")
	for _, event := range l.buffer {
		fmt.Printf("[%s] %s\n", event.Timestamp.Format("2006-01-02 15:04:05"), string(event.Data))
	}
	l.buffer = l.buffer[:0] // Clear the buffer
	l.mu.Unlock()
}

// Close stops the event handler goroutine and waits for it to finish
func (l *Logger) Close() {
	close(l.eventChan)
	l.wg.Wait()
	fmt.Printf("Dropped events: %d\n", l.dropped)
}

func main() {
	logger := NewLogger()

	// Generate and log events continuously
	go func() {
		for {
			time.Sleep(time.Millisecond * 10)
			eventData := []byte(fmt.Sprintf("Event at %s", time.Now().Format("2006-01-02 15:04:05")))
			// Truncate the event data if it exceeds the maximum size
			if len(eventData) > maxEventSize {
				eventData = eventData[:maxEventSize]
			}
			logger.LogEvent(Event{Timestamp: time.Now(), Data: eventData})
		}
	}()

	// Flush events periodically
	go func() {
		for {
			time.Sleep(time.Second * 5)
			logger.FlushEvents()
		}
	}()

	// Run for a specified duration
	time.Sleep(time.Second * 30)

	// Stop the logger and print dropped events
	logger.Close()
}
