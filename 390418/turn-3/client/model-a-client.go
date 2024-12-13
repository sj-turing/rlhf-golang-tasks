package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	address := "localhost:8080"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type messages to send to the server. Type 'exit' to quit.")

	for {
		// Read a line from standard input
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			break
		}

		// Send the message to the server
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
