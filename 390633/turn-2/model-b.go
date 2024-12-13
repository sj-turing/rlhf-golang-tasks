package main

import (
	"flag"
	"fmt"
	"strings"
)

var newProcessing = flag.Bool("new-processing", false, "enable the new data processing feature")

// Old processing function
func oldProcessData(data string) string {
	return strings.ToUpper(data)
}

// New processing function
func newProcessData(data string) string {
	return strings.Replace(data, " ", "-", -1)
}

func processData(data string) string {
	if *newProcessing {
		return newProcessData(data)
	} else {
		return oldProcessData(data)
	}
}

func main() {
	flag.Parse()

	inputData := "hello world"
	processedData := processData(inputData)

	fmt.Println("Input Data:", inputData)
	fmt.Println("Processed Data:", processedData)
}
