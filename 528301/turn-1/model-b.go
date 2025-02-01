package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// AuditTrail is a struct to store audit trail entries
type AuditTrail struct {
	sync.Mutex          // to ensure thread safety when multiple goroutines are accessing the audit trail simultaneously
	Entries    []string // slice of audit trail entries
}

// logEntry is a struct to represent an audit trail entry
type logEntry struct {
	IP                string
	RequestURI        string
	QueryParameters   url.Values
	Encoding          string
	Decoding          string
	StatusCode        int
	TransactionStatus string
}

// NewAuditTrail creates a new AuditTrail instance
func NewAuditTrail() *AuditTrail {
	return &AuditTrail{Entries: make([]string, 0)}
}

// Log logs an audit trail entry
func (a *AuditTrail) Log(entry logEntry) {
	a.Lock()
	defer a.Unlock()

	// Format the audit trail entry as a string and append it to the Entries slice
	entryStr := fmt.Sprintf("IP: %s, RequestURI: %s, QueryParameters: %v, Encoding: %s, Decoding: %s, StatusCode: %d, TransactionStatus: %s", entry.IP, entry.RequestURI, entry.QueryParameters, entry.Encoding, entry.Decoding, entry.StatusCode, entry.TransactionStatus)
	a.Entries = append(a.Entries, entryStr)
}

// Print prints all the audit trail entries
func (a *AuditTrail) Print() {
	a.Lock()
	defer a.Unlock()

	for _, entry := range a.Entries {
		fmt.Println(entry)
	}
}

// decodeQueryParams decodes URL query parameters using given encoding and returns the decoded values
func decodeQueryParams(values url.Values, encoding string) (url.Values, error) {
	// ... Decoding logic goes here ...
	// For simplicity, let's assume we are decoding URL-encoded values
	decodedValues := make(url.Values)
	for key, values := range values {
		decodedValues[key] = make([]string, len(values))
		for i, value := range values {
			decoded, err := url.QueryUnescape(value)
			if err != nil {
				return nil, err
			}
			decodedValues[key][i] = decoded
		}
	}
	return decodedValues, nil
}

// httpHandler is a HTTP handler function that tracks URL query parameter encoding and decoding
func httpHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the IP address from the request context
	ip := r.RemoteAddr

	// Extract the RequestURI, QueryParameters, and Encoding
	requestURI := r.RequestURI
	queryParameters := r.URL.Query()
	encoding := r.URL.Query().Get("encoding")

	// Decode the query parameters using the given encoding
	decodedQueryParameters, err := decodeQueryParams(queryParameters, encoding)
	if err != nil {
		// Handle decoding error
		log.Printf("Error decoding query parameters: %v", err)
		http.Error(w, "Error decoding query parameters", http.StatusBadRequest)
		return
	}

	// Perform the required transaction or browsing logic here
	// For simplicity, let's assume we are just checking if a parameter "purchase" is present
	transactionStatus := "Failed"
	if _, ok := decodedQueryParameters["purchase"]; ok {
		transactionStatus = "Successful"
	}

	// Create an audit trail entry and log it
	auditTrailEntry := logEntry{
		IP:                ip,
		RequestURI:        requestURI,
		QueryParameters:   queryParameters,
		Encoding:          encoding,
		Decoding:          "URL-encoded", // Assuming URL-encoded for simplicity
		StatusCode:        http.StatusOK, // Assuming successful for simplicity
		TransactionStatus: transactionStatus,
	}
	auditTrail.Log(auditTrailEntry)

	// Return a response to the client
	fmt.Fprintf(w, "Transaction status: %s", transactionStatus)
}

func main() {
	// Create a new AuditTrail instance
	auditTrail = NewAuditTrail()

	// Start the HTTP server
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
