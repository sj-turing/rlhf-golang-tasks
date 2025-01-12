package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type notification struct {
	method  string
	message string
}

func sendNotification(notificationChan chan<- notification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.URL.Query().Get("method")
		message := r.URL.Query().Get("message")
		if method == "" || message == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request parameters")
			return
		}
		notificationChan <- notification{method, message}
		fmt.Fprintln(w, "Notification sent successfully!")
	}
}

func processNotifications(ctx context.Context, notifications <-chan notification) {
	for {
		select {
		case n := <-notifications:
			switch strings.ToLower(n.method) {
			case "sms":
				sendSMS(n.message)
			case "email":
				sendEmail(n.message)
			default:
				log.Println("Invalid notification method:", n.method)
			}
		case <-ctx.Done():
			log.Println("Notification processing exiting...")
			return
		}
	}
}

func sendSMS(message string) {
	// Implement SMS sending logic
	fmt.Println("Sending SMS:", message)
}

func sendEmail(message string) {
	// Implement Email sending logic
	fmt.Println("Sending Email:", message)
}

func main() {
	notificationChan := make(chan notification)
	ctx, cancel := context.WithCancel(context.Background())

	go processNotifications(ctx, notificationChan)

	svr := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/notify", sendNotification(notificationChan))

	go func(svr *http.Server) {
		fmt.Println("Notification system running...")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while svr.ListenAndServe: %v", err)
		}
	}(svr)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	log.Println("Received SIGTERM. Gracefully shutting down...")

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(childCtx); err != nil {
		fmt.Printf("Error while svr.Shutdown: %v", err)

		// in case of failure to graceful shutdown
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("Executing force shutdown")
			// trying to forcefully shutdown
			if err := svr.Close(); err != nil {
				fmt.Printf("Server close failed: %v", err)
				fmt.Printf("Proceeding for server exit")
			} else {
				fmt.Printf("Server close completed")
			}
		}
	}

	close(notificationChan)
	cancel()
	log.Println("Graceful shutdown complete.")
	os.Exit(0)
}
