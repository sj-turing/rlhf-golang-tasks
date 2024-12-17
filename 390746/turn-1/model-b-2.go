package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	certFile = flag.String("cert", "./cert.pem", "Path to SSL certificate")
	keyFile  = flag.String("key", "./key.pem", "Path to SSL key")
)

func main() {
	flag.Parse()

	if _, err := os.Stat(*certFile); os.IsNotExist(err) {
		log.Fatal("Certificate file not found:", *certFile)
	}

	if _, err := os.Stat(*keyFile); os.IsNotExist(err) {
		log.Fatal("Key file not found:", *keyFile)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, secure world! Method: %s\n", r.Method)
	})

	log.Println("Listening on :443 with HTTPS")
	if err := http.ListenAndServeTLS(":443", *certFile, *keyFile, nil); err != nil {
		log.Fatal(err)
	}
}
