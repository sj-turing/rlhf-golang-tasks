package main

import (
	"fmt"
	"time"
)

// FormatDuration formats a time.Duration into a human-readable string
func FormatDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24

	hours %= 24
	minutes %= 60
	seconds %= 60

	return fmt.Sprintf("%dd %02dh %02dm %02ds", days, hours, minutes, seconds)
}

func main() {
	duration := 2*time.Hour + 30*time.Minute + 15*time.Second
	fmt.Println(duration) // Output: 2h30m15s

	// Using our custom format function
	formattedDuration := FormatDuration(duration)
	fmt.Println(formattedDuration) // Output: 0d 02h 30m 15s
}
