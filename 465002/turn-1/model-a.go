package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// readFileHandler is a simple HTTP handler that reads a file and returns its contents
func readFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filepath")
	if filePath == "" {
		http.Error(w, "No file path specified", http.StatusBadRequest)
		return
	}

	// Defer the closing of the file to ensure it's closed regardless of how the function exits
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		http.Error(w, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file contents
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the file contents to the HTTP response
	w.Write(content)
}

// main sets up and runs the HTTP server with the file reading handler
func main() {
	ctx := context.Background()
	defer context.Cancel(ctx) // Defer canceling the context to ensure proper resource cleanup

	http.HandleFunc("/readfile", readFileHandler)

	// Start the server
	fmt.Println("Server is running at http://localhost:8080/readfile")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
