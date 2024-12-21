package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// readFileHandler is a simple HTTP handler that reads a file and returns its contents
func readFileHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		defer file.Close() // defer to close the file

		// Read the file contents
		content, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error reading file: %v", err)
			http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
			return
		}

		// Write the file contents to the HTTP response
		w.Write(content)
	}
}

// main sets up and runs the HTTP server with the file reading handler
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Defer canceling the context to ensure proper resource cleanup

	http.HandleFunc("/readfile", readFileHandler(ctx))

	// Start the server
	fmt.Println("Server is running at http://localhost:8080/readfile")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
