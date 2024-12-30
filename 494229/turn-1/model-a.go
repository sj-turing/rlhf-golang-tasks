package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	bufferSize = 1000 // Adjust based on memory constraints and expected throughput
	logEntry   = "Event: %s at %s\n"
)

type logEntryStruct struct {
	message   string
	timestamp time.Time
}

type logBuffer struct {
	entries []logEntryStruct
	mu      sync.Mutex
	head    int
	tail    int
}

func newLogBuffer() *logBuffer {
	return &logBuffer{entries: make([]logEntryStruct, bufferSize)}
}

func (b *logBuffer) write(entry logEntryStruct) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.entries[b.tail] = entry
	b.tail = (b.tail + 1) % bufferSize
	if b.tail == b.head {
		b.head = (b.head + 1) % bufferSize
	}
}

func (b *logBuffer) read() (logEntryStruct, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.head == b.tail {
		return logEntryStruct{}, false // Buffer is empty
	}

	entry := b.entries[b.head]
	b.head = (b.head + 1) % bufferSize
	return entry, true
}

func logProducer(buffer *logBuffer) {
	for {
		message := fmt.Sprintf("Example event from producer: %d", time.Now().UnixNano())
		entry := logEntryStruct{message: message, timestamp: time.Now()}
		buffer.write(entry)
		time.Sleep(100 * time.Millisecond) // Simulate event generation
	}
}

func logConsumer(buffer *logBuffer, conn net.Conn) {
	defer conn.Close()
	writer := bufio.NewWriter(conn)

	for {
		entry, ok := buffer.read()
		if !ok {
			time.Sleep(10 * time.Millisecond) // Wait for new entries
			continue
		}

		logString := fmt.Sprintf(logEntry, entry.message, entry.timestamp)
		_, err := writer.WriteString(logString)
		if err != nil {
			log.Println("Error writing to network:", err)
			break
		}

		err = writer.Flush()
		if err != nil {
			log.Println("Error flushing buffer to network:", err)
			break
		}
	}
}

func main() {
	buffer := newLogBuffer()

	// Connect to a network server (replace with actual server address)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Start producer and consumer goroutines
	go logProducer(buffer)
	go logConsumer(buffer, conn)

	select {} // Block main goroutine to keep the program running
}
