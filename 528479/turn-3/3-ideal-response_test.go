package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	log.SetOutput(buf)
}

var address = []string{"city1", "city2", "city3"}

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func getRandomString(size int) string {
	buf := &bytes.Buffer{}
	defer buf.Reset()

	for size > 0 {
		buf.WriteByte(alphabets[rand.Intn(len(alphabets))])
		size--
	}

	return buf.String()
}

func BenchmarkGetUsersByAddress(b *testing.B) {
	userService := setup() // Setup users and indices (both old and new)

	b.Run("Old Implementation", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/getusersbyaddress?address=%s", address[rand.Intn(len(address))]), nil)
			w := httptest.NewRecorder()
			userService.GetUsersByAddress(w, req)
		}
	})

	b.Run("New Implementation", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/getusersbyaddressnew?address=%s", address[rand.Intn(len(address))]), nil)
			w := httptest.NewRecorder()
			userService.GetUsersByAddressOptimal(w, req)
		}
	})
}

func setup() *UserService {
	// Generate random users and populate both userService and userAddressIndex
	// ... (Your implementation to generate random users and populate the indices)
	//
	userService := NewUserService()
	size := 10000

	for size > 0 {
		size--
		user := User{
			Username: getRandomString(rand.Intn(10)),
			Password: getRandomString(rand.Intn(20)),
			Address:  address[rand.Intn(len(address))],
		}

		bb, _ := json.Marshal(user)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(bb))
		w := httptest.NewRecorder()
		userService.SignUp(w, req)
	}

	return userService

}
