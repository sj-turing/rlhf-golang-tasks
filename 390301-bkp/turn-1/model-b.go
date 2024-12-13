package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Shared data structure
type Counter struct {
	Value int        `json:"value"`
	mux   sync.Mutex `json:"-"` // Mutex to protect concurrent access
}

// Global shared counter
var sharedCounter = Counter{Value: 0}

// Endpoint to increment the counter
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	sharedCounter.mux.Lock()
	defer sharedCounter.mux.Unlock()

	sharedCounter.Value++

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sharedCounter); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Endpoint to get the current counter value
func getCounterHandler(w http.ResponseWriter, r *http.Request) {
	sharedCounter.mux.RLock() // Read lock, allows multiple readers
	defer sharedCounter.mux.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sharedCounter); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/get", getCounterHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
