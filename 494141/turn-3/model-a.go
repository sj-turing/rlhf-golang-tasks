package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var requestCounts sync.Map

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Log the request
	log.Printf("Request from %s: %s", r.RemoteAddr, r.URL)

	// Extract user ID from request header or query parameter (assumed for simplicity)
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		// Respond with an error if user_id is missing
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Status:  "error",
			Message: "user_id parameter is missing",
		})
		return
	}

	// Increment request count for the user ID
	if existingCount, exists := requestCounts.LoadOrStore(userID, 0); exists {
		count := existingCount.(int) + 1
		requestCounts.Store(userID, count)
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		Status:  "success",
		Message: fmt.Sprintf("Request received from user ID: %s", userID),
	})
}

func main() {
	http.HandleFunc("/api/requests", handler)
	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
