package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// User represents a user structure
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	PhotoURL string `json:"photo_url"`
}

func main() {
	// Open a database connection pool
	db, err := sql.Open("postgres", "postgres://your_user:your_password@localhost/your_database?sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close() // Close the database connection pool on program exit

	// Create a sync.WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Get the list of JSON files in a directory
	files, err := ioutil.ReadDir("./json_files")
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Iterate through each file and process it concurrently
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			wg.Add(1) // Increment the waitgroup for each file
			go processFile(db, &wg, file.Name())
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All files processed successfully.")
}

// processFile reads a JSON file, decodes the users, and stores them in the database
func processFile(db *sql.DB, wg *sync.WaitGroup, fileName string) {
	// Decrement the waitgroup when the function returns
	defer wg.Done()

	// Read the JSON file
	data, err := ioutil.ReadFile("./json_files/" + fileName)
	if err != nil {
		log.Printf("Error reading file %s: %v", fileName, err)
		return
	}

	// Decode the JSON data into a slice of User
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Printf("Error decoding JSON in file %s: %v", fileName, err)
		return
	}

	// Insert users into the database using prepared statements
	insertStmt, err := db.Prepare("INSERT INTO users (username, email, photo_url) VALUES ($1, $2, $3)")
	if err != nil {
		log.Printf("Error preparing insert statement: %v", err)
		return
	}
	defer insertStmt.Close() // Close the prepared statement on return

	// Execute insert statements for each user in the slice
	for _, user := range users {
		_, err := insertStmt.Exec(user.Username, user.Email, user.PhotoURL)
		if err != nil {
			log.Printf("Error inserting user %s: %v", user.Username, err)
			continue
		}
	}

	fmt.Println("File", fileName, "processed successfully.")
}
