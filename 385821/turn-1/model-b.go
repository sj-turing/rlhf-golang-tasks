package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// logErrorWithContext logs an error message along with relevant context data
func logErrorWithContext(ctx context.Context, err error) {
	// Create a new logger with the function name and line number
	l := log.New(os.Stderr, fmt.Sprintf("%s:%d: ", funcName(2), getLine()), log.LstdFlags)

	// Extract values from the context
	requestID, _ := ctx.Value("request_id").(string)
	userID, _ := ctx.Value("user_id").(string)

	// Build the error message with context data
	var errorMsg string
	if err != nil {
		errorMsg = fmt.Sprintf("Request ID: %s, User ID: %s, Error: %s", requestID, userID, err.Error())
	} else {
		errorMsg = fmt.Sprintf("Request ID: %s, User ID: %s, Unknown error occurred.", requestID, userID)
	}

	// Log the error message
	l.Println(errorMsg)
}

// funcName returns the name of the calling function
func funcName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "???"
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	i := strings.LastIndex(name, "/")
	if i >= 0 {
		name = name[i+1:]
	}
	i = strings.Index(name, "(")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// getLine returns the line number of the calling function
func getLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

func main() {
	// Create a new context with request ID and user ID
	ctx := context.WithValue(context.Background(), "request_id", "req-123")
	ctx = context.WithValue(ctx, "user_id", "user456")

	// Simulate an error
	err := someFunction(ctx)
	if err != nil {
		logErrorWithContext(ctx, err)
	}
}

func someFunction(ctx context.Context) error {
	// Some code here that might return an error
	time.Sleep(2 * time.Second)
	return fmt.Errorf("some function failed")
}
