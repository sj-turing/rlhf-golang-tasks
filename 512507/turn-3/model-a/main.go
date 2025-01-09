package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"512507/turn-3/model-a/userstore" // Replace with your package path
)

var store *userstore.UserStore

func init() {
	filename := "users.gob"
	var err error
	store, err = userstore.NewUserStore(filename)
	if err != nil {
		log.Fatalf("Error initializing user store: %v", err)
	}

	// Register a signal handler to gracefully shut down the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("Received interrupt signal, shutting down gracefully...")
		if err := store.save(); err != nil {
			log.Printf("Failed to save data before shutdown: %v", err)
		}
		os.Exit(0)
	}()
}

func createUser(w http.ResponseWriter, r *http.Request) {
	type UserReq struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	var req UserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := store.Create(req.Email, req.Name, req.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func readUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email parameter required", http.StatusBadRequest)
		return
	}

	user, err := store.Read(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	type UserReq struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	var req UserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := store.Update(req.Email, req.Name, req.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email parameter required", http.StatusBadRequest)
		return
	}

	err := store.Delete(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func main() {
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/read", readUser)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)

	log.Println("Starting HTTP server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
