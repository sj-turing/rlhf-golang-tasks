package main

/*

import (
	"fmt"
	"net/url"
)

func legacyParseURL(input string) (*url.URL, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	// Manually handle query parameters
	rawQuery := u.Query().Encode()
	components := rawQuery.Split("&")
	for _, component := range components {
		key, value := component
		key = url.QueryEscape(key)
		value = url.QueryEscape(value)
		u.RawQuery += fmt.Sprintf("&%s=%s", key, value)
	}

	return u, nil
}

func main() {
	input := "https://example.com/path?key1=value1&key2=value2"
	parsedURL, err := legacyParseURL(input)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println(parsedURL)
}
*/
