package main

import (
	"flag"
	"fmt"
)

// Define logging flags
var logGroup = struct {
	debug   *bool
	verbose *bool
}{
	debug:   flag.Bool("log.debug", false, "enable debug logging"),
	verbose: flag.Bool("log.verbose", false, "increase logging verbosity"),
}

// Define file handling flags
var fileGroup = struct {
	force *bool
}{
	force: flag.Bool("file.force", false, "overwrite existing files"),
}

func main() {
	flag.Parse()

	setupLogging(*logGroup.debug, *logGroup.verbose)
	setupFileHandling(*fileGroup.force)
}

func setupLogging(debug bool, verbose bool) {
	if debug {
		fmt.Println("Debug logging enabled.")
	}
	if verbose {
		fmt.Println("Verbose logging enabled.")
	}
}

func setupFileHandling(force bool) {
	if force {
		fmt.Println("File force mode enabled.")
	}
}
