package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"464944/turn-3/model-b/dispose"
)

type FileCloser struct {
	*os.File
}

func (fc *FileCloser) Close() error {
	return fc.File.Close()
}

type HttpServer struct {
	*http.Server
}

func (s *HttpServer) Close() error {
	return s.Server.Shutdown(nil)
}

func main() {
	d := dispose.NewDisposer()

	// Open a file for example
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	d.Add(&FileCloser{file})

	// Start a simple HTTP server for example
	server := &HttpServer{&http.Server{Addr: ":8080"}}
	go func() {
		fmt.Println("HTTP server started on :8080")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()
	d.Add(server)

	// Simulate doing some work
	fmt.Println("Doing some work...")
	time.Sleep(2 * time.Second)

	// Now, trigger graceful shutdown
	if err := d.Dispose(); err != nil {
		panic(err)
	}

	fmt.Println("Shutdown completed gracefully")
}
