package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// handleConnection manages a single TCP connection from a client.
func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from:", conn.RemoteAddr())

	// Set a read deadline for the connection to prevent hanging
	readDeadline := time.Now().Add(10 * time.Second)
	if err := conn.SetReadDeadline(readDeadline); err != nil {
		log.Printf("Failed to set read deadline: %v", err)
		return
	}

	buf := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			log.Println("Connection cancelled:", conn.RemoteAddr())
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					log.Println("Client disconnected:", conn.RemoteAddr())
					return
				}
				log.Printf("Error reading from client: %v", err)
				return
			}

			// Process the received message
			message := string(buf[:n])
			log.Printf("Received message: %s", message)

			// Send acknowledgment back to the client
			response := fmt.Sprintf("Acknowledged: %s", message)
			if _, err := conn.Write([]byte(response)); err != nil {
				log.Printf("Error writing to client: %v", err)
				return
			}

			// Reset the read deadline after successful read
			readDeadline = time.Now().Add(10 * time.Second)
			if err := conn.SetReadDeadline(readDeadline); err != nil {
				log.Printf("Failed to reset read deadline: %v", err)
				return
			}
		}
	}
}

// main function starts the TCP server.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddr := ":8080"
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", serverAddr, err)
	}
	defer listener.Close()

	log.Println("TCP Server started and listening on", serverAddr)

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down server...")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v", err)
				continue
			}

			log.Println("Accepted connection from:", conn.RemoteAddr())
			go handleConnection(ctx, conn)
		}
	}
}
