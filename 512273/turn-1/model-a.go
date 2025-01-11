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
)

const dbUrl = "postgres://user:password@localhost:5432/mydb?sslmode=disable"

// User struct represents a user's data
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Profile  string `json:"profile_image"`
}

func insertUser(db *sql.DB, user User) error {
	stmt, err := db.Prepare("INSERT INTO users (username, email, profile_image) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Profile)
	return err
}

func readFileAndInsertUsers(filePath string, db *sql.DB) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	users := []User{}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return err
	}

	for _, user := range users {
		if err := insertUser(db, user); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure the users table exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			profile_image VARCHAR(255)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	// Example JSON file paths
	filePaths := []string{"file1.json", "file2.json"}

	for _, filePath := range filePaths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			if err := readFileAndInsertUsers(path, db); err != nil {
				log.Errorf("Failed to process file %s: %v", path, err)
			} else {
				fmt.Println("File processed successfully:", path)
			}
		}(filePath)
	}

	wg.Wait()
}

