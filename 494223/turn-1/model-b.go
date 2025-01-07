package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// UserPref represents the user preferences structure.
type UserPref struct {
	Username string `json:"username"`
	Language string `json:"language"`
}

// WritePreferencesUTF8 writes user preferences to a file in UTF-8 encoding.
func WritePreferencesUTF8(filePath string, pref UserPref) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(pref); err != nil {
		return err
	}
	return w.Flush()
}

// ReadPreferencesUTF8 reads user preferences from a file in UTF-8 encoding.
func ReadPreferencesUTF8(filePath string) (UserPref, error) {
	var pref UserPref
	file, err := os.Open(filePath)
	if err != nil {
		return pref, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&pref); err != nil {
		return pref, err
	}
	return pref, nil
}


// WritePreferencesASCII writes user preferences to a file in ASCII encoding.
func WritePreferencesASCII(filePath string, pref UserPref) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    // Create an encoder for UTF-8 to ASCII conversion.
    utf8ToASCII := transform.NewReader(strings.NewReader(string(pref)), charmap.Windows1252.NewEncoder())
    r := bufio.NewReader(utf8ToASCII)
    
    encoder := json.NewEncoder(w)
    encoder.SetIndent("", "  ")
    var prefASCII UserPref
    decode := json.NewDecoder(r)
    if err := decode.Decode(&prefASCII); err != nil {
        return err
    }
    // Convert the non-ASCII characters to '?'
    prefASCII.Username = strings.Map(func(r rune) rune {
        if r < 128 {
            return r
        }
        return '?'
    }, prefASCII.Username)
    prefASCII.Language = strings.Map(func(r rune) rune {
        if r < 128 {
            return r
        }
        return '?'
    }, prefASCII.Language)

    if err := encoder.Encode(prefASCII); err != nil {
        return err
    }
    return w.Flush()
}

// ReadPreferencesASCII reads user preferences from a file in ASCII encoding.
func ReadPreferencesASCII(filePath string) (UserPref, error) {
    var pref UserPref
    file, err := os.Open(filePath)
    if err != nil {
        return pref, err
    }
    defer file.Close()
