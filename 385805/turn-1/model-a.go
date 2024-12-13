package main

import (
	"fmt"
	_ "github.com/lib/pq" // Import Postgres driver
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Database connection setup (not included in this example)

	// Profile the CPU for 30 seconds
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create profile: ", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Your main application logic goes here

	// Server setup and start
	http.HandleFunc("/", handleRequest)
	port := ":8080"
	fmt.Println("Server starting on", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Your database transaction logic
	dbTransaction()

	// Execute other parts of your request handling here

	elapsed := time.Since(start)
	fmt.Fprintf(w, "Request processed in %s\n", elapsed)
}
func dbTransaction() {
	// Sample database transaction code
	fmt.Println("This is dbTransaction")
}
