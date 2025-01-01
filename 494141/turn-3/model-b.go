package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	// Use a sync.Map to store user ID to request count mapping
	userRequestCount sync.Map
)

func main() {
	http.HandleFunc("/", logUserRequest)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func logUserRequest(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request header
	userID := r.Header.Get("User-ID")
	if userID == "" {
		http.Error(w, "User-ID header is missing", http.StatusBadRequest)
		return
	}

	// Parse the user ID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid User-ID header", http.StatusBadRequest)
		return
	}

	// Update the user request count in the sync.Map
	userRequestCount.Store(userIDInt, incrementRequestCount(userIDInt))

	fmt.Fprintf(w, "User ID %d requested.\n", userIDInt)
}

func incrementRequestCount(userID int) int {
	// Load the current request count for the user
	if count, ok := userRequestCount.Load(userID); ok {
		return count.(int) + 1
	}
	return 1 // If the user ID doesn't exist, return 1 as the first request
}
