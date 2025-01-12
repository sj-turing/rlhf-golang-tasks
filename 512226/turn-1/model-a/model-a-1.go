package main

import (
	"fmt"
	"net/http"
	"url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// Example of extracting parameters
	userID := params.Get("user_id")
	page := params.Get("page")

	// Log or store the parameters
	fmt.Printf("UserID: %s, Page: %s\n", userID, page)
}

func main() {
	http.HandleFunc("/track", handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
