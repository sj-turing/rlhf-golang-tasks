package main

/*

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// User represents a user in our application.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserProcessor defines the operations that can be performed on a user.
type UserProcessor interface {
	Process(*User) error
}

// Option represents a functional option for the user processor.
type Option func(p *UserProcessor)

// NewUserProcessor creates a new UserProcessor with default settings.
func NewUserProcessor(opts ...Option) *UserProcessor {
	p := &UserProcessorImpl{
		processTimeout: 5 * time.Second, // Default timeout
	}

	// Apply functional options
	for _, opt := range opts {
		opt(p)
	}

	return p
}

// UserProcessorImpl is the concrete implementation of UserProcessor.
type UserProcessorImpl struct {
	processTimeout time.Duration // Timeout for processing user data
}

// WithProcessTimeout sets the timeout for processing user data.
func WithProcessTimeout(t time.Duration) Option {
	return func(p *UserProcessor) {
		p.(*UserProcessorImpl).processTimeout = t
	}
}

// Process implements the UserProcessor interface.
func (p *UserProcessorImpl) Process(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), p.processTimeout)
	defer cancel()

	// Simulate a processing operation
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(1 * time.Second):
		fmt.Printf("Processed user %d successfully\n", user.ID)
		return nil
	}
}

// UserHandler handles user data processing.
type UserHandler struct {
	users     []*User
	wg        sync.WaitGroup
	processor UserProcessor
}

// NewUserHandler creates a new UserHandler with the specified UserProcessor.
func NewUserHandler(processor UserProcessor) *UserHandler {
	return &UserHandler{
		processor: processor,
	}
}

// AddUser adds a user to the handler for processing.
func (h *UserHandler) AddUser(user *User) {
	h.users = append(h.users, user)
	h.wg.Add(1)

	go h.processUser(user)
}

// processUser processes a user in a Goroutine.
func (h *UserHandler) processUser(user *User) {
	if err := h.processor.Process(user); err != nil {
		log.Printf("Error processing user %d: %v\n", user.ID, err)
	}
	h.wg.Done()
}

// Wait waits for all users to be processed.
func (h *UserHandler) Wait() {
	h.wg.Wait()
}

func main() {
	// Create a user processor with a custom timeout
	userProcessor := NewUserProcessor(WithProcessTimeout(10 * time.Second))

	// Create a user handler
	userHandler := NewUserHandler(userProcessor)

	// Add users to be processed
	users := []*User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}

	for _, user := range users {
		userHandler.AddUser(user)
	}

	// Wait for all users to be processed
	userHandler.Wait()
}
*/
