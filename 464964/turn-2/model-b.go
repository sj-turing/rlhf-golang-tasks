package main

import (
	"fmt"
	"github.com/goware/urlx"
)

func main() {
	inputURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2#fragment"

	// Parse the URL using urlx.Parse()
	u, err := urlx.Parse(inputURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)

	// Extract query parameters using urlx.Values()
	query := u.Values()
	fmt.Println("Query Parameters:")
	for key, values := range query {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	fmt.Println("Fragment:", u.Fragment)
}
