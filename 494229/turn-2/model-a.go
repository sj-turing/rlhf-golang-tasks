package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents a logged network event
type Event struct {
	Timestamp time.Time
	Data      []byte
}

// Logger manages the event buffer
type Logger struct {
	bufferSize int
	buffer     []Event
	mu         sync.Mutex
	cond       *sync.Cond
}

// NewLogger initializes a new Logger with a specified buffer size
func NewLogger(bufferSize int) *Logger {
	l := &Logger{bufferSize: bufferSize, buffer: make([]Event, 0, bufferSize)}
	l.cond = sync.NewCond(&l.mu)
	go l.flushLoop()
	return l
}

// LogEvent appends an event to the buffer
func (l *Logger) LogEvent(event Event) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for len(l.buffer) >= l.bufferSize {
		l.cond.Wait()
	}

	l.buffer = append(l.buffer, event)
	l.cond.Signal()
}

// flushLoop continuously flushes events
func (l *Logger) flushLoop() {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		l.mu.Lock()
		defer l.mu.Unlock()

		if len(l.buffer) == 0 {
			continue
		}

		l.cond.Broadcast()
		eventsToFlush := l.buffer
		l.buffer = l.buffer[:0]

		// Flush logic goes here, e.g., writing to disk or sending to a remote service
		fmt.Println("Flushing events:")
		for _, event := range eventsToFlush {
			fmt.Printf("[%s] %s\n", event.Timestamp.Format("2006-01-02 15:04:05"), string(event.Data))
		}
	}
}

func main() {
	logger := NewLogger(100)

	// Generate and log events
	go func() {
		for {
			eventData := []byte(fmt.Sprintf("Event at %s", time.Now().Format("2006-01-02 15:04:05")))
			logger.LogEvent(Event{Timestamp: time.Now(), Data: eventData})
			time.Sleep(50 * time.Millisecond) // Simulate event generation interval
		}
	}()

	select {} // Run forever
}
