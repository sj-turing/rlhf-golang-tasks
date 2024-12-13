package main

import (
	"bufio"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const chunkSize = 4096 // 4KB chunk size

// FileMetadata holds the metadata of a file
type FileMetadata struct {
	Name string
	Size int64
}

// sendFile sends a file along with its metadata over a TCP connection
func sendFile(conn net.Conn, filePath string) {
	defer conn.Close()

	// Read the file metadata
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}

	metadata := FileMetadata{
		Name: info.Name(),
		Size: info.Size(),
	}

	// Encode the metadata and send it
	enc := gob.NewEncoder(conn)
	if err := enc.Encode(metadata); err != nil {
		log.Fatalf("Error encoding metadata: %v", err)
	}

	// Send the file content in chunks
	for {
		buf := make([]byte, chunkSize)
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		// Write the chunk size
		if err := binary.Write(conn, binary.LittleEndian, int32(n)); err != nil {
			log.Fatalf("Error writing chunk size: %v", err)
		}

		// Write the chunk content
		if _, err := conn.Write(buf[:n]); err != nil {
			log.Fatalf("Error writing chunk: %v", err)
		}
	}

	log.Printf("File %s sent successfully.", filePath)
}

// receiveFile receives a file along with its metadata over a TCP connection
func receiveFile(conn net.Conn) {
	defer conn.Close()

	// Decode the metadata
	dec := gob.NewDecoder(conn)
	var metadata FileMetadata
	if err := dec.Decode(&metadata); err != nil {
		log.Fatalf("Error decoding metadata: %v", err)
	}

	// Create the output file
	outputFile, err := os.Create(metadata.Name)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Receive the file content in chunks
	for {
		var chunkSize int32
		if err := binary.Read(conn, binary.LittleEndian, &chunkSize); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading chunk size: %v", err)
		}

		buf := make([]byte, chunkSize)
		if _, err := io.ReadFull(conn, buf); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading chunk: %v", err)
		}

		if _, err := outputFile.Write(buf); err != nil {
			log.Fatalf("Error writing chunk to file: %v", err)
		}
	}

	log.Printf("File %s received successfully.", metadata.Name)
