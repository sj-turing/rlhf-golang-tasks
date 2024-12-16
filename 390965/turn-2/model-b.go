package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

// Logger provides asynchronous logging with minimal allocations.
type Logger struct {
	logChan chan []byte
	wg      sync.WaitGroup
	pool    *sync.Pool
}

func NewLogger(bufferSize int) *Logger {
	l := &Logger{
		logChan: make(chan []byte, bufferSize),
		pool:    &sync.Pool{New: func() interface{} { return make([]byte, 256) }},
	}
	l.wg.Add(1)
	go l.startLogging()
	return l
}

func (l *Logger) startLogging() {
	defer l.wg.Done()
	for logMsg := range l.logChan {
		// Write log message to standard output
		bw := bufio.NewWriter(os.Stdout)
		_, _ = bw.Write(logMsg)
		_, _ = bw.WriteString("\n")
		_ = bw.Flush()
		// Release the buffer back to the pool
		l.pool.Put(logMsg)
	}
}

func (l *Logger) Log(format string, args ...interface{}) {
	buf := l.pool.Get().([]byte)
	defer l.pool.Put(buf)

	// Use fmt.Sprintf to minimize allocations; you could optimize further by
	// using manual appends for fixed format logs.
	msg := fmt.Sprintf(format+"\n", args...)
	lenMsg := len(msg)

	// Ensure the buffer is large enough
	if len(buf) < lenMsg {
		buf = make([]byte, lenMsg)
	}

	copy(buf, msg)
	l.logChan <- buf
}

func (l *Logger) Close() {
	close(l.logChan)
	l.wg.Wait()
}

func main() {
	logger := NewLogger(10)

	// Simulating concurrent log writes
	for i := 0; i < 10000; i++ {
		logger.Log("Log message #%d at %s", i, time.Now())
		time.Sleep(10 * time.Microsecond) // Simulate some processing time
	}

	// Ensure all logs are flushed before exiting
	logger.Close()
}
