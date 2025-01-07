package main

import (
	"fmt"
	"net/http"
	"testing"
)

// handler handles HTTP requests at the /hello endpoint.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// TestHttpServer tests that the HTTP server responds correctly.
func TestHttpServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)
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

	body, err := ioutil.ReadAll(res.Body)
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
	mux.HandleFunc("/hello", handler)
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

// ExampleHttpServer demonstrates starting the HTTP server.
func ExampleHttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	fmt.Println("Starting HTTP server on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
