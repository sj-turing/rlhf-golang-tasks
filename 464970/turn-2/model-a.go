package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func queryDatabase(db *sql.DB) error {
	defer db.Close()

	query := "SELECT * FROM users"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var userID int
	var username string
	for rows.Next() {
		err := rows.Scan(&userID, &username)
		if err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}
		fmt.Printf("User ID: %d, Username: %s\n", userID, username)
	}

	return nil
}

func main() {
	// Create a database connection
	dsn := "user:pass@tcp(127.0.0.1:3306)/my_database?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
		return
	}
	defer db.Close()

	// Create an HTTP client
	httpClient := &http.Client{}
	defer httpClient.CloseIdleConnections()

	// Define a handler for incoming HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")

		// Simulate a database query
		if err := queryDatabase(db); err != nil {
			http.Error(w, fmt.Sprintf("error querying database: %v", err), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("error starting server: %v\n", err)
		return
	}
}

// Graceful shutdown
func gracefulShutdown(server *http.Server, stop chan struct{}) {
	// Wait for a shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Received shutdown signal")

	// Stop accepting new connections
	close(stop)

	// Give the server time to gracefully handle active connections
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("error shutting down server: %v\n", err)
	}
}
