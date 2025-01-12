package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

type notification struct {
	method  string
	message string
}

func sendNotification(notificationChan chan<- notification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := strings.ToLower(r.URL.Query().Get("method"))
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

func processNotifications(ctx context.Context, wg *sync.WaitGroup, notifications <-chan notification) {
	defer wg.Done()
	for {
		select {
		case n := <-notifications:
			switch n.method {
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
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go processNotifications(ctx, &wg, notificationChan)

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM)
		<-sigChan
		log.Println("Received SIGTERM. Gracefully shutting down...")
		close(notificationChan)
		cancel()
		wg.Wait()
		log.Println("Graceful shutdown complete.")
		os.Exit(0)
	}()

	fmt.Println("Notification system running...")
	http.HandleFunc("/notify", sendNotification(notificationChan))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
