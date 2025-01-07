package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)

	res := Response{Message: message}

	json.NewEncoder(w).Encode(res)
}

func main() {

	http.HandleFunc("/defer-cleanup", deferCleanupHandler)
	http.HandleFunc("/manual-cleanup", manualCleanupHandler)

	log.Println("Starting application on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func deferCleanupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method not supported")
		return
	}

	var user User
	defer func() {
		log.Println("Closing the request body")
		r.Body.Close()
	}()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	log.Printf("incoming request for defer cleanup: %+v\n", user)
	jsonResponse(w, http.StatusCreated, "User created successfully")
}

func manualCleanupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method not supported")
		return
	}

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// reason to close body after the decoding is
	// once body get closed, json decoder will not able to read it
	r.Body.Close()

	log.Printf("incoming request for manual cleanup: %+v\n", user)
	jsonResponse(w, http.StatusCreated, "User created successfully")
}
