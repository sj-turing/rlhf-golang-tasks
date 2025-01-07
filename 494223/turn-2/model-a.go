package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type UserPreferences struct {
	Username string `json:"username"`
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

func readPreferencesUTF8(filename string) (*UserPreferences, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var pref UserPreferences
	err = json.Unmarshal(data, &pref)
	if err != nil {
		return nil, err
	}

	return &pref, nil
}

func writePreferencesUTF8(filename string, pref UserPreferences) error {
	data, err := json.MarshalIndent(pref, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	return err
}

func readPreferencesASCII(filename string) (*UserPreferences, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Convert ASCII (with escapes) back to UTF-8
	b := make([]byte, len(data))
	copy(b, data)
	decodedData, err := json.Unmarshal(b, &UserPreferences{})
	if err != nil {
		return nil, err
	}

	var pref UserPreferences
	err = json.Unmarshal(data, &pref)
	if err != nil {
		return nil, err
	}

	return &pref, nil
}

func writePreferencesASCII(filename string, pref UserPreferences) error {
	data, err := json.MarshalIndent(pref, "", "    ")
	if err != nil {
		return err
	}

	// Encode JSON string to ASCII (with escapes)
	newData := make([]byte, len(data))
	for i, c := range data {
		if !strings.IsPrint(string(c)) {
			newData[i] = '\\'
			i++
			newData[i] = 'x'
			i++
			newData[i] = byte(fmt.Sprintf("%02x", int(c))[0])
			i++
			newData[i] = byte(fmt.Sprintf("%02x", int(c))[1])
		} else {
			newData[i] = c
		}
	}

	err = ioutil.WriteFile(filename, newData, 0644)
	return err
}

func main() {
	filename := "user_preferences.json"
	pref := UserPreferences{
		Username: "johndoe",
		Theme:    "dark",
		Language: "español",
	}

	// Write preferences in UTF-8
	if err := writePreferencesUTF8(filename, pref); err != nil {
		fmt.Println("Error writing preferences in UTF-8:", err)
		return
	}

	// Read preferences in UTF-8
	readPrefUTF8, err := readPreferencesUTF8(filename)
	if err != nil {
		fmt.Println("Error reading preferences in UTF-8:", err)
		return
	}
	fmt.Println("Read Preferences (UTF-8):")
	fmt.Printf("Username: %s\n", readPrefUTF8.Username)
	fmt.Printf("Theme: %s\n", readPrefUTF8.Theme)
	fmt.Printf("Language: %s\n", readPrefUTF8.Language)

	// Write preferences in ASCII (with escapes)
	if err := writePreferencesASCII(filename, pref); err != nil {
		fmt.Println("Error writing preferences in ASCII:", err)
		return
	}

	// Read preferences in ASCII
	readPrefASCII, err := readPreferencesASCII(filename)
	if err != nil {
		fmt.Println("Error reading preferences in ASCII:", err)
		return
	}
	fmt.Println("\nRead Preferences (ASCII):")
	fmt.Printf("Username: %s\n", readPrefASCII.Username)
	fmt.Printf("Theme: %s\n", readPrefASCII.Theme)
	fmt.Printf("Language: %s\n", readPrefASCII.Language)
}
