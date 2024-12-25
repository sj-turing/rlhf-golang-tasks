package main

import (
	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
	"os"

	"path/filepath"

	"464953/turn-1/model-a/dbpkg"
)

func main() {
	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	defer db.Close()

	// Create the query_params table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS query_params (
			id SERIAL PRIMARY KEY,
			params TEXT
		)
	`)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	// Define the URL with special characters
	urlString := "http://example.com/search?q=special+characters&key=value%20with%20spaces"

	// Parse and store the URL parameters
	err = dbpkg.ParseAndStoreQueryParams(db, urlString)
	if err != nil {
		log.Fatalf("error parsing and storing query parameters: %v", err)
	}

	log.Println("Query parameters successfully stored in the database.")
}
