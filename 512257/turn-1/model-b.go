package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type notification struct {
	method  string
	message string
}

func sendNotification(w http.ResponseWriter, r *http.Request) {

	wg := sync.WaitGroup{}
	notifications := make(chan notification, 2)

	defer close(notifications)
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
	go func() {
		wg.Add(1)
		defer wg.Done()
		notifications <- notification{method: method, message: message}
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		processNotifications(notifications)
	}()
	wg.Wait()
	fmt.Fprintln(w, "Notification sent successfully!")
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
	fmt.Println("Notification system running...")
	http.HandleFunc("/notify", sendNotification)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
