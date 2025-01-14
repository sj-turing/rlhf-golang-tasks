package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type DistanceResponse struct {
	Distance float64 `json:"distance"`
}

func Distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Hypot(dx, dy)
}

func GetDistance(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	x1Str := r.FormValue("x1")
	x2Str := r.FormValue("x2")
	y1Str := r.FormValue("y1")
	y2Str := r.FormValue("y2")

	x1, err := strconv.ParseFloat(x1Str, 64)
	if err != nil {
		http.Error(w, "Invalid 'x1' parameter", http.StatusBadRequest)
		return
	}

	x2, err := strconv.ParseFloat(x2Str, 64)
	if err != nil {
		http.Error(w, "Invalid 'x2' parameter", http.StatusBadRequest)
		return
	}

	y1, err := strconv.ParseFloat(y1Str, 64)
	if err != nil {
		http.Error(w, "Invalid 'y1' parameter", http.StatusBadRequest)
		return
	}

	y2, err := strconv.ParseFloat(y2Str, 64)
	if err != nil {
		http.Error(w, "Invalid 'y2' parameter", http.StatusBadRequest)
		return
	}

	p1 := Point{X: x1, Y: y1}
	p2 := Point{X: x2, Y: y2}
	distance := Distance(p1, p2)

	response := DistanceResponse{Distance: distance}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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

func main() {
	http.HandleFunc("/getDistance", GetDistance)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
