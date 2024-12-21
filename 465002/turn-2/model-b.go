// filename: main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", middleware(helloHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// middleware is a simple middleware function that logs any panics that occur during the request processing.
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic occurred: %v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

// helloHandler is an example handler that panics if the user agent is Chrome.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() == "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36" {
		panic("User-Agent is Chrome!")
	}
	fmt.Fprintf(w, "Hello, World! Your User-Agent is: %s\n", r.UserAgent())
}
