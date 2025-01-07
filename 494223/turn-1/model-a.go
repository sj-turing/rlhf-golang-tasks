package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserPreferences struct {
	Username string `json:"username"`
	Theme    string `json:"theme"`
	Language string `json:"language"`
	// Add more fields as needed
}

func readPreferences(filename string, encoding string) (*UserPreferences, error) {
	var pref UserPreferences
	var err error

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	switch encoding {
	case "utf-8":
		err = json.Unmarshal(data, &pref)
	case "ascii":
		err = json.Unmarshal(data, &pref)
	default:
		return nil, fmt.Errorf("unsupported encoding: %s", encoding)
	}

	if err != nil {
		return nil, err
	}

	return &pref, nil
}

func writePreferences(filename string, encoding string, pref UserPreferences) error {
	var err error

	switch encoding {
	case "utf-8":
		data, err := json.MarshalIndent(pref, "", "    ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, data, 0644)
	case "ascii":
		data, err := json.MarshalIndent(pref, "", "    ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, data, 0644)
	default:
		return fmt.Errorf("unsupported encoding: %s", encoding)
	}

	return err
}

func main() {
	filename := "user_preferences.json"
	encoding := "utf-8"

	// Example preference
	pref := UserPreferences{
		Username: "johndoe",
		Theme:    "dark",
		Language: "español",
	}

	// Write preferences to file
	err := writePreferences(filename, encoding, pref)
	if err != nil {
		fmt.Println("Error writing preferences:", err)
		return
	}

	// Read preferences from file
	readPref, err := readPreferences(filename, encoding)
	if err != nil {
		fmt.Println("Error reading preferences:", err)
		return
	}

	// Display read preferences
	fmt.Println("Read Preferences:")
	fmt.Printf("Username: %s\n", readPref.Username)
	fmt.Printf("Theme: %s\n", readPref.Theme)
	fmt.Printf("Language: %s\n", readPref.Language)
}
