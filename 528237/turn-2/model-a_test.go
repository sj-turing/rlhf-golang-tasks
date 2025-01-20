package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// UserDetail represents the user details.
type UserDetail struct {
	Username  string
	Email     string
	LastVisit time.Time
}

// Mock datastore of user details for simulation purposes.
var userData = generateUserData(100000)

// generateUserData creates mock user data for testing.
func generateUserData(num int) []UserDetail {
	data := make([]UserDetail, num)
	for i := range data {
		data[i] = UserDetail{
			Username:  "User" + strconv.Itoa(i),
			Email:     "user" + strconv.Itoa(i) + "@example.com",
			LastVisit: time.Now().Add(-time.Duration(rand.Intn(24)) * time.Hour),
		}
	}
	return data
}

// GetUserDetails returns details of users who visited within the last 12 hours.
func GetUserDetails() []UserDetail {
	cutoff := time.Now().Add(-12 * time.Hour)
	var recentUsers []UserDetail
	for _, user := range userData {
		if user.LastVisit.After(cutoff) {
			recentUsers = append(recentUsers, user)
		}
	}
	return recentUsers
}

// BenchmarkGetUserDetails benchmarks the GetUserDetails function.
func BenchmarkGetUserDetails(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = GetUserDetails()
	}
}

// Profiling and debugging could be added here or enabled in test files.
func main() {
	// Example usage of GetUserDetails
	users := GetUserDetails()
	println("Number of recent users:", len(users))

	// Example usage of runtime/pprof for manual profiling could be added
}
