package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type User struct {
	Username    string  `json:"username"`
	SpendAmount float64 `json:"spendAmount"`
	Email       string  `json:"email"`
}

var serverCtx context.Context
var serverCancel context.CancelFunc
var shutdownWG sync.WaitGroup

func readUsersFromFile(ctx context.Context, filePath string) ([]User, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// Create a child context with a timeout derived from the request context
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "file query parameter is required", http.StatusBadRequest)
		return
	}

	users, err := readUsersFromFile(ctx, filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusServiceUnavailable)
		log.Printf("Error reading file %s: %v", filePath, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func disposeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shutting down gracefully...\n"))
	serverCancel()
}

func main() {
	serverCtx, serverCancel = context.WithCancel(context.Background())

	http.HandleFunc("/readfile", fileHandler)
	http.HandleFunc("/dispose", disposeHandler)

	server := &http.Server{Addr: ":8080"}

	// Run server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Listen for shutdown signal
	<-serverCtx.Done()
	log.Println("Shutdown signal received")

	// Gracefully shutdown the server with a timeout context
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
	log.Println("Server exited properly")
}
