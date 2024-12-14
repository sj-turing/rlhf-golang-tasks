package main

import (
	"log"
	"sync"
	"time"
)

// LogMessage struct to hold log information
type LogMessage struct {
	Level   string // e.g., "INFO", "ERROR"
	Message string
}

// Logger manages asynchronous logging
type Logger struct {
	mu      sync.Mutex
	logChan chan LogMessage
	wg      sync.WaitGroup
}

// NewLogger creates a new asynchronous logger
func NewLogger() *Logger {
	return &Logger{
		logChan: make(chan LogMessage, 1000), // Buffer size to prevent blocking
	}
}

// Start starts the logging goroutine
func (l *Logger) Start() {
	l.wg.Add(1)
	go l.loggingGoroutine()
}

// Stop stops the logging goroutine and waits for it to finish
func (l *Logger) Stop() {
	l.mu.Lock()
	defer l.mu.Unlock()

	close(l.logChan)
	l.wg.Wait()
}

// loggingGoroutine listens on the log channel and writes logs asynchronously
func (l *Logger) loggingGoroutine() {
	for logMessage := range l.logChan {
		log.Printf("[%s] %s\n", logMessage.Level, logMessage.Message)
	}
	l.wg.Done()
}

// Log sends a log message to the logger
func (l *Logger) Log(level, message string) {
	l.logChan <- LogMessage{Level: level, Message: message}
}

func main() {
	logger := NewLogger()
	logger.Start()

	// Simulate data stream processing with logging
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		logger.Log("INFO", "Processing item", i)
	}

	// Simulate application shutdown
	logger.Stop()
	log.Println("Application shutdown complete.")
}
