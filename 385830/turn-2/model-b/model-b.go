package main

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	chunkSize  = 4096 // Size of each chunk in bytes
	maxRetries = 3    // Number of retries for lost chunks
)

// FileChunk contains metadata and content of a file chunk
type FileChunk struct {
	SeqNum   uint32 // Sequence number of the chunk
	IsLast   bool   // True if it's the last chunk
	Metadata []byte // File metadata (e.g., filename, size, etc.)
	Content  []byte // File content
	Checksum []byte // MD5 checksum of the content
}

func generateChecksum(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

// SendFile sends a file over TCP in chunks with lossless transmission
func SendFile(conn net.Conn, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	metadata := []byte{
		'F', 'I', 'L', 'E', // File header
		byte(info.Mode() & 0777), // File permissions
		uint8(info.Size() >> 24),
		uint8((info.Size() >> 16) & 0xFF),
		uint8((info.Size() >> 8) & 0xFF),
		uint8(info.Size() & 0xFF),
	}

	var seqNum uint32 = 1
	for {
		buf := make([]byte, chunkSize)
		n, err := file.Read(buf)
		if err == io.EOF {
			buf = buf[:n]
			isLast := true
			chunk := &FileChunk{
				SeqNum:   seqNum,
				IsLast:   isLast,
				Metadata: metadata,
				Content:  buf,
				Checksum: generateChecksum(buf),
			}
			err = sendChunk(conn, chunk)
			if err != nil {
				return err
			}
			fmt.Println("File sent successfully!")
			break
		} else if err != nil {
			return err
		}

		isLast := false
		chunk := &FileChunk{
			SeqNum:   seqNum,
			IsLast:   isLast,
			Metadata: metadata,
			Content:  buf,
			Checksum: generateChecksum(buf),
		}
		err = sendChunk(conn, chunk)
		if err != nil {
			return err
		}
		seqNum++
	}

	return nil
}

func sendChunk(conn net.Conn, chunk *FileChunk) error {
	// Encode the chunk data
	buf := make([]byte, 4+len(chunk.Metadata)+len(chunk.Content)+len(chunk.Checksum))
	binary.BigEndian.PutUint32(buf[:4], chunk.SeqNum)
	buf[4] = uint8(0) // Reserved byte (can be used for flags in future)
	copy(buf[5:], chunk.Metadata)
	copy(buf[5+len(chunk.Metadata):], chunk.Content)
	copy(buf[5+len(chunk.Metadata)+len(chunk.Content):], chunk.Checksum)

	for retries := 0; retries < maxRetries; retries++ {
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("Failed to send chunk", chunk.SeqNum, ", retrying...")
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			continue
		}
		fmt.Println("Sent chunk", chunk.SeqNum)
		return nil
	}
	return errors.New("failed to send chunk after maximum retries")
}

func main() {
	port := "8080"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Got a connection!")

		go func(conn net.Conn) {
			defer conn.Close()
			SendFile(conn, "example.txt")
			//	err := receiveFile(conn)
			if err != nil {
				fmt.Println("Error receiving file:", err)
			}
		}(conn)
	}
}
