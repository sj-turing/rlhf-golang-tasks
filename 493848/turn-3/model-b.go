package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Middleware represents a callback function that takes an http.Handler and returns an http.Handler.
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware is a middleware that logs the request information.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method:", r.Method, "Path:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware is a middleware that performs basic authentication.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "user" || password != "pass" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ContentTypeMiddleware is a middleware that sets the Content-Type header.
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// MainHandler is the actual request handler.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", strings.TrimPrefix(r.URL.Path, "/"))
}

// composeMiddleware takes a list of middlewares and returns a composed middleware.
func composeMiddleware(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}

func main() {
	// Chain together the middlewares
	middlewareChain := composeMiddleware(LoggingMiddleware, AuthMiddleware, ContentTypeMiddleware)

	// Use the middleware chain with the main handler
	http.Handle("/", middlewareChain(http.HandlerFunc(MainHandler)))

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
