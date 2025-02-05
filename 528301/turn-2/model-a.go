package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AuditTrail struct {
	UserID     string `json:"userID"`
	Action     string `json:"action"`
	ReceivedAt int64  `json:"receivedAt"`
}

func main() {
	http.HandleFunc("/track-transactions-lifecycle", auditMiddleware(TrackTransactionLifecycleHandler))
	http.HandleFunc("/track-subscription-browsing", auditMiddleware(trackSubscriptionBrowsingHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Middleware function to log request details
func auditMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request details
		log.Printf("Request: %s %s, Processed in %v\n", r.Method, r.URL.Path, time.Since(start))
	}
}

func TrackTransactionLifecycleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	encodedString := queryParams.Get("encoding")

	decodedByte, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendToAudit(decodedByte)
	w.WriteHeader(http.StatusOK)
}

func trackSubscriptionBrowsingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	encodedString := queryParams.Get("encoding")

	decodedByte, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendToAudit(decodedByte)
	w.WriteHeader(http.StatusOK)
}

func sendToAudit(data []byte) {
	var auditTrail AuditTrail
	json.Unmarshal(data, &auditTrail)
	auditTrail.ReceivedAt = time.Now().UnixMilli()

	// Log audit trail data
	log.Printf("Audit Data: %+v\n", auditTrail)
}
