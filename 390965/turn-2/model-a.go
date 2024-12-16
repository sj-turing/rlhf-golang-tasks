package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	// Control the buffer size of the log channel
	logChannelSize = 100
)

type LogLevel uint8

const (
	InfoLevel LogLevel = iota
	WarnLevel
	ErrorLevel
)

type LogMessage struct {
	Level   LogLevel
	Message string
	Time    time.Time
}

// Logger encapsulates logging functionality
type Logger struct {
	logChan chan LogMessage
	wg      sync.WaitGroup
}

// NewLogger instantiates a new Logger
func NewLogger() *Logger {
	l := &Logger{
		logChan: make(chan LogMessage, logChannelSize),
	}
	l.wg.Add(1)
	go l.processLogs()
	return l
}

// processLogs is a goroutine that processes log messages
func (l *Logger) processLogs() {
	defer l.wg.Done()
	for logMsg := range l.logChan {
		l.writeLog(&logMsg)
	}
}

func (l *Logger) writeLog(logMsg *LogMessage) {
	var level string
	switch logMsg.Level {
	case InfoLevel:
		level = "INFO"
	case WarnLevel:
		level = "WARN"
	case ErrorLevel:
		level = "ERROR"
	}
	fmt.Printf("[%s] %s: %s\n", logMsg.Time.Format(time.RFC3339), level, logMsg.Message)
}

// Log asynchronously emits a log message
func (l *Logger) Log(level LogLevel, message string) {
	logMsg := LogMessage{
		Level:   level,
		Message: message,
		Time:    time.Now(),
	}
	select {
	case l.logChan <- logMsg:
	default:
		// Log channel is full, consider dropping the log or handling it accordingly
		log.Printf("Log channel is full, dropped message: %s", message)
	}
}

func (l *Logger) Close() {
	close(l.logChan)
	l.wg.Wait()
}

func main() {
	logger := NewLogger() // Create a new async logger

	// Simulating concurrent log writes
	for i := 0; i < 100; i++ {
		go logger.Log(InfoLevel, fmt.Sprintf("Log message #%d", i))
		time.Sleep(10 * time.Millisecond) // Simulate some processing time
	}

	// Give some time for all logs to be processed before exiting
	time.Sleep(2 * time.Second)
	logger.Close() // Ensure all logs are flushed before exiting
}
