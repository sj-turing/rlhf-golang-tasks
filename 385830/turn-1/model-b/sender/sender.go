package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

// FileData struct to hold file content and metadata
type FileData struct {
	Filename string `json:"filename"`
	Content  []byte `json:"content"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: sender <server_ip> <server_port>")
		os.Exit(1)
	}

	serverIP := os.Args[1]
	serverPort := os.Args[2]

	filePath := "example.txt" // Replace this with the file you want to send
	fileData, err := readFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Connect to the server
	conn, err := net.Dial("tcp", serverIP+":"+serverPort)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send file data to the server
	if err := sendFileData(conn, fileData); err != nil {
		log.Fatalf("Error sending file data: %v", err)
	}

	fmt.Println("File sent successfully.")
}

func readFile(filePath string) (*FileData, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileData := &FileData{
		Filename: filePath,
		Content:  content,
	}

	return fileData, nil
}

func sendFileData(conn net.Conn, fileData *FileData) error {
	// Encode the file data to JSON
	jsonData, err := json.Marshal(fileData)
	if err != nil {
		return err
	}

	// Send the JSON data length first
	length := len(jsonData)
	if err := sendInt(conn, length); err != nil {
		return err
	}

	// Send the JSON data
	_, err = conn.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func sendInt(conn net.Conn, value int) error {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		buf[i] = byte(value >> (24 - 8*i))
	}
	_, err := conn.Write(buf)
	return err
}
