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

func prettifyJSON(v interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
func main() {
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	// Unprettified JSON output
	unprettifiedData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Unprettified JSON:", string(unprettifiedData))

	// Prettified JSON output
	prettifiedData, err := prettifyJSON(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nPrettified JSON:")
	fmt.Println(string(prettifiedData))
}
