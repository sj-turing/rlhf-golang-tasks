package main

import (
	"fmt"
)

// Define a struct
type Person struct {
	Name    string
	Age     int
	Country string
}

// Implement the Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Country: %s", p.Name, p.Age, p.Country)
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		Name:    "Alice",
		Age:     30,
		Country: "Wonderland",
	}

	// Use fmt.Sprintf to format and print the struct
	formattedString := fmt.Sprintf("Hello, %s! You are %d years old and from %s.", person, person.Age, person.Country)
	fmt.Println(formattedString)

	// Print the struct directly
	fmt.Println(person)
}
