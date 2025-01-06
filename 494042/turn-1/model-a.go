package main

import (
	"log"
	"net/http"
	"time"
)

// Callback is a function type representing the callback function
type Callback func(error)

// handleSecureOperation runs a secure operation with error handling and callback execution
func handleSecureOperation(operation func() error, callback Callback) {
	err := operation()
	if err != nil {
		log.Errorf("Operation failed: %v", err)
		callback(err)
		return
	}
	log.Infof("Operation completed successfully")
	callback(nil)
}

// secureOperation is a mock of a secure operation that might fail
func secureOperation() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Simulated secure operation error")
}

// handleErrorLogging is a callback function that logs the error
func handleErrorLogging(err error) {
	if err != nil {
		log.Printf("Error in callback: %v", err)
	} else {
		log.Println("Callback executed successfully")
	}
}

func main() {
	// Listen for incoming connections on port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Perform a secure operation with callback
		handleSecureOperation(secureOperation, handleErrorLogging)
	})

	log.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
