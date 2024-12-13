package main

/*

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func logContextError(ctx context.Context, err error) {
	// extract the user-defined values from the context
	userID, ok := ctx.Value("userID").(int)
	method, ok := ctx.Value("method").(string)

	if !ok {
		log.Printf("Error retrieving context data: %v", err)
	}
	// Custom log format with context and error message
	log.Printf("Request: User ID: %d, Method: %s, Error: %v\n", userID, method, err)
}

func sampleFunctionWithContext(ctx context.Context) error {
	// Simulate an error
	return errors.New("some internal error occurred")
}

func main() {
	// Sample context creation
	userID := 42
	method := "GET"

	ctx := context.Background() // Starting with a basic context
	ctx = context.WithValue(ctx, "userID", userID)
	ctx = context.WithValue(ctx, "method", method)

	// Call the function with the context and check for errors
	err := sampleFunctionWithContext(ctx)
	if err != nil {
		logContextError(ctx, err)
	}
}

*/
