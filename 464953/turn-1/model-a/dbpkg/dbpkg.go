package dbpkg

import (
	"database/sql"
	"encoding/url"
	"fmt"
)

// StoreQueryParams stores the query parameters in a database
func StoreQueryParams(db *sql.DB, params map[string]string) error {
	encodedParams, err := url.Values(params).Encode()
	if err != nil {
		return fmt.Errorf("error encoding parameters: %v", err)
	}

	query := fmt.Sprintf(`
		INSERT INTO query_params (params)
		VALUES ('$%s$')
	`, encodedParams)

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}

// ParseAndStoreQueryParams parses a URL and stores its query parameters in a database
func ParseAndStoreQueryParams(db *sql.DB, urlString string) error {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("error parsing URL: %v", err)
	}

	return StoreQueryParams(db, parsedURL.Query())
}
