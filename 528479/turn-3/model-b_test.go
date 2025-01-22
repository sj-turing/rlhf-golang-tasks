package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func BenchmarkGetUsersByAddress(b *testing.B) {
	setup() // Setup users and indices (both old and new)

	b.Run("Old Implementation", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			req, _ := http.NewRequest(http.MethodGet, "/getusersbyaddress?address=address"+strconv.Itoa(rand.Intn(10000)), nil)
			w := httptest.NewRecorder()
			userService.GetUsersByAddress(w, req)
		}
	})

	b.Run("New Implementation", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			req, _ := http.NewRequest(http.MethodGet, "/getusersbyaddress?address=address"+strconv.Itoa(rand.Intn(10000)), nil)
			w := httptest.NewRecorder()
			userAddressIndex.GetUsersByAddress(w, req)
		}
	})
}

func setup() {
	// Generate random users and populate both userService and userAddressIndex
	// ... (Your implementation to generate random users and populate the indices)
}
