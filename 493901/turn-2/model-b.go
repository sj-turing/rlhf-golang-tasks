package main

import (
	"fmt"
	"net/http"
	"testing"
)

// HelloHandler responds to the "/hello" route with a message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// main serves as the entry point of the application.
func main() {
	http.HandleFunc("/hello", HelloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// BenchmarkHelloHandler benchmarks the HelloHandler function.
func BenchmarkHelloHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := NewMockResponseWriter()

	for i := 0; i < b.N; i++ {
		HelloHandler(w, req)
	}
}

type MockResponseWriter struct {
	code int
}

func (w *MockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (w *MockResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w *MockResponseWriter) WriteHeader(code int) {
	w.code = code
}

// ExampleHelloHandler shows an example usage of the HelloHandler function.
func ExampleHelloHandler() {
	w := NewMockResponseWriter()
	req, _ := http.NewRequest("GET", "/hello", nil)
	HelloHandler(w, req)

	fmt.Printf("Status code: %d\n", w.code)
	// Output: Status code: 200
}
