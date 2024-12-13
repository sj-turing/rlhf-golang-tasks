package main

/*

import (
	"errors"
	"fmt"
	"testing"
)

var (
	errInvalidInput = errors.New("invalid input")
	errInternal     = errors.New("internal error")
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errInvalidInput
	}
	return x / y, nil
}

func handleDivideError(x, y float64) error {
	result, err := divide(x, y)
	if err != nil {
		return fmt.Errorf("failed to divide: %w", err)
	}
	return nil
}

func TestHandleDivideError(t *testing.T) {
	type testCase struct {
		x, y float64
		err  error
	}
	testCases := []testCase{
		{x: 10, y: 2, err: nil},
		{x: 10, y: 0, err: errors.New("failed to divide: invalid input")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("x=%.1f, y=%.1f", tc.x, tc.y), func(t *testing.T) {
			if got := handleDivideError(tc.x, tc.y); !errors.Is(got, tc.err) {
				t.Errorf("got %v, want %v", got, tc.err)
			}
		})
	}
}

*/
