// filename: main.go
package main

/*

import (
	"fmt"
	"log"
	"net/http"
)

// recoverMiddleware is a middleware that catches panics and logs them, returning a 500 error.
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				// Log the panic
				log.Printf("Recovered panic: %v\n", p)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

// divideByZero is a simple handler that panics when trying to divide by zero.
func divideByZero(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	if n == "" {
		http.Error(w, "No number specified", http.StatusBadRequest)
		return
	}
	num, err := fmt.Sscanf(n, "%d", &n)
	if err != nil || num != 1 {
		http.Error(w, "Invalid number specified", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "1 / %d = %d\n", n, 1/n)
}

// main sets up and runs the HTTP server with the middleware and the handler
func main() {
	http.HandleFunc("/divide", recoverMiddleware(http.HandlerFunc(divideByZero)))

	fmt.Println("Server is running at http://localhost:8080/divide")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
