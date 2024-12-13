package main

/*

import (
	"context"
	"errors"
	"fmt"
)

// Function to perform some service operation
func serviceOperation(ctx context.Context) error {
	return errors.New("Service operation failed")
}

// Function to handle HTTP request
func handleRequest(ctx context.Context) error {
	// Create a new context with added user ID value
	userID := "user123"
	ctxWithUserID := context.WithValue(ctx, "userID", userID)

	// Perform the service operation with the updated context
	err := serviceOperation(ctxWithUserID)
	if err != nil {
		return logErrorWithContext(ctxWithUserID, err)
	}
	return nil
}

// Function to log errors with context information
func logErrorWithContext(ctx context.Context, err error) error {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		userID = "Unknown User"
	}

	log := fmt.Sprintf("Error occurred for User: %s, Error: %v", userID, err)
	fmt.Println(log)
	return err
}

func main() {
	// Create the initial context
	ctx := context.Background()

	// Handle the request with the initial context
	err := handleRequest(ctx)
	if err != nil {
		fmt.Println("Application error:", err)
	}
}
*/
