package main

import (
	"encoding/json"
	"fmt"
)

// Define a sample Go struct representing some JSON data
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		Name: "John Doe",
		Age:  30,
		City: "New York",
	}

	// Encode the struct into JSON data
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Print the raw JSON data
	fmt.Println("Raw JSON Data:")
	fmt.Println(string(jsonData))

	// Format the JSON string using fmt.Sprintf
	formattedJSON := fmt.Sprintf("{\n\t\"name\": \"%s\",\n\t\"age\": %d,\n\t\"city\": \"%s\"\n}",
		person.Name, person.Age, person.City)

	// Print the formatted JSON data
	fmt.Println("\nFormatted JSON Data:")
	fmt.Println(formattedJSON)
}
