package main

import (
	"log"
	"net/http"
	"testing"
)

func deferHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer w.Body.Close()

	_, err := io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func manualHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer w.Body.Close()

	err := r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = w.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestDeferHandler(t *testing.T) {
	http.HandleFunc("/", deferHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func TestManualHandler(t *testing.T) {
	http.HandleFunc("/", manualHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
