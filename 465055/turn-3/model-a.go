package main

import (
	"encoding/json"
	"fmt"
)

// Define a sample Go struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// PrettifyJSON takes a struct and returns its prettified JSON representation
func PrettifyJSON(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Use json.MarshalIndent to prettify the JSON
	prettifiedData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return "", err
	}

	return string(prettifiedData), nil
}

func main() {
	person := Person{
		Name: "John Doe",
		Age:  30,
		City: "New York",
	}

	// Get the prettified JSON
	prettifiedJSON, err := PrettifyJSON(person)
	if err != nil {
		fmt.Println("Error prettifying JSON:", err)
		return
	}

	// Print the prettified JSON
	fmt.Println("Prettified JSON:")
	fmt.Println(prettifiedJSON)
}
