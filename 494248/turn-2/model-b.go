package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

type WeatherNotification struct {
	UserID   string `json:"user_id"`
	Message  string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	WeatherType string `json:"weather_type"`
}

type NotificationSystem struct {
	notifications chan WeatherNotification
	lock sync.Mutex
	notificationFileDir string
}

func NewNotificationSystem(dir string) *NotificationSystem {
	// Create the directory to store notifications if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			log.Fatalf("Failed to create notification directory: %v", err)
		}
	}
	return &NotificationSystem{
		notifications:       make(chan WeatherNotification),
		notificationFileDir: dir,
	}
}

func (ns *NotificationSystem) Start() {
	go ns.writeNotifications()
}

func (ns *NotificationSystem) AddNotification(notification WeatherNotification) {
	ns.notifications <- notification
}

func (ns *NotificationSystem) writeNotifications() {
	for notification := range ns.notifications {
		filePath := filepath.Join(ns.notificationFileDir, fmt.Sprintf("%s_notifications.txt", notification.WeatherType))
		err := ns.writeToFile(filePath, notification)
		if err != nil {
			log.Printf("Error writing notification: %v", err)
		}
	}
}

// writeToFile writes a single notification to a file while handling concurrent access.
func (ns *NotificationSystem) writeToFile(filePath string, notification WeatherNotification) error {
	// Lock the file to prevent multiple processes from writing simultaneously
	ns.lock.Lock()
	defer ns.lock.Unlock()

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		return err
	}
	defer func() {
		if err := syscall.Flock(int(file.Fd()), syscall.LOCK_UN); err != nil {
			log.Printf("Error unlocking file: %v", err)
		}
	}()

	data, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}

func main() {
	notificationDir := "weather_notifications"
	ns := NewNotificationSystem(notificationDir)
	ns.Start()

	go func() {
		for i := 0; i < 10; i++ {
			ns.AddNotification(WeatherNotification{
				UserID:       fmt.Sprintf("user%d", i),
				Message:      "It's going to rain!",
				Timestamp:    time.Now(),
				WeatherType: "Rain",
			})
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ns.AddNotification(WeatherNotification{
				UserID:       fmt.Sprintf("user%d", i+10),
				Message:      "Sunny skies ahead!",
				Timestamp:    time.Now(),
