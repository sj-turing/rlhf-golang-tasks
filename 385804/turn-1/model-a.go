package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func jsonSerialization(person Person) error {
	data, err := json.Marshal(person)
	if err != nil {
		return fmt.Errorf("error serializing JSON: %w", err)
	}

	err = os.WriteFile("person.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}
	return nil
}

func jsonDeserialization() (Person, error) {
	data, err := os.ReadFile("person.json")
	if err != nil {
		return Person{}, fmt.Errorf("error reading JSON file: %w", err)
	}

	var person Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return Person{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return person, nil
}

func xmlSerialization(person Person) error {
	xmlData, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		return fmt.Errorf("error serializing XML: %w", err)
	}

	xmlData = []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + string(xmlData))

	err = os.WriteFile("person.xml", xmlData, 0644)
	if err != nil {
		return fmt.Errorf("error writing XML to file: %w", err)
	}

	return nil
}

func xmlDeserialization() (Person, error) {
	data, err := os.ReadFile("person.xml")
	if err != nil {
		return Person{}, fmt.Errorf("error reading XML file: %w", err)
	}

	var person Person
	err = xml.Unmarshal(data, &person)
	if err != nil {
		return Person{}, fmt.Errorf("error unmarshalling XML: %w", err)
	}

	return person, nil
}

func main() {
	person := Person{
		Name: "Abak",
		Age:  30,
	}

	err := jsonSerialization(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON serialization successful")

	err = xmlSerialization(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("XML serialization successful")

	deserializedPerson, err := jsonDeserialization()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON deserialization successful:")
	fmt.Println(deserializedPerson)

	deserializedPerson, err = xmlDeserialization()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("XML deserialization successful:")
	fmt.Println(deserializedPerson)
}
