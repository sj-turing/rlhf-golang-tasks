package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandler tests that the HTTP server responds correctly.
func TestHelloHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Errorf("Error fetching response: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading body: %v", err)
	}

	if string(body) != "Hello, World!" {
		t.Errorf("Expected body 'Hello, World!', got %s", string(body))
	}
}

// BenchmarkHttpServer benchmarks the HTTP server handler.
func BenchmarkHttpServer(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := &http.Client{}

	for i := 0; i < b.N; i++ {
		_, err := client.Get(ts.URL + "/hello")
		if err != nil {
			b.Errorf("Error performing request: %v", err)
		}
	}
}
