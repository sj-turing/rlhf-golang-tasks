package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"


	"512507/turn-3/model-b/store"

	"github.com/gorilla/mux"
)

var (
	userStore *store.UserStore
)

func main() {
	filename := "users.gob"
	store, err := store.NewUserStore(filename)
	if err != nil {
		log.Fatalf("Error creating user store: %v", err)
	}
	userStore = store

	// Set up HTTP router
	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{email}", readUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{email}", updateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/users/{email}", deleteUserHandler).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nGracefully shutting down...")
		if err := srv.Shutdown(nil); err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
		fmt.Println("Server shut down gracefully.")

		// Try to save the data on shutdown
		if err := userStore.save(); err != nil {
			log.Fatalf("Error saving user data on shutdown: %v", err)
		}
		os.Exit(0)
	}()

	fmt.Println("Server running on :8080")
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
