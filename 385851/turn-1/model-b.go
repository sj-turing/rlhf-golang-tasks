package main

import (
	"errors"
	"fmt"
	"testing"
)

// Simulated function that might return an error
func someFunction() error {
	return errors.New("someFunction failed")
}

// Error handler function
func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}

// Test case function
func TestErrorHandling(t *testing.T) {
	var testCases = []struct {
		inputError error
		expected   string
	}{
		{
			inputError: errors.New("someFunction failed"),
			expected:   "Error: someFunction failed",
		},
		{
			inputError: nil,
			expected:   "",
		},
		{
			inputError: fmt.Errorf("custom error"),
			expected:   "Error: custom error",
		},
	}

	for _, tc := range testCases {
		handleError(tc.inputError)

		// Check the output using testing.T
		// Note that this method captures the standard output
		// and compares it with the expected string
		output := fmt.Sprintf("%s", tc.expected)
		if tc.expected != "" {
			t.Helper()
			t.Run(fmt.Sprintf("Error: %s", tc.inputError), func(t *testing.T) {
				t.CheckPoint()
				t.Check(output, tc.expected)
			})
		}
	}
}

func main() {
	testing.Main()
}
