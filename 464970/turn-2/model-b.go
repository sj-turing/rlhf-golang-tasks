package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB     *sql.DB
	Client *http.Client
}

func (a *App) HandleUsers(w http.ResponseWriter, r *http.Request) {
	// Query the database using a context
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Sample database query
	query := "SELECT * FROM users"
	rows, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process the results and send response
	// ...
	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"users": []}`)
}

func (a *App) Run(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: a,
	}

	// Start the HTTP server
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for signals to initiate graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Shutting down server...")

	// Create a context with a timeout for closing connections
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the HTTP server
	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server shut down")
}

func main() {
	// Open database connection
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/my_database")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Create a new App with database connection and HTTP client
	a := &App{DB: db, Client: &http.Client{Timeout: 5 * time.Second}}

	var wg sync.WaitGroup
