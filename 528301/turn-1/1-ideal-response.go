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

	http.HandleFunc("/track-transactions-lifecycle", TrackTransactionLifecycleHandler)
	http.HandleFunc("/track-subscription-browsing", trackSubscriptionBrowingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TrackTransactionLifecycleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	encodedString := queryParams["encoding"][0]

	decodedByte, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendToAudit(decodedByte)
	w.WriteHeader(http.StatusOK)
}

func trackSubscriptionBrowingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	encodedString := queryParams["encoding"][0]

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

	log.Printf("this is audit data = %+v\n", auditTrail)
}
