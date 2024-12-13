package main

/*

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type OTPRateLimiter struct {
	mu              sync.Mutex
	customerRecords map[string]*CustomerRecord
}

type CustomerRecord struct {
	hourlyCount       int
	dailyCount        int
	lastHourlyRequest time.Time
	lastDailyRequest  time.Time
}

const (
	hourLimit = 3
	dayLimit  = 10
)

func NewOTPRateLimiter() *OTPRateLimiter {
	return &OTPRateLimiter{
		customerRecords: make(map[string]*CustomerRecord),
	}
}

func (rl *OTPRateLimiter) RequestOTP(customerID string) (bool, string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	record, exists := rl.customerRecords[customerID]
	if !exists {
		// Create a new customer record if it does not exist
		record = &CustomerRecord{
			hourlyCount:       0,
			dailyCount:        0,
			lastHourlyRequest: now,
			lastDailyRequest:  now,
		}
		rl.customerRecords[customerID] = record
	}

	// Reset hourly limit if the last request was more than an hour ago
	if now.Sub(record.lastHourlyRequest) >= time.Hour {
		record.hourlyCount = 0
		record.lastHourlyRequest = now
	}

	// Reset daily limit if the last request was more than a day ago
	if now.Sub(record.lastDailyRequest) >= 24*time.Hour {
		record.dailyCount = 0
		record.lastDailyRequest = now
	}

	// Check hourly limit
	if record.hourlyCount < hourLimit {
		record.hourlyCount++
	} else {
		return false, "Rate limit exceeded: You can only request 3 OTPs per hour."
	}

	// Check daily limit
	if record.dailyCount < dayLimit {
		record.dailyCount++
	} else {
		return false, "Rate limit exceeded: You can only request 10 OTPs per day."
	}

	// Generate OTP Logic (This is a placeholder)
	otp := generateOTP()
	return true, otp
}

func generateOTP() string {
	// Replace with actual OTP generation logic
	return "123456"
}

func otpRequestHandler(rl *OTPRateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := r.URL.Query().Get("customer_id") // Get customer ID from request

		if customerID == "" {
			http.Error(w, "Customer ID is required", http.StatusBadRequest)
			return
		}

		allowed, response := rl.RequestOTP(customerID)
		if !allowed {
			http.Error(w, response, http.StatusTooManyRequests)
			return
		}

		fmt.Fprintln(w, "Your OTP is:", response)
	}
}

func main() {
	rateLimiter := NewOTPRateLimiter()
	http.HandleFunc("/request-otp", otpRequestHandler(rateLimiter))

	fmt.Println("OTP server running on :8080")
	http.ListenAndServe(":8080", nil)
}

*/
