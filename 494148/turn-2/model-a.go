package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Created string `json:"created_at"`
}

var dbConn *sql.DB

func init() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}

	var err error
	dbConn, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ensure the users table exists
	_, err = dbConn.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)

	res := Response{Message: message}

	json.NewEncoder(w).Encode(res)
}

func userSignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method not supported")
		return
	}

	var user User

	defer func() {
		r.Body.Close()
	}()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if _, err := dbConn.Exec("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email); err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Failed to sign up user")
		return
	}

	jsonResponse(w, http.StatusCreated, "User signed up successfully")
}

func main() {
	http.HandleFunc("/signup", userSignupHandler)

	log.Println("Starting application on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Response struct {
	Message string `json:"message"`
}
