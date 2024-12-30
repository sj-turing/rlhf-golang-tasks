package main

/*

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"
)

// WeatherNotification represents a weather notification.
type WeatherNotification struct {
	UserID    string    `json:"user_id"`
	Weather   string    `json:"weather"`
	Timestamp time.Time `json:"timestamp"`
}

// NotificationSystem manages weather notifications.
type NotificationSystem struct {
	filePath string
	lock     *sync.Mutex
}

// NewNotificationSystem creates a new notification system.
func NewNotificationSystem(filePath string) *NotificationSystem {
	return &NotificationSystem{
		filePath: filePath,
		lock:     &sync.Mutex{},
	}
}

// Start starts the notification system.
func (ns *NotificationSystem) Start() {
	// In this simple implementation, we don't have a separate worker thread,
	// but you could add one to handle writes in the background.
}

// AddNotification adds a weather notification to the system.
func (ns *NotificationSystem) AddNotification(notification WeatherNotification) {
	ns.lock.Lock()
	defer ns.lock.Unlock()

	data, err := ioutil.ReadFile(ns.filePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		data = []byte{}
	}

	data = append(data, []byte(fmt.Sprintln(notification))...)

	err = ioutil.WriteFile(ns.filePath, data, 0644)
	if err != nil {
		log.Printf("Error writing file: %v", err)
	} else {
		log.Printf("Notification written for user %s: %s", notification.UserID, notification.Weather)
	}
}

func main() {
	ns := NewNotificationSystem("weather_notifications.txt")

	// Simulate multiple users adding notifications concurrently
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(userID string) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ns.AddNotification(WeatherNotification{
					UserID:    userID,
					Weather:   fmt.Sprintf("Weather for user %s: %s", userID, getRandomWeather()),
					Timestamp: time.Now(),
				})
				time.Sleep(time.Second)
			}
		}(fmt.Sprintf("user%d", i+1))
	}

	wg.Wait()
}

func getRandomWeather() string {
	weathers := []string{"Sunny", "Cloudy", "Rainy", "Snowy", "Thunderstorm"}
	return weathers[rand.Intn(len(weathers))]
}
*/
