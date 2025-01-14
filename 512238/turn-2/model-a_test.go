package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDistance(t *testing.T) {
	testCases := []struct {
		name             string
		url              string
		expectedStatus   int
		expectedDistance float64
	}{
		{
			name:             "Origin to Origin",
			url:              "/distance?x1=0&y1=0&x2=0&y2=0",
			expectedStatus:   http.StatusOK,
			expectedDistance: 0,
		},
		{
			name:             "Positive X and Y",
			url:              "/distance?x1=3&y1=4&x2=0&y2=0",
			expectedStatus:   http.StatusOK,
			expectedDistance: 5,
		},
		{
			name:             "Negative X and Y",
			url:              "/distance?x1=-3&y1=-4&x2=0&y2=0",
			expectedStatus:   http.StatusOK,
			expectedDistance: 5,
		},
		{
			name:             "Positive X and Negative Y",
			url:              "/distance?x1=3&y1=-4&x2=0&y2=0",
			expectedStatus:   http.StatusOK,
			expectedDistance: 5,
		},
		{
			name:             "Negative X and Positive Y",
			url:              "/distance?x1=-3&y1=4&x2=0&y2=0",
			expectedStatus:   http.StatusOK,
			expectedDistance: 5,
		},
		{
			name:             "All points with non-zero value",
			url:              "/distance?x1=7&y1=13&x2=10&y2=17",
			expectedStatus:   http.StatusOK,
			expectedDistance: 5,
		},
		{
			name:           "Invalid coordinates",
			url:            "/distance?x1=abc&y1=def&x2=ghi&y2=jkl",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing coordinates",
			url:            "/distance?x1=3&y1=4",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.url, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetDistance)
			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("status code: got %v, want %v", rr.Code, tc.expectedStatus)
			}

			if rr.Code == http.StatusOK {
				var response map[string]float64
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal JSON response: %v", err)
				}

				if response["distance"] != tc.expectedDistance {
					t.Errorf("distance: got %v, want %v", response["distance"], tc.expectedDistance)
				}
			}
		})
	}
}
