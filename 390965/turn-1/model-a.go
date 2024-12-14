package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type LogMessage struct {
	Message string
	Time    time.Time
}

type Logger struct {
	logChan chan LogMessage
	wg      sync.WaitGroup
}

func NewLogger(bufferSize int) *Logger {
	l := &Logger{
		logChan: make(chan LogMessage, bufferSize),
	}
	l.wg.Add(1)
	go l.startLogging()
	return l
}

func (l *Logger) startLogging() {
	defer l.wg.Done()
	for logMsg := range l.logChan {
		// Here you can format the log message and decide how to handle it.
		fmt.Printf("[%s] %s\n", logMsg.Time.Format(time.RFC3339), logMsg.Message)
	}
}

func (l *Logger) Log(message string) {
	logMsg := LogMessage{
		Message: message,
		Time:    time.Now(),
	}
	select {
	case l.logChan <- logMsg:
		// Log message successfully sent to channel
	default:
		// Log channel is full, consider dropping the log or handle it accordingly
		log.Printf("Log channel is full, dropped message: %s", message)
	}
}

func (l *Logger) Close() {
	close(l.logChan)
	l.wg.Wait()
}

func main() {
	logger := NewLogger(10) // Buffered channel size of 10

	// Simulating concurrent log writes
	for i := 0; i < 100; i++ {
		logger.Log(fmt.Sprintf("Log message #%d", i))
		time.Sleep(10 * time.Millisecond) // Simulate some processing time
	}

	// Ensure all logs are flushed before exiting
	logger.Close()
}
