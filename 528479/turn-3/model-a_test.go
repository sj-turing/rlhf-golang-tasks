package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkGetUsersByAddressOriginal(b *testing.B) {
	userService := NewUserService()
	// Populate with test data
	for i := 0; i < 1000; i++ {
		userService.users[string(i)] = &User{Username: string(i), Address: "abc"}
	}

	req, _ := http.NewRequest("GET", "/getusersbyaddress?address=abc", nil)
	rr := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		userService.GetUsersByAddress(rr, req)
	}
}

func BenchmarkGetUsersByAddressOptimized(b *testing.B) {
	userService := NewUserService()
	// Populate with test data
	for i := 0; i < 1000; i++ {
		user := &User{Username: string(i), Address: "abc"}
		userService.users[string(i)] = user
		userService.addressIndex["abc"] = append(userService.addressIndex["abc"], user)
	}

	req, _ := http.NewRequest("GET", "/getusersbyaddress?address=abc", nil)
	rr := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		userService.GetUsersByAddress(rr, req)
	}
}
