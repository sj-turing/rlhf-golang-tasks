package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	n, err := file.WriteString("Hello, World!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	if err := file.Sync(); err != nil {
		fmt.Println("Error synchronizing file:", err)
		return
	}
	fmt.Println("Wrote", n, "bytes to file successfully.")
}
