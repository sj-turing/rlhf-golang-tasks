package main

import (
	"encoding/json"
	"fmt"
)

// Define a sample Go struct representing some JSON data
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	City    string   `json:"city"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		Name:    "Jane Smith",
		Age:     28,
		City:    "Los Angeles",
		Hobbies: []string{"reading", "hiking", "coding"},
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

	// Basic formatting
	formattedJSON1 := fmt.Sprintf(`{"name": "%s", "age": %d, "city": "%s", "hobbies": %v}`,
		person.Name, person.Age, person.City, person.Hobbies)

	fmt.Println("\nBasic Formatted JSON:")
	fmt.Println(formattedJSON1)

	// Indentation with newlines
	formattedJSON2 := fmt.Sprintf(`{
		"name": "%s",
		"age": %d,
		"city": "%s",
		"hobbies": %v
	}`,
		person.Name, person.Age, person.City, person.Hobbies)

	fmt.Println("\nIndented JSON:")
	fmt.Println(formattedJSON2)

	// Indentation with tabs
	formattedJSON3 := fmt.Sprintf(`{
\t"name": "%s",
\t"age": %d,
\t"city": "%s",
\t"hobbies": %v
}`,
		person.Name, person.Age, person.City, person.Hobbies)

	fmt.Println("\nIndented with Tabs JSON:")
	fmt.Println(formattedJSON3)

	// Indentation with alignment
	formattedJSON4 := fmt.Sprintf(`{
\t"name":        "%s",
\t"age":         %d,
\t"city":        "%s",
\t"hobbies":     %v
}`,
		person.Name, person.Age, person.City, person.Hobbies)

	fmt.Println("\nAligned JSON:")
	fmt.Println(formattedJSON4)
}
