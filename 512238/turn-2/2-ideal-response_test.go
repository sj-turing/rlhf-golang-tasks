package main

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDistanceHandler(t *testing.T) {
	testCases := []struct {
		name             string
		queryParams      map[string]string
		expectedStatus   int
		expectedResponse DistanceResponse
	}{
		{
			name:             "Valid parameters",
			queryParams:      map[string]string{"x1": "1", "y1": "2", "x2": "3", "y2": "4"},
			expectedStatus:   http.StatusOK,
			expectedResponse: DistanceResponse{Distance: math.Hypot(2, 2)},
		},
		{
			name:             "Invalid x1 parameter",
			queryParams:      map[string]string{"x1": "invalid", "y1": "2", "x2": "3", "y2": "4"},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: DistanceResponse{},
		},
		// Add more test cases for other invalid parameter scenarios
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/getDistance", nil)
			q := req.URL.Query()
			for key, value := range tt.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()
			GetDistance(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				var actualResponse DistanceResponse
				if err := json.NewDecoder(resp.Body).Decode(&actualResponse); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}

				if actualResponse.Distance != tt.expectedResponse.Distance {
					t.Errorf("Expected distance %f, got %f", tt.expectedResponse.Distance, actualResponse.Distance)
				}
			}
		})
	}
}
