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
	items   []Item    // Shared data structure
	writeCh chan Item // Channel for synchronizing writes
)

func init() {
	// Create a buffer of size 1 for the write channel
	writeCh = make(chan Item, 1)
}

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

	// Send the item to the write channel
	writeCh <- item

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Endpoint to get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	// Create a copy of the items slice to safely return data
	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)

	// Respond with the list of items
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemsCopy)
}

// Background goroutine to handle writes to the shared list
func writeHandler() {
	for item := range writeCh {
		// Append the new item to the shared list
		items = append(items, item)
	}
}

func main() {
	// Start the write handler goroutine
	go writeHandler()

	http.HandleFunc("/items", addItem)
	http.HandleFunc("/items/all", getItems)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
