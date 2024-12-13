package main

/*

import (
	"context"
	"log"
	"net"
	"time"
)

// handleConnection is a function to handle a single TCP connection.
func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from", conn.RemoteAddr())

	// Set a shorter read deadline to demonstrate connection timing out
	if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Fatalf("failed to set read deadline: %v", err)
	}

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		if ctx.Err() != nil {
			log.Println("Connection cancelled")
		} else if net.IsTimeout(err) {
			log.Println("Connection timed out")
		} else {
			log.Printf("Read error: %v", err)
		}
		return
	}

	log.Printf("Received data: %s", string(buf))

	// Simulate processing time
	time.Sleep(3 * time.Second)

	// Set a write deadline to ensure we don't block indefinitely
	if err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second)); err != nil {
		log.Fatalf("failed to set write deadline: %v", err)
	}

	_, err = conn.Write([]byte("Hello, client!"))
	if err != nil {
		if ctx.Err() != nil {
			log.Println("Write cancelled")
		} else if net.IsTimeout(err) {
			log.Println("Write timed out")
		} else {
			log.Printf("Write error: %v", err)
		}
		return
	}

	log.Println("Response sent")
}

// main starts a TCP server and handles incoming connections with context.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Starting TCP server on :8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer listener.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("Server shutting down")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				if ctx.Err() != nil {
					log.Println("Accept cancelled")
					return
				}
				log.Printf("Accept error: %v", err)
				continue
			}

			go handleConnection(ctx, conn)
		}
	}
}

*/
