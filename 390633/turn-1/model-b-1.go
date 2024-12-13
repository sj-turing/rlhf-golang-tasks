package main

import (
	"flag"
	"fmt"
)

var debug = flag.Bool("debug", false, "enable debug mode")

func main() {
	flag.Parse()

	if *debug {
		fmt.Println("Debug mode is enabled.")
	} else {
		fmt.Println("Debug mode is disabled.")
	}
}
