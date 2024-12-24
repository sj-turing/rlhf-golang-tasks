package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB     *sql.DB
	Client *http.Client
}

type User struct {
	ID       int
	Username string
}

func (a *App) HandleUsers(w http.ResponseWriter, r *http.Request) {
	// Query the database using a context
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Sample database query
	query := "SELECT id, username FROM users"
	rows, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users = []User{}

	for rows.Next() {
		user := User{}
		// mapping the database data with struct
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			log.Printf("Error while scanning data from rows: %v", err)
			continue
		}
		users = append(users, user)
	}

	// Process the results and send response
	// ...
	// Send JSON response
	res, err := a.Client.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Printf("Error getting the data from API: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// using defer to close body
	defer res.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, users)
}

func Run(app *App, addr string) *http.Server {
	srv := &http.Server{
		Addr: addr,
	}

	http.HandleFunc("/", app.HandleUsers)

	// Start the HTTP server
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	return srv
}

func gracefulShutdown(srv *http.Server) {
	// Wait for signals to initiate graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Shutting down server...")

	// Create a context with a timeout for closing connections
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if srv != nil {
		// Shutdown the HTTP server
		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
	}

	log.Println("Server shut down")
}

type rabbitMQ struct{}

func NewRabbitMQConnection() (*rabbitMQ, error) {
	return &rabbitMQ{}, nil
}

func (rmq *rabbitMQ) Close() error {
	if rmq != nil {
		rmq = nil
	}
	log.Println("Closing the RabbitMQ open connection")
	return nil
}

func main() {

	// Open database connection
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/my_database")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer func() {
		log.Print("Closing the mySQL connection")
		db.Close()
	}()

	db.SetMaxIdleConns(2) // setting max Idle connection
	db.SetMaxOpenConns(5) // setting max open connections

	// mimic to create a rabbitMQ connection
	rmq, err := NewRabbitMQConnection()
	if err != nil {
		log.Fatalf("Error getting RMQ connection: %v", err)
	}
	defer rmq.Close()

	// Create a new App with database connection and HTTP client
	a := &App{DB: db, Client: &http.Client{Timeout: 5 * time.Second}}

	// running server on port 8080
	srv := Run(a, ":8080")

	// gracefulShutdow if receives any interrupt
	gracefulShutdown(srv)
}
