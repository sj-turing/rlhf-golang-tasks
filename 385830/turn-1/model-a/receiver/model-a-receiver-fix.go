package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type File struct {
	Name    string    // file name
	Size    int       // file size
	ModTime time.Time // modifiton time
	Data    []byte    // actual file data
}

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

	_, err = outputFile.WriteString(fmt.Sprintf("%s", receivedFile.Data))
	if err != nil {
		fmt.Println("Error writing file data:", err)
		return
	}

	fmt.Println("File saved successfully!")
}

func main() {
	receiveAndSaveFile("127.0.0.1") // Replace with the sender's IP address
}

func receiveFile(conn net.Conn) *File {
	var receivedFile File
	// Decode the file information (metadata)
	if err := gob.NewDecoder(conn).Decode(&receivedFile); err != nil {
		fmt.Println("Error receiving file metadata:", err)
		return nil
	}

	fmt.Println("Received file metadata:")
	fmt.Printf("Name: %s\nSize: %d\nModified Time: %s\n",
		receivedFile.Name, receivedFile.Size, receivedFile.ModTime)

	// Read the file content
	temp := make([]byte, receivedFile.Size)
	n, err := conn.Read(temp)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error receiving file data:", err)
			return nil
		}

	}

	receivedFile.Data = append(receivedFile.Data, temp[:n]...)
	fmt.Printf("Received %d bytes of file data\n", n)
	return &receivedFile
}
