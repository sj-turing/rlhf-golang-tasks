package main

import (
	"fmt"
	"net/url"

	"github.com/goware/urlx"
)

func refactorParseURL(input string) (*url.URL, error) {
	normalizedURL, err := urlx.NormalizeString(input)
	if err != nil {
		return nil, err
	}

	return urlx.Parse(normalizedURL)
}

// func legacyParseURL(input string) (map[string]interface{}, error) {
func legacyParseURL(input string) (*url.URL, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	return u, nil

}

func main() {
	input := "localhost:80///x///y/z/../././index.html?b=y&a=x#t=20"

	legarcyURL, err := legacyParseURL(input)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Printf("Using Legacy Code:   %+v\n", legarcyURL)

	refactorURL, err := refactorParseURL(input)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Printf("Using Refactor Code: %+v\n", refactorURL)
}
