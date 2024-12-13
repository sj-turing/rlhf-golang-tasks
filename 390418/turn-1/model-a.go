package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Define the address of the TCP server
	address := "localhost:8080"

	// Establish a TCP connection
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Set custom read and write buffer sizes
	readBufferSize := 8192  // 8 KB
	writeBufferSize := 8192 // 8 KB

	err = conn.(*net.TCPConn).SetReadBuffer(readBufferSize)
	if err != nil {
		fmt.Println("Error setting read buffer:", err)
		return
	}

	err = conn.(*net.TCPConn).SetWriteBuffer(writeBufferSize)
	if err != nil {
		fmt.Println("Error setting write buffer:", err)
		return
	}

	fmt.Printf("Read buffer size set to: %d bytes\n", readBufferSize)
	fmt.Printf("Write buffer size set to: %d bytes\n", writeBufferSize)

	// Write some data to the server
	message := "Hello, TCP Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	// Read the response from the server
	buffer := make([]byte, readBufferSize)
	conn.SetDeadline(time.Now().Add(5 * time.Second)) // Set a deadline for read
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	fmt.Printf("Received from server: %s\n", string(buffer[:n]))
}
