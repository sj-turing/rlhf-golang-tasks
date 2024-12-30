package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"syscall"
	"time"
)

// Notification represents a notification message.
type Notification struct {
	UserID    string    `json:"user_id"`
	Message   string    `json:"message"`
	Weather   string    `json:"weather"`
	Timestamp time.Time `json:"timestamp"`
}

// NotificationSystem manages notifications.
type NotificationSystem struct {
	filePath      string
	notifications chan Notification
	lock          sync.Mutex
}

// NewNotificationSystem creates a new notification system.
func NewNotificationSystem(filePath string) *NotificationSystem {
	return &NotificationSystem{
		filePath:      filePath,
		notifications: make(chan Notification),
	}
}

// Start starts the notification system, writing to the file concurrently.
func (ns *NotificationSystem) Start() {
	go ns.writeNotifications()
}

// AddNotification adds a notification to the system.
func (ns *NotificationSystem) AddNotification(notification Notification) {
	ns.notifications <- notification
}

// writeNotifications writes notifications to the file in a concurrent manner.
func (ns *NotificationSystem) writeNotifications() {
	file, err := os.OpenFile(ns.filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	flocker := &Flocker{file}

	for notification := range ns.notifications {
		err = flocker.Flock()
		if err != nil {
			log.Printf("Failed to flock file: %v", err)
			continue
		}

		defer flocker.Unflock()

		data, err := json.Marshal(notification)
		if err != nil {
			log.Printf("Failed to marshal notification: %v", err)
			continue
		}

		_, err = file.Write(append(data, '\n'))
		if err != nil {
			log.Printf("Failed to write notification to file: %v", err)
			continue
		}

		log.Printf("Notification written for user %s: %s", notification.UserID, notification.Message)
	}
}

// Flocker provides file locking functionality.
type Flocker struct {
	*os.File
}

// Flock locks the file.
func (f *Flocker) Flock() error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

// Unflock unlocks the file.
func (f *Flocker) Unflock() error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}

func main() {
	ns := NewNotificationSystem("weather_notifications.txt")
	ns.Start()

	go func() {
		for i := 0; i < 10; i++ {
			ns.AddNotification(Notification{
				UserID:    fmt.Sprintf("user%d", i),
				Message:   fmt.Sprintf("Message for user%d", i),
				Timestamp: time.Now(),
				Weather:   fmt.Sprintf("Weather for user %d: %s", i, getRandomWeather()),
			})
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ns.AddNotification(Notification{
				UserID:    fmt.Sprintf("user%d", i+10),
				Message:   fmt.Sprintf("Message for user%d", i+10),
				Weather:   fmt.Sprintf("Weather for user %d: %s", i+10, getRandomWeather()),
				Timestamp: time.Now(),
			})
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}

func getRandomWeather() string {
	weathers := []string{"Sunny", "Cloudy", "Rainy", "Snowy", "Thunderstorm"}
	return weathers[rand.Intn(len(weathers))]
}
