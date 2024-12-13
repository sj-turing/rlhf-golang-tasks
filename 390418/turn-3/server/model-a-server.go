package main

/*

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Start TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		go processMessage(message) // Process messages concurrently
	}
}

func processMessage(message string) {
	// Simulate processing time with a simple print statement
	fmt.Printf("Received message: %s", message)
	// Add more processing logic as needed
}

*/
