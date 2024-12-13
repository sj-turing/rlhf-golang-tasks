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

	// Set a read timeout for the connection to avoid hanging indefinitely
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
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

			// Acknowledge the received message
			message := string(buf[:n])
			log.Printf("Received message: %s", message)
			response := fmt.Sprintf("Acknowledged: %s", message)
			_, err = conn.Write([]byte(response))
			if err != nil {
				log.Printf("Error writing to client: %v", err)
				return
			}

			// Reset read deadline after a successful read
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
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

			go handleConnection(ctx, conn)
		}
	}
}
