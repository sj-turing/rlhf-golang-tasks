package main

import (
	"fmt"
	"log"
	"time"
)

// LogEntry represents a log entry for user activities
type LogEntry struct {
	Timestamp time.Time
	UserID    int
	Action    string
}

// String returns a custom formatted string representation of a LogEntry
func (l LogEntry) String() string {
	// Use fmt.Sprintf to format the timestamp, user ID, and action
	return fmt.Sprintf("%s | UserID: %d | Action: %s",
		l.Timestamp.Format("2006-01-02 15:04:05"), // Format the timestamp
		l.UserID,
		l.Action)
}

func main() {
	// Create a log entry
	entry := LogEntry{
		Timestamp: time.Now(),
		UserID:    123,
		Action:    "Login",
	}

	// Print the log entry using fmt.Println
	fmt.Println(entry) // Outputs: 2023-10-11 12:00:00 | UserID: 123 | Action: Login

	// Log the entry using the standard log package, which uses the String method
	log.Println(entry) // Outputs to the standard output with the same format

	// You can also format the log entry explicitly with fmt.Sprintf
	formattedString := fmt.Sprintf("User %d %s at %s", entry.UserID, entry.Action, entry.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Println(formattedString) // Outputs: User 123 Login at 2023-10-11 12:00:00
}
