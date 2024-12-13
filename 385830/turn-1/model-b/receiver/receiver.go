package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type FileData struct {
	Filename string `json:"filename"`
	Content  []byte `json:"content"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: receiver <listen_ip> <listen_port>")
		os.Exit(1)
	}

	listenIP := os.Args[1]
	listenPort := os.Args[2]

	// Listen for incoming connections
	listener, err := net.Listen("tcp", listenIP+":"+listenPort)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Receive the JSON data length first
	length, err := receiveInt(conn)
	if err != nil {
		log.Printf("Error receiving data length: %v", err)
		return
	}

	// Receive the JSON data
	jsonData := make([]byte, length)
	_, err = conn.Read(jsonData)
	if err != nil {
		log.Printf("Error receiving JSON data: %v", err)
		return
	}

	// Decode the JSON data back to FileData
	var fileData FileData
	if err := json.Unmarshal(jsonData, &fileData); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	// Save the received file
	outputPath := fileData.Filename
	if err := ioutil.WriteFile(outputPath, fileData.Content, 0644); err != nil {
		log.Printf("Error writing file: %v", err)
		return
	}

	fmt.Println("File received successfully:", outputPath)
}

func receiveInt(conn net.Conn) (int, error) {
	buf := make([]byte, 4)
	_, err := conn.Read(buf)
	if err != nil {
		return 0, err
	}

	value := 0
	for i := 0; i < 4; i++ {
		value |= int(buf[i]) << (24 - 8*i)
	}

	return value, nil
}
