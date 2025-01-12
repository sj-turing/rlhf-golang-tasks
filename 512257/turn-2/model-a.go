package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
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

		// Validate parameters
		if method == "" || message == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request parameters")
			return
		}

		// Convert method to lowercase
		method = strings.ToLower(method)

		// Dispatching notifications to channel
		notificationChan <- notification{method: method, message: message}
		fmt.Fprintln(w, "Notification sent successfully!")
	}
}

func processNotifications(ctx context.Context, notifications <-chan notification) {
	defer close(notifications)

	for n := range notifications {
		if n.method == "sms" {
			sendSMS(n.message)
		} else if n.method == "email" {
			sendEmail(n.message)
		} else {
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

	// Create a cancel context
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		cancel()
	}()

	// Start the notification processor
	go processNotifications(ctx, notificationChan)

	// Start the server
	fmt.Println("Notification system running...")
	http.HandleFunc("/notify", sendNotification(notificationChan))
	s := &http.Server{
		Addr: ":8080",
	}

	// Handle OS termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for s := range quit {
			fmt.Println("Received terminate signal:", s)
			log.Println("Shutting down server gracefully...")
			ctx.Done() // Cancel the context
			s.Shutdown(ctx)
			os.Exit(0)
		}
	}()

	log.Fatal(s.ListenAndServe()) // blocks until Shutdown or ListenAndServe throws an error
}
