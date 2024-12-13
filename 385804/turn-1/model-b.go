package main

/*
import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

	err = ioutil.WriteFile("person.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}
	return nil
}

func jsonDeserialization() (Person, error) {
	data, err := ioutil.ReadFile("person.json")
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
	data, err := xml.Marshal(person)
	if err != nil {
		return fmt.Errorf("error serializing XML: %w", err)
	}

	// Prepend XML declaration
	data = []byte(xml.Header) + data

	err = ioutil.WriteFile("person.xml", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing XML to file: %w", err)
	}
	return nil
}
func xmlDeserialization() (Person, error) {
	data, err := ioutil.ReadFile("person.xml")
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
	person := Person{Name: "John", Age: 30}

	if err := jsonSerialization(person); err != nil {
		fmt.Println("JSON serialization failed:", err)
		os.Exit(1)
	}

	if err := xmlSerialization(person); err != nil {
		fmt.Println("XML serialization failed:", err)
		os.Exit(1)
	}

	var jsonPerson, errJson = jsonDeserialization()
	if errJson != nil {
		fmt.Println("JSON deserialization failed:", errJson)
		os.Exit(1)
	}

	var xmlPerson, errXml = xmlDeserialization()
	if errXml != nil {
		fmt.Println("XML deserialization failed:", errXml)
		os.Exit(1)
	}

	fmt.Println("JSON Person:", jsonPerson)
	fmt.Println("XML Person:", xmlPerson)
}

*/
