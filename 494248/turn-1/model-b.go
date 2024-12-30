package main  
import (  
    "fmt"
    "os"
    "sync"
    "io/ioutil"
    "encoding/json"
    "time"
    "log"
    "sync/atomic"
    "path/filepath"
)

// Notification represents a notification message.
type Notification struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
}

// NotificationSystem manages file-based notifications.
type NotificationSystem struct {
    fileName string // name of the file to store notifications
    mutex    sync.Mutex // mutex for file access
    counter  uint64 // atomic counter for notification IDs
    fileLock sync.Mutex // lock for exclusive file access
}

// NewNotificationSystem creates a new NotificationSystem.
func NewNotificationSystem(fileName string) *NotificationSystem {
    return &NotificationSystem{
        fileName: fileName,
    }
}

// AddNotification adds a new notification to the system.
func (ns *NotificationSystem) AddNotification(text string) error {
    // Increment the atomic counter and get the new ID
    id := atomic.AddUint64(&ns.counter, 1)

    n := Notification{
        ID:   int(id),
        Text: text,
    }

    // Serialize the notification to JSON
    data, err := json.Marshal(n)
    if err != nil {
        return err
    }

    // Acquire the file lock for exclusive access
    ns.fileLock.Lock()
    defer ns.fileLock.Unlock()

    // Write the notification to the file
    f, err := os.OpenFile(ns.fileName, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    _, err = f.WriteString("\n")
    if err != nil {
        return err
    }

    return nil
}

// GetNotifications reads all notifications from the file.
func (ns *NotificationSystem) GetNotifications() ([]Notification, error) {
    // Acquire the file lock for shared access
    ns.fileLock.Lock()
    defer ns.fileLock.Unlock()

    // Read the entire file content
    data, err := ioutil.ReadFile(ns.fileName)
    if err != nil {
        return nil, err
    }

    // Split the content into lines (notifications)
    lines := bytes.Split(data, []byte("\n"))

    var notifications []Notification
    for _, line := range lines {
        if len(line) == 0 {
            continue
        }

        var n Notification
        err := json.Unmarshal(line, &n)
        if err != nil {
            return nil, err
        }
        notifications = append(notifications, n)
    }

    return notifications, nil
}

// DeleteNotification deletes a notification by ID.
func (ns *NotificationSystem) DeleteNotification(id int) error {
    // Acquire the file lock for exclusive access
    ns.fileLock.Lock()
    defer ns.fileLock.Unlock()

    // Read all notifications
    notifications, err := ns.GetNotifications()
    if err != nil {
        return err
    }

    // Create a new temporary file
    tempFile, err := ioutil.TempFile(filepath.Dir(ns.fileName), "notifications_temp_")
    if err != nil {
        return err
    }
    defer os.Remove(tempFile.Name())

    // Write all notifications except the one to be deleted to the temporary file
    for _, n := range notifications {
        if n.ID != id {
            data, err := json.Marshal(n)
            if err != nil {
                return err
            }
            _, err = tempFile.Write(data)
            if err != nil {
                return err
            }
            _, err = tempFile.WriteString("\n")
            if err != nil {
