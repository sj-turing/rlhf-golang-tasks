package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// maxClients defines the maximum number of concurrent client connections
const maxClients = 2

// handleConnection manages a single TCP connection from a client.
func handleConnection(ctx context.Context, conn net.Conn, clientID int) {
	defer conn.Close()
	log.Printf("New connection from client %d: %s\n", clientID, conn.RemoteAddr())

	// Set a read timeout for the connection to avoid hanging indefinitely
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	buf := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Connection cancelled: client %d\n", clientID)
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					log.Printf("Client %d disconnected\n", clientID)
					return
				}
				log.Printf("Error reading from client %d: %v\n", clientID, err)
				return
			}

			// Acknowledge the received message
			message := string(buf[:n])
			log.Printf("Received message from client %d: %s\n", clientID, message)
			response := fmt.Sprintf("Acknowledged: %s\n", message)
			_, err = conn.Write([]byte(response))
			if err != nil {
				log.Printf("Error writing to client %d: %v\n", clientID, err)
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

	log.Printf("TCP Server started and listening on %s\n", serverAddr)

	var clientCount int

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down server...")
			return
		default:
			if clientCount >= maxClients {
				log.Println("Server at maximum client capacity. Refusing new connections.")
				continue
			}

			clientCount++

			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v\n", err)
				continue
			}

			go handleConnection(ctx, conn, clientCount)
		}
	}
}
