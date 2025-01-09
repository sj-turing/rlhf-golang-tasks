package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"512507/turn-3/ideal-response/store"

	"github.com/gorilla/mux"
)

var userStore *store.UserStore

type UserReq struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req UserReq

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := userStore.Create(req.Email, req.Name, req.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func readUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	if email == "" {
		http.Error(w, "Email parameter required", http.StatusBadRequest)
		return
	}

	user, err := userStore.Read(email)
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

	var req UserReq
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := userStore.Update(req.Email, req.Name, req.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	if email == "" {
		http.Error(w, "Email parameter required", http.StatusBadRequest)
		return
	}

	err := userStore.Delete(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func main() {

	filename := "users.gob"
	var err error
	userStore, err = store.NewUserStore(filename)
	if err != nil {
		log.Fatalf("Error initializing user store: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/user", createUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{email}", readUser).Methods(http.MethodGet)
	r.HandleFunc("/user", updateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{email}", deleteUser).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Register a signal handler to gracefully shut down the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	log.Println("Starting HTTP server on port 8080")

	go func(svr *http.Server) {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error while svr.ListenAndServe: %v", err)
		}
	}(srv)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Error while svr.Shutdown: %v", err)

		// in case of failure to graceful shutdown
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Executing force shutdown")
			// trying to forcefully shutdown
			if err := srv.Close(); err != nil {
				log.Printf("Server close failed: %v", err)
				log.Printf("Proceeding for server exit")
			} else {
				log.Printf("Server close completed")
			}
		}
	}

	log.Println("Received interrupt signal, shutting down gracefully...")
	if err := userStore.Save(); err != nil {
		log.Printf("Failed to save data before shutdown: %v", err)
	}
	os.Exit(0)
}
