package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name string `json:"name"`
}

type Items []Item

// Endpoint to add an item
func (i *Items) addItem(w http.ResponseWriter, r *http.Request) {
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

	*i = append(*i, item)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Endpoint to get all items
func (i *Items) getItems(w http.ResponseWriter, r *http.Request) {
	// Create a copy of the items slice to safely return data
	// Respond with the list of items
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}

func main() {
	// Start the write handler goroutine
	//	go writeHandler()

	items := Items{}

	http.HandleFunc("/items", items.addItem)
	http.HandleFunc("/items/all", items.getItems)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
