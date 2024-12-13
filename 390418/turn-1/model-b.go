package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	// Define buffer sizes in bytes
	readBufferSize := 65536  // 64KB
	writeBufferSize := 65536 // 64KB

	// Connect to a TCP server
	addr := "example.com:8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Set socket buffer sizes
	setSocketBuffer(conn.Fd(), syscall.SO_RCVBUF, readBufferSize)
	setSocketBuffer(conn.Fd(), syscall.SO_SNDBUF, writeBufferSize)

	fmt.Println("Connected to", addr)
	fmt.Println("Read buffer size:", getSocketBuffer(conn.Fd(), syscall.SO_RCVBUF))
	fmt.Println("Write buffer size:", getSocketBuffer(conn.Fd(), syscall.SO_SNDBUF))

	// Example of reading and writing data
	_, err = conn.Write([]byte("Hello, Server!"))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Response:", string(buf[:n]))
}

// setSocketBuffer sets the socket buffer size
func setSocketBuffer(fd uintptr, opt int, size int) {
	if err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, opt, size); err != nil {
		fmt.Println("Error setting socket buffer:", err)
	}
}

// getSocketBuffer gets the socket buffer size
func getSocketBuffer(fd uintptr, opt int) int {
	var size int
	if err := syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, opt, &size); err != nil {
		fmt.Println("Error getting socket buffer:", err)
	}
	return size
}
