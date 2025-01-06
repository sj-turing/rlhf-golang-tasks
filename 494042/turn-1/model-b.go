package main

import (
	"errors"
	"fmt"
)

type ErrorCallback func(error)

func ProcessData(data string, callback ErrorCallback) (result string, err error) {
	// Simulate an error condition
	if data == "invalid" {
		err = errors.New("Invalid data provided")
		callback(err)
		return "", err
	}
	result = fmt.Sprintf("Processed data: %s", data)
	return result, nil
}
func main() {
	data := "invalid" // Replace this with valid data to avoid errors
	errorCallback := func(err error) {
		// Handle error here, e.g., log it, send notification, etc.
		fmt.Println("Error:", err.Error())
	}
	result, err := ProcessData(data, errorCallback)
	if err != nil {
		// Handle error as needed
		fmt.Println("Error occurred:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}
}
