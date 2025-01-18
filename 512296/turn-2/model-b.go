package main

import (
	"flag"
	"log"
	"net/http"
	"runtime/pprof"
)

func main() {
	var pprofAddr string
	flag.StringVar(&pprofAddr, "pprof", "", "pprof address to listen on")
	flag.Parse()

	if pprofAddr != "" {
		go func() {
			log.Println(http.ListenAndServe(pprofAddr, nil))
		}()
	}

	// Your DNA matching algorithm code here
}
