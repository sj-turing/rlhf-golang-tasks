package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Port        string
	DatabaseURI string
}

var config Config

func init() {
	// Read environment variables
	config.Port = getEnv("PORT", "8080")            // Default to "8080"
	config.DatabaseURI = getEnv("DATABASE_URI", "") // Default to empty if not set
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "hello world"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Printf("Starting server on port %s\n", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
