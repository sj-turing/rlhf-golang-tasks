package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the content of "example.txt" file
	content, err := os.ReadFile("example.txt")
	if err != nil {
		// Error handling if there is an issue with reading the file
		fmt.Println("Error reading file:", err)
		return
	}

	// Successfully read the file, now process the content
	fmt.Println("File content:")
	fmt.Println(string(content))

	// Stat the file after reading to get some additional information
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		// Error handling for stat operation
		fmt.Println("Error getting file stats:", err)
		return
	}

	// Now you can access various file metadata like size, mode, modification time, etc.
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("File permissions:", fileInfo.Mode())
}
