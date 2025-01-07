package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerDefer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func handlerManual(w http.ResponseWriter, r *http.Request) {
	err := r.Body.Close()
	if err != nil {
		// Handle error
	}
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func BenchmarkHandlerDefer(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(handlerDefer))
	defer ts.Close()

	for i := 0; i < b.N; i++ {
		resp, err := http.Get(ts.URL)
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}

func BenchmarkHandlerManual(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(handlerManual))
	defer ts.Close()

	for i := 0; i < b.N; i++ {
		resp, err := http.Get(ts.URL)
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
