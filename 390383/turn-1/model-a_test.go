package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// Test Example
func TestGetUser(t *testing.T) {
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"id":1,"name":"John Doe"}`)),
	}
	mockClient := &MockHTTPClient{
		MockResponse: mockResponse,
		MockError:    nil,
	}

	apiClient := NewAPIClient(mockClient)
	user, err := apiClient.GetUser(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := &User{ID: 1, Name: "John Doe"}
	if *user != *expected {
		t.Errorf("Expected user to be %+v, got %+v", expected, user)
	}
}
