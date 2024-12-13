package main

/*

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

const port = "12345"

// File struct includes metadata and file content
type File struct {
	Name    string         // file name
	Size    int            // file size
	ModTime os.FileModTime // modifiton time
	Data    []byte         // actual file data
}

// sendFile implements the functionality to send the file structure using TCP.
func sendFile(conn net.Conn, file *File) {
	defer conn.Close()
	// Encode the file information (metadata)
	if err := gob.NewEncoder(conn).Encode(file); err != nil {
		fmt.Println("Error sending file metadata:", err)
		return
	}
	// Send the file content
	n, err := conn.Write(file.Data)
	if err != nil {
		fmt.Println("Error sending file data:", err)
		return
	}
	fmt.Printf("Sent %d bytes of file data\n", n)
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
	receivedFile.Data = make([]byte, receivedFile.Size)
	n, err := io.ReadFull(conn, receivedFile.Data)
	if err != nil {
		fmt.Println("Error receiving file data:", err)
		return nil
	}
	fmt.Printf("Received %d bytes of file data\n", n)
	return &receivedFile
}

func main() {
	// Die if failed to open localFile
	localFile, err := os.Open("example.txt") // Replace this with the actual file path you want to send
	if err != nil {
		fmt.Println("Error opening local file:", err)
		return
	}
	defer localFile.Close()

	fileStat, err := localFile.Stat()
	if err != nil {
		fmt.Println("Error stat'ing file:", err)
		return
	}
	// Read the file content into memory
	fileBytes, err := io.ReadAll(localFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Send the file
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening tcp:", err)
		return
	}
	defer socket.Close()

	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("Error accepting connections:", err)
			return
		}

		go func(conn net.Conn) {
			fileToSend := &File{
				Name:    fileStat.Name(),
				Size:    int(fileStat.Size()),
				ModTime: fileStat.ModTime(),
				Data:    fileBytes,
			}
			sendFile(conn, fileToSend)
		}(conn)
	}
}
*/
