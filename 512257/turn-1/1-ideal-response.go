package main

import (
	"fmt"
	"log"
	"net/http"
)

type notification struct {
	method  string
	message string
}

func sendNotification(notificationChan chan<- notification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Query parameters for method and message
		method := r.URL.Query().Get("method")
		message := r.URL.Query().Get("message")

		//Validate parameters
		if method == "" || message == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request parameters")
			return
		}
		// Dispatching notifications for each channel
		notificationChan <- notification{method: method, message: message}
		fmt.Fprintln(w, "Notification sent successfully!")
	}
}

func processNotifications(notifications <-chan notification) {

	for n := range notifications {
		switch n.method {
		case "sms":
			sendSMS(n.message)
		case "email":
			sendEmail(n.message)
		default:
			log.Println("Invalid notification method")
		}
	}
}

func sendSMS(message string) {
	// Implementation for sending SMS.
	fmt.Println("Sending SMS:", message)
}

func sendEmail(message string) {
	// Implementation for sending Email.
	fmt.Println("Sending Email:", message)
}

func main() {

	notificationChan := make(chan notification)
	go processNotifications(notificationChan)

	fmt.Println("Notification system running...")
	http.HandleFunc("/notify", sendNotification(notificationChan))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
