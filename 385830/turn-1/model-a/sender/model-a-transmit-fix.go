package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const port = "12345"

// File struct includes metadata and file content
type File struct {
	Name    string    // file name
	Size    int       // file size
	ModTime time.Time // modifiton time
	Data    []byte    // actual file data
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
