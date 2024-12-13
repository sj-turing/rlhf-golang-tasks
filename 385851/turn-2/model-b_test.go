package main

import (
	"errors"
	"fmt"
	"testing"
)

// Custom error type
type customError struct {
	message string
}

func (e customError) Error() string {
	return e.message
}

// Function with error handler
func processData(input string) error {
	if input == "" {
		return customError{"input cannot be empty"}
	}
	return nil
}

func handleError(err error) error {
	if err != nil {
		return fmt.Errorf("error processing: %w", err)
	}
	return nil
}

func TestErrorHandler(t *testing.T) {

	type testCase struct {
		input string
		want  error
	}

	tests := []testCase{
		{
			input: "",
			want:  fmt.Errorf("error processing: %w", customError{"input cannot be empty"}),
		},
		{
			input: "hello",
			want:  nil,
		},
	}

	for index, tt := range tests {
		t.Run(tt.input+fmt.Sprintf("%d", index), func(t *testing.T) {
			got := handleError(processData(tt.input))
			if !errors.Is(tt.want, got) {
				t.Errorf("handleError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorHandlerWithWrappedError(t *testing.T) {
	err := handleError(processData(""))

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	underlyingErr := errors.Unwrap(err)
	if underlyingErr == nil {
		t.Fatalf("expected underlying error, got nil")
	}

	customErr, ok := underlyingErr.(customError)
	if !ok {
		t.Fatalf("underlying error should be of type customError, got %T", underlyingErr)
	}

	if customErr.message != "input cannot be empty" {
		t.Errorf("underlying error message incorrect: got %q, want %q", customErr.message, "input cannot be empty")
	}
}
