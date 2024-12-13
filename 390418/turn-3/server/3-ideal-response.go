package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	maxMessageSize = 65536 // 64 KB limit per message
	delimiter      = '\n'
)

func main() {
	// Start TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		message, err := readMessage(reader)
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected")
				return
			}
			log.Println("Error reading message:", err)
			return
		}

		// Broadcast message to other clients
		broadcastMessage(string(message), writer)
	}
}

func readMessage(reader *bufio.Reader) ([]byte, error) {
	message := []byte{}

	for {
		b, isPrefix, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		if isPrefix {
			return nil, fmt.Errorf("message exceeds %d bytes", maxMessageSize)
		}

		message = b

		// Check if we've reached the end of a message
		if len(message) > 0 {
			break
		}
	}

	return message, nil // Remove the delimiter
}

func broadcastMessage(message string, writer *bufio.Writer) {
	// Assuming there's a way to get all connected clients
	// For simplicity, we'll just log the message here
	log.Printf("Received: %s", message)

	// Write message back to the sender (for testing purposes)
	_, err := writer.WriteString(message + "\n")
	if err != nil {
		log.Println("Error writing message:", err)
	}
	err = writer.Flush()
	if err != nil {
		log.Println("Error flushing writer:", err)
	}
}
