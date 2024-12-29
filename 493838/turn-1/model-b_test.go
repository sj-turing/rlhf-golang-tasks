package main  
import (  
    "fmt"
    "sync"
    "testing"
    "time"
)

type Notification struct {  
    Id       int
    Message string
}
type Notifier struct {  
    notifications []Notification  
    mu          sync.Mutex
}

func (n *Notifier) AddNotification(notification Notification) {  
    n.mu.Lock()  
    defer n.mu.Unlock()  
    n.notifications = append(n.notifications, notification)  
}
func (n *Notifier) ProcessNotifications() {  
    n.mu.Lock()  
    defer n.mu.Unlock()  
    for _, notification := range n.notifications {  
        fmt.Printf("Processing notification: %+v\n", notification)  
    }  
    n.notifications = []Notification{}
}

func BenchmarkBatchProcessing(b *testing.B) {  
    notifier := NewNotifier()
