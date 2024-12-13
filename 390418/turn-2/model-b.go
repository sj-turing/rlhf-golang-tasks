package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

const (
	batchSize = 8192 // 8 KB
)

func main() {
	// Define the address of the TCP server
	address := "localhost:8080"

	// Establish a TCP connection
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create a buffer reader for efficient reading
	reader := bufio.NewReader(conn)

	// Read the total file size sent by the server
	var fileSize uint64
	err = binary.Read(reader, binary.LittleEndian, &fileSize)
	if err != nil {
		fmt.Println("Error reading file size:", err)
		return
	}

	fmt.Printf("Receiving file of size: %d bytes\n", fileSize)

	// Open a file to write the received data
	outputFile, err := os.Create("received_file.bin")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Receive data in batches
	bytesReceived := 0
	for bytesReceived < int64(fileSize) {
		batch := make([]byte, batchSize)
		n, err := io.ReadFull(conn, batch)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from server:", err)
			return
		}

		bytesReceived += int64(n)
		fmt.Printf("Received %d/%d bytes\r", bytesReceived, fileSize)

		_, err = outputFile.Write(batch[:n])
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("\nFile received successfully!")
}
