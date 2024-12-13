package main

import (
	"flag"
	"fmt"
)

// Feature flags
var enableNewFeature = flag.Bool("enable-new-feature", false, "enable the experimental new feature")

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Core application logic
	fmt.Println("Running the application...")

	if *enableNewFeature {
		runNewFeature()
	} else {
		runOldFeature()
	}
}

// runNewFeature simulates implementation of the new feature
func runNewFeature() {
	fmt.Println("Running the new experimental feature!")
	// Add new feature logic here
}

// runOldFeature simulates the existing functionality
func runOldFeature() {
	fmt.Println("Running the old feature.")
	// Add existing feature logic here
}
