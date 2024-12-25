package main

/*

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

// LegacyParseURL parses a URL string manually
func LegacyParseURL(rawURL string) *url.URL {
	// Manual parsing logic goes here
	scheme := strings.Split(rawURL, "://")[0]
	host := strings.Split(strings.Split(rawURL, "://")[1], "/")[0]
	path := strings.Replace(rawURL, scheme+"://"+host, "", -1)

	u := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	return u
}

// NewParseURL parses a URL string using the net/url package
func NewParseURL(rawURL string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func main() {
	// Example usage of legacy parsing code
	rawURL := "http://example.com/path/to/resource"
	legacyURL := LegacyParseURL(rawURL)
	fmt.Println("Legacy URL:", legacyURL)

	// Example usage of new parsing code
	newURL, err := NewParseURL(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New URL:", newURL)
}
*/
