package main

/*

import (
	"fmt"
	"io"
	"net"
	"os"
)

const port = "12345"

func receiveAndSaveFile(remoteAddress string) {
	conn, err := net.Dial("tcp", remoteAddress+":"+port)
	if err != nil {
		fmt.Println("Error dialing tcp:", err)
		return
	}
	defer conn.Close()
	receivedFile := receiveFile(conn)
	if receivedFile == nil {
		return // Return if there was an error receiving the file
	}

	// Create a new file to store the received data
	outputFile, err := os.Create(receivedFile.Name)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	_, err = io.WriteFull(outputFile, receivedFile.Data)
	if err != nil {
		fmt.Println("Error writing file data:", err)
		return
	}

	fmt.Println("File saved successfully!")
}

func main() {
	receiveAndSaveFile("127.0.0.1") // Replace with the sender's IP address
}
*/
