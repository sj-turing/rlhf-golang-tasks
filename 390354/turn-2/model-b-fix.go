package main

import (
	"context"
	"log"
	"net"
	"time"
)

// echoHandler handles a single TCP connection by echoing back received messages.
func echoHandler(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from", conn.RemoteAddr())

	for {
		select {
		case <-ctx.Done():
			log.Println("Connection cancelled")
			return
		default:
			// Set a read deadline
			if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
				log.Printf("failed to set read deadline: %v", err)
				return
			}

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}

			log.Printf("Received: %s", string(buf[:n]))

			// Set a write deadline
			if err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second)); err != nil {
				log.Printf("failed to set write deadline: %v", err)
				return
			}

			_, err = conn.Write([]byte("Echo: " + string(buf[:n])))
			if err != nil {
				log.Printf("Write error: %v", err)
				return
			}
		}
	}
}

// main starts a TCP server on port 8080.
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

			// Create a new context for this connection with a cancellation function
			connCtx, connCancel := context.WithCancel(ctx)
			defer connCancel()

			go echoHandler(connCtx, conn)
		}
	}
}
