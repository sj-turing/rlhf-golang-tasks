package main

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"testing"
)

// Distance calculates the Euclidean distance between two points on a coordinate plane.
func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// TestDistance tests the Distance function with various test cases.
func TestDistance(t *testing.T) {
	tests := []struct {
		x1, y1, x2, y2 float64
		want           float64
	}{
		{0, 0, 3, 4, 5},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 3, 3},
		{3, 4, 3, 4, 0},
		{1.5, 2.5, 5.5, 7.5, 5},
	}

	for _, test := range tests {
		got := Distance(test.x1, test.y1, test.x2, test.y2)
		if !math.IsNaN(got) && !math.IsNaN(test.want) && math.Abs(got-test.want) > 0.0001 {
			t.Errorf("Distance(%f, %f, %f, %f) = %f, want %f", test.x1, test.y1, test.x2, test.y2, got, test.want)
		}
	}
}

// TestApiDistance validates the distance calculation via an API endpoint.
func TestApiDistance(t *testing.T) {
	// Assume API URL
	apiURL := "http://localhost:8080/distance"

	tests := []struct {
		x1, y1, x2, y2 float64
		want           float64
	}{
		{0, 0, 3, 4, 5},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 3, 3},
		{3, 4, 3, 4, 0},
		{1.5, 2.5, 5.5, 7.5, 5},
	}

	client := &http.Client{}

	for _, test := range tests {
		requestBody, err := json.Marshal(map[string]float64{
			"x1": test.x1,
			"y1": test.y1,
			"x2": test.x2,
			"y2": test.y2,
		})
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
			continue
		}

		var response map[string]float64
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		got, ok := response["distance"]
		if !ok {
			t.Fatalf("Response missing 'distance' field")
		}

		if !math.IsNaN(got) && !math.IsNaN(test.want) && math.Abs(got-test.want) > 0.0001 {
			t.Errorf("API Distance(%f, %f, %f, %f) = %f, want %f", test.x1, test.y1, test.x2, test.y2, got, test.want)
		}
	}
}

func main() {
	testing.Main()
}
