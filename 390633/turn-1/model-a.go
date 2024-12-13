package main

import (
	"flag"
	"fmt"
	"os"
)

// Application logic using boolean flags
func main() {
	// define boolean flags
	debugPtr := flag.Bool("debug", false, "enable debugging logs")
	verbosePtr := flag.Bool("verbose", false, "enable verbose logging")
	silentPtr := flag.Bool("silent", false, "disable all logging")

	// parse flags
	flag.Parse()

	// Validate flags: make sure silent is not used with others
	if *silentPtr && (*debugPtr || *verbosePtr) {
		fmt.Fprintf(os.Stderr, "Error: 'silent' flag cannot be used with 'debug' or 'verbose' flags.\n")
		os.Exit(1)
	}

	// Implementing the boolean logic
	if *silentPtr {
		fmt.Println("All logging is disabled.")
		return
	}

	if *debugPtr {
		fmt.Println("Debugging is enabled.")
	}

	if *verbosePtr {
		fmt.Println("Verbose logging is enabled. Detailed log info will be displayed.")
	}

	// Application logic...
	fmt.Println("Running application logic...")
}
