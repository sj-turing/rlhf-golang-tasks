package main

import (
	"errors"
	"fmt"
	"log"
)

type ErrorCallback func(error)

// Function processing data with a callback to handle errors.
func processDataWithCallback(data string, callback ErrorCallback) error {
	if data == "invalid" {
		err := errors.New("Invalid data provided")
		callback(err)
		return err
	}
	return nil
}

func main() {
	var data string = "invalid" // Replace with valid data

	// Define an error callback to log the error.
	errorLogger := func(err error) {
		log.Println("Error encountered:", err.Error())
	}

	// Processing the data with callback handling
	err := processDataWithCallback(data, errorLogger)
	if err != nil {
		// Optional: Also print or handle the error here
		fmt.Println("Error occurred:", err.Error())
	}
}
