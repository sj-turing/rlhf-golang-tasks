package main

/*

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User visit storage
type VisitorStore struct {
	visitors map[string]struct{} // map to store visited users
}

// Channel request struct
type VisitorRequest struct {
	User     string
	Check    bool      // true if we are checking, false if we are visiting
	Response chan bool // channel to send the result back
}

// VisitorStore goroutine to handle concurrent access
func (store *VisitorStore) ManageVisitors(requests chan VisitorRequest) {
	for req := range requests {
		if req.Check {
			// Return true if the user has visited
			_, exists := store.visitors[req.User]
			req.Response <- exists
		} else {
			// Mark the user as visited
			store.visitors[req.User] = struct{}{} // Use empty struct for storage
			req.Response <- true
		}
	}
}

// Handler for visiting a user
func visitHandler(visitors chan VisitorRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user") // Get user from query param
		if user == "" {
			http.Error(w, "User missing", http.StatusBadRequest)
			return
		}

		responseChan := make(chan bool)
		request := VisitorRequest{User: user, Check: false, Response: responseChan}
		visitors <- request // Send request to visit

		// Wait for response
		<-responseChan
		fmt.Fprintf(w, "User %s has been visited.\n", user)
	}
}

// Handler for checking if user has visited
func checkHandler(visitors chan VisitorRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user") // Get user from query param
		if user == "" {
			http.Error(w, "User missing", http.StatusBadRequest)
			return
		}

		responseChan := make(chan bool)
		request := VisitorRequest{User: user, Check: true, Response: responseChan}
		visitors <- request // Send request to check

		// Wait for response
		visited := <-responseChan
		if visited {
			response := map[string]string{"status": "visited"}
			jsonResponse(w, response)
		} else {
			response := map[string]string{"status": "not visited"}
			jsonResponse(w, response)
		}
	}
}

// Helper function to return JSON response
func jsonResponse(w http.ResponseWriter, response map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Create the shared storage structure
	visitorStore := VisitorStore{
		visitors: make(map[string]struct{}),
	}

	// Create a channel for visitor requests
	visitorRequests := make(chan VisitorRequest)

	// Start the visitor management goroutine
	go visitorStore.ManageVisitors(visitorRequests)

	// Set up the HTTP handlers
	http.HandleFunc("/visit", visitHandler(visitorRequests))
	http.HandleFunc("/check", checkHandler(visitorRequests))

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

*/
