package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, secure world!"))
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Assuming you have your certificates ready
	log.Println("Starting server on :443 for HTTPS")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
