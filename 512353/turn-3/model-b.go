package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"
)

type User struct {
	Username    string  `json:"username"`
	SpendAmount float64 `json:"spendAmount"`
	Email       string  `json:"email"`
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func readUsersFromJSON(ctx context.Context, filePath string) ([]User, error) {
	// Don't proceed if the context is canceled
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

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	// Create a child context with the time limit of 2 seconds for file reading
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	filePath := "users.json"
	users, err := readUsersFromJSON(ctx, filePath)
	if err != nil {
		// Handle intermediate errors
		switch err {
		case context.DeadlineExceeded:
			http.Error(w, "Service unavailable due to timeout", http.StatusServiceUnavailable)
		default:
			http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func handleDisposeRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Shutting down the service...")
	// Signal the main goroutine to shut down
	shutdownChan <- os.Interrupt
	http.Error(w, "Service shutting down...", http.StatusServiceUnavailable)
}

var (
	shutdownChan = make(chan os.Signal, 1)
)

func main() {
	// Handle graceful shutdown
	go func() {
		<-shutdownChan
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/data", handleDataRequest)
	mux.HandleFunc("/dispose", handleDisposeRequest)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Service started on port 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Wait for shutdown signal
	<-shutdownChan

	log.Println("Graceful shutdown in progress...")
	// Create a context with 5 seconds for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shut down the server: %v", err)
	}

	log.Println("Server stopped.")
}
