package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Retrieve environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("Using default port:", port)
	} else {
		fmt.Println("Using port from environment:", port)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// Define a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the API!")
	})

	// Start the server
	log.Println("Server starting on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
