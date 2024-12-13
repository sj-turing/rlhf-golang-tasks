package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Connecting to the TCP server
	address := "localhost:8080"

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Create a file to save the received data
	outFile, err := os.Create("received_file.dat") // Change to your desired output file
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Buffer for reading data
	buffer := make([]byte, 8192) // 8 KB buffer size
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // End of file, exit loop
			}
			fmt.Println("Error reading data:", err)
			return
		}
		if n > 0 {
			_, err = outFile.Write(buffer[:n])
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}

	fmt.Println("File received successfully")
}
