package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserPreferences struct {
	Username string `json:"username"`
	Theme    string `json:"theme"`
	Language string `json:"language"`
	Note     string `json:"note"` // Added a multilingual field
}

func readPreferences(filename, encoding string) (*UserPreferences, error) {
	var pref UserPreferences
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	switch encoding {
	case "utf-8":
		err = json.Unmarshal(data, &pref)
	case "ascii":
		// Convert to UTF-8 to handle Unicode data (in field Note)
		utf8Data := make([]byte, len(data))
		for i, b := range data {
			utf8Data[i] = b & 0x7F
		}
		err = json.Unmarshal(utf8Data, &pref)
	default:
		return nil, fmt.Errorf("unsupported encoding: %s", encoding)
	}
	return &pref, err
}

func writePreferences(filename, encoding string, pref UserPreferences) error {
	var data []byte
	var err error
	switch encoding {
	case "utf-8":
		data, err = json.MarshalIndent(pref, "", "    ")
	case "ascii":
		// Convert to ASCII by stripping Unicode characters
		asciiPref := UserPreferences{
			Username: pref.Username,
			Theme:    pref.Theme,
			Language: pref.Language,
			Note:     removeNonASCII(pref.Note),
		}
		data, err = json.MarshalIndent(asciiPref, "", "    ")
	default:
		return fmt.Errorf("unsupported encoding: %s", encoding)
	}
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Helper function to remove non-ASCII characters from a string
func removeNonASCII(s string) string {
	var r []rune
	for _, c := range s {
		if c < 128 {
			r = append(r, c)
		}
	}
	return string(r)
}

func main() {
	filename := "user_preferences.json"

	prefUTF8 := UserPreferences{
		Username: "johndoe",
		Theme:    "dark",
		Language: "español",
		Note:     "Hello, 世界! This is a multilingual note.",
	}

	fmt.Println("Writing UTF-8 preferences...")
	err := writePreferences(filename, "utf-8", prefUTF8)
	if err != nil {
		fmt.Println("Error writing UTF-8 preferences:", err)
		return
	}

	fmt.Println("Reading UTF-8 preferences...")
	readPrefUTF8, err := readPreferences(filename, "utf-8")
	if err != nil {
		fmt.Println("Error reading UTF-8 preferences:", err)
		return
	}
	fmt.Println(readPrefUTF8)

	fmt.Println("\nWriting ASCII preferences...")
	prefASCII := prefUTF8
	err = writePreferences(filename, "ascii", prefASCII)
	if err != nil {
		fmt.Println("Error writing ASCII preferences:", err)
		return
	}

	fmt.Println("Reading ASCII preferences...")
	readPrefASCII, err := readPreferences(filename, "ascii")
	if err != nil {
		fmt.Println("Error reading ASCII preferences:", err)
		return
	}
	fmt.Println(readPrefASCII)
}
