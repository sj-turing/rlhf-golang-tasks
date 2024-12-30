package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// RequestHandler is the type for an HTTP request handler function.
type RequestHandler func(http.ResponseWriter, *http.Request)

// LoggingMiddleware adds logging to an HTTP request.
func LoggingMiddleware(next RequestHandler) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling request: %s %s\n", r.Method, r.URL)
		next(w, r)
	}
}

// AuthenticationMiddleware performs basic authentication on an HTTP request.
func AuthenticationMiddleware(next RequestHandler) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "Basic "+"YWRtaW46cGFzc3dvcmQ=" { // Basic auth for "admin:password"
			log.Println("Authentication successful")
			next(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
}

// RateLimitingMiddleware limits the number of requests per minute.
func RateLimitingMiddleware(next RequestHandler, limit int) RequestHandler {
	requestsPerMinute := make(chan struct{}, limit)
	return func(w http.ResponseWriter, r *http.Request) {
		requestsPerMinute <- struct{}{}
		defer func() {
			<-requestsPerMinute
		}()
		log.Println("Rate limiting request")
		next(w, r)
	}
}

// SimpleHandler is a sample HTTP request handler.
func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	// Create a chain of middleware
	handler := SimpleHandler
	handler = RateLimitingMiddleware(handler, 5) // Limit to 5 requests per minute
	handler = AuthenticationMiddleware(handler)
	handler = LoggingMiddleware(handler)

	// Serve the middleware chain
	http.HandleFunc("/", handler)
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
