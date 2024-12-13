package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Struct to hold the user data and synchronization primitives
type UserTracker struct {
	users   map[string]bool
	rwmutex sync.RWMutex
}

// Initialize the user tracker
func NewUserTracker() *UserTracker {
	return &UserTracker{users: make(map[string]bool)}
}

// Method to mark a user as visited
func (ut *UserTracker) MarkVisited(username string) {
	ut.rwmutex.Lock()
	defer ut.rwmutex.Unlock()
	ut.users[username] = true
}

// Method to check if a user has been visited
func (ut *UserTracker) IsVisited(username string) string {
	ut.rwmutex.RLock()
	defer ut.rwmutex.RUnlock()
	if _, visited := ut.users[username]; visited {
		return "visited"
	}
	return "not visited"
}

// HTTP handler to mark a user as visited
func markVisitedHandler(w http.ResponseWriter, r *http.Request, ut *UserTracker) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username parameter is required", http.StatusBadRequest)
		return
	}

	ut.MarkVisited(username)
	fmt.Fprintf(w, "User '%s' marked as visited.\n", username)
}

// HTTP handler to check if a user has been visited
func checkVisitedHandler(w http.ResponseWriter, r *http.Request, ut *UserTracker) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username parameter is required", http.StatusBadRequest)
		return
	}

	status := ut.IsVisited(username)
	fmt.Fprintf(w, "User '%s' is %s.\n", username, status)
}

func main() {
	// Create an instance of UserTracker
	ut := NewUserTracker()

	// Set up HTTP handlers
	http.HandleFunc("/mark-visited", func(w http.ResponseWriter, r *http.Request) {
		markVisitedHandler(w, r, ut)
	})
	http.HandleFunc("/check-visited", func(w http.ResponseWriter, r *http.Request) {
		checkVisitedHandler(w, r, ut)
	})

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
