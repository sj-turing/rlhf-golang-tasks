package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	delimiter      = '\n'
	maxMessageSize = 65536 // 64 KB limit per message
)

func main() {
	// Connecting to the TCP server
	address := "localhost:8080"

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	go sendMessages(writer)
	receiveMessages(reader)
}

func sendMessages(writer *bufio.Writer) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if len(message) > maxMessageSize {
			fmt.Println("Message exceeds size limit.")
			continue
		}

		_, err := writer.WriteString(message + "\n")
		if err != nil {
			log.Println("Error writing message:", err)
		}
		err = writer.Flush()
		if err != nil {
			log.Println("Error flushing writer:", err)
		}
	}
}

func receiveMessages(reader *bufio.Reader) {
	for {
		message, err := readMessage(reader)
		if err != nil {
			if err == io.EOF {
				log.Println("Server disconnected")
				return
			}
			log.Println("Error reading message:", err)
			return
		}

		fmt.Printf("Received: %s", message)
	}
}

func readMessage(reader *bufio.Reader) (string, error) {
	message := []byte{}

	for {
		b, isPrefix, err := reader.ReadLine()
		if err != nil {
			return "", err
		}
		if isPrefix {
			return "", fmt.Errorf("message exceeds %d bytes", maxMessageSize)
		}

		message = append(message, b...)

		// Check if we've reached the end of a message
		if len(message) > 0 && message[len(message)-1] == delimiter {
			break
		}
	}

	return string(message[:len(message)-1]), nil // Remove the delimiter
}
