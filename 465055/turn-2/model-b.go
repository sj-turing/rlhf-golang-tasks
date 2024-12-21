package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	// 1. Simple formatting with newlines
	formattedJSON1 := fmt.Sprintf("%s\n", string(jsonData))
	fmt.Println("Formatted JSON 1:")
	fmt.Println(formattedJSON1)

	// 2. Adding indentation using tabs
	formattedJSON2 := fmt.Sprintf("%s\n", string(jsonData))
	formattedJSON2 = fmt.Sprintf("%s\n", formattedJSON2)
	formattedJSON2 = fmt.Sprintf("%s\t%s", formattedJSON2, formattedJSON2)
	fmt.Println("\nFormatted JSON 2:")
	fmt.Println(formattedJSON2)

	// 3. Using a custom indentation string
	indent := "    "
	formattedJSON3 := fmt.Sprintf("%s\n", string(jsonData))
	formattedJSON3 = fmt.Sprintf("%s\n", formattedJSON3)
	formattedJSON3 = fmt.Sprintf("%s%s", formattedJSON3, indent)
	formattedJSON3 = fmt.Sprintf("%s%s", formattedJSON3, formattedJSON3)
	fmt.Println("\nFormatted JSON 3:")
	fmt.Println(formattedJSON3)

	// 4. Using json.MarshalIndent for better readability
	formattedJSON4, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nFormatted JSON 4 (using json.MarshalIndent):")
	fmt.Println(string(formattedJSON4))
}
