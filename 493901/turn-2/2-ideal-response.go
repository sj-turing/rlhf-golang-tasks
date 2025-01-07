package main

import (
	"fmt"
	"net/http"
)

// HelloHandler responds to the "/hello" route with a message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// main serves as the entry point of the application.
func main() {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Starting application on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
