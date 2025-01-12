import (
	"log"
	"sync"
	"time"
)

var notificationQueue = make(chan string, 100)
var wg sync.WaitGroup

func processSubscriptionEvent(userID, notificationType string) {
	// Add the event to the queue
	notificationQueue <- fmt.Sprintf("%s,%s", userID, notificationType)
	wg.Add(1)
	defer wg.Done()
}

func sendNotification(event string) {
	userID, notificationType := eventSplit(event)

	// Simulate sending a notification
	log.Printf("Sending %s notification to user %s\n", notificationType, userID)
	time.Sleep(1 * time.Second) // Simulate delay
}

func eventSplit(event string) (userID, notificationType string) {
	parts := strings.Split(event, ",")
	return parts[0], parts[1]
}

func main() {
	// Start notification processing goroutines
	for i := 0; i < 5; i++ {
		go processNotifications()
	}

	// Run the web server
	runServer()

	// Wait for all goroutines to finish
	wg.Wait()
}

func processNotifications() {
	for event := range notificationQueue {
		sendNotification(event)
	}
}
