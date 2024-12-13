package main

import (
	"fmt"
	"io"
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

	// Open the file to send
	file, err := os.Open("largefile.dat") // Change to your file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the file
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
	}
	fmt.Println("File sent successfully")
}
