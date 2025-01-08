package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// User represents the structure of a user
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response represents the response structure for API calls
type Response struct {
	Message string `json:"message"`
}

var db *sql.DB

func init() {
	// Initialize the database connection
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// Set the maximum number of open connections in the pool
	db.SetMaxOpenConns(100)
	// Set the maximum number of idle connections in the pool
	db.SetMaxIdleConns(20)
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := Response{Message: message}
	json.NewEncoder(w).Encode(res)
}

func userSignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	// Decode the JSON request body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	// Validate user input
	if user.Name == "" || user.Email == "" {
		jsonResponse(w, http.StatusBadRequest, "Invalid user information")
		return
	}

	// Insert the user into the database using a prepared statement
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		jsonResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		jsonResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		jsonResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	log.Printf("User signed up successfully: ID %d, Name: %s, Email: %s", lastInsertID, user.Name, user.Email)

	jsonResponse(w, http.StatusCreated, "User signed up successfully")
}

func main() {
	http.HandleFunc("/signup", userSignupHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
