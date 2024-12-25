package main

import (
	"fmt"
	"github.com/mattn/urlquery"
	"net/url"
)

func refactoredParseURL(input string) (*url.URL, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	// Use urlquery for enhanced query parameter handling
	q := urlquery.Query{}
	for _, pair := range u.Query() {
		for _, value := range pair {
			q[pair[0]] = append(q[pair[0]], value)
		}
	}

	// Modify query parameters if needed (this is just an example)
	q["key1"] = append(q["key1"], "updated_value1")

	// Update URL with the modified query
	u.RawQuery = q.Encode()

	return u, nil
}

func main() {
	input := "https://example.com/path?key1=value1&key2=value2"
	parsedURL, err := refactoredParseURL(input)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println(parsedURL)
}
