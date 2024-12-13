package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Item struct {
	Name string `json:"name"`
}

var (
	items []Item     // Shared data structure
	mu    sync.Mutex // Mutex for protecting shared data
)

// Endpoint to add an item
func addItem(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var item Item
	// Decode the incoming JSON data
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Lock the mutex for writing
	mu.Lock()
	// Append the new item to the shared list
	items = append(items, item)
	mu.Unlock() // Unlock the mutex

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Endpoint to get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex for reading
	mu.Lock()
	// Create a copy of the items slice to safely return data
	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)
	mu.Unlock() // Unlock the mutex

	// Respond with the list of items
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemsCopy)
}

func main() {
	http.HandleFunc("/items", addItem)
	http.HandleFunc("/items/all", getItems)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
