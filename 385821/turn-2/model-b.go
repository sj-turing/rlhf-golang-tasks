package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/gorilla/mux"
)

func doSomethingExpensive() {
	time.Sleep(time.Second)
}

// A simple service function that uses the context
func someServiceFunction(ctx context.Context, reqID string) error {
	doSomethingExpensive() // Simulate some work
	// Generate an error for demonstration purposes
	return fmt.Errorf("An error occurred for request %s: something broke", reqID)
}

// HTTP handler function that calls the service function
func handleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqID := vars["request_id"]

	ctx := context.WithValue(r.Context(), "request_id", reqID)

	err := someServiceFunction(ctx, reqID)
	if err != nil {
		logErrorWithContext(ctx, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request %s processed successfully", reqID)
}

// Function to log errors along with context data
func logErrorWithContext(ctx context.Context, err error) {
	reqID := ctx.Value("request_id")
	log.Printf("Error for request %s: %v", reqID, err)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/process/{request_id}", handleRequest)

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
