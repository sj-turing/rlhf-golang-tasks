package main

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/go-gax/v2/gax"
	"github.com/google/go-analytics/ga"
	"net/url"
	"sync"
)

const (
	propertyID   = "YOUR_GA4_PROPERTY_ID" // Replace with your GA4 property ID
	userID       = "YOUR_USER_ID"         // Replace with your user ID
	dimensionKey = "dimension1"           // Replace with your user dimension key
)

var (
	gAnalyticsClient *ga.Client
	gA4Client        *ga.BetaAnalyticsClient
	mu               sync.Mutex
)

// Initiate the Google Analytics 4 client
func initGA4Client(ctx context.Context) error {
	// Initialize Google Analytics 4 Client
	var err error
	gA4Client, err = ga.NewBetaAnalyticsClient(ctx)
	if err != nil {
		return fmt.Errorf("creating GA4 client: %v", err)
	}
	return nil
}

// Get the Google Analytics 4 Client
func getGA4Client() (*ga.BetaAnalyticsClient, error) {
	mu.Lock()
	defer mu.Unlock()
	if gA4Client == nil {
		ctx := context.Background()
		if err := initGA4Client(ctx); err != nil {
			return nil, err
		}
	}
	return gA4Client, nil
}

// Track the User Interaction with Google Analytics 4
func trackUserInteraction(dimensionValues map[string]string) error {
	// Get the GA4 Client
	client, err := getGA4Client()
	if err != nil {
		return fmt.Errorf("tracking user interaction: %v", err)
	}

	// Prepare the User Data
	userData := &ga.CreateBetaUserRequest{
		Property: fmt.Sprintf("properties/%s", propertyID),
		User:     &ga.User{},
	}

	// Set the User ID and Custom Dimensions
	userData.User.UserId = userID
	userData.User.UserProperties = map[string]*ga.UserProperty{}
	for key, value := range dimensionValues {
		userData.User.UserProperties[key] = &ga.UserProperty{
			Value: value,
			Parameters: map[string]*ga.UserPropertyParameters{
				"scope": {Value: "USER"},
			},
		}
	}

	// Send User Data
	_, err = client.CreateBetaUser(context.Background(), userData)
	if err != nil {
		return fmt.Errorf("tracking user interaction: %v", err)
	}

	return nil
}

// Initialize the Google Analytics client
func initGoogleAnalyticsClient() error {
	var err error
	gAnalyticsClient, err = ga.NewClient(nil, propertyID)
	if err != nil {
		return fmt.Errorf("creating GA client: %w", err)
	}
	return nil
}

// Get the Google Analytics client
func getGoogleAnalyticsClient() (*ga.Client, error) {
	mu.Lock()
	defer mu.Unlock()
	if gAnalyticsClient == nil {
		if err := initGoogleAnalyticsClient(); err != nil {
			return nil, err
		}
	}
	return gAnalyticsClient, nil
}

// Send a page view hit to Google Analytics
func sendPageViewHit(ctx context.Context, pagePath, pageTitle string) error {
	client, err := getGoogleAnalyticsClient()
	if err != nil {
		return fmt.Errorf("sending pageview hit: %w", err)
	}

	v := client.NewPageview(ctx, pagePath, pageTitle)
	return client.Send(ctx, v)
}

// Send an event hit to Google Analytics
func sendEventHit(ctx context.Context, eventCategory, eventAction, eventLabel string) error {
	client, err := getGoogleAnalyticsClient()
	if err != nil {
		return fmt.Errorf("sending event hit: %w", err)
	}

	v := client.NewEvent(ctx, eventCategory, eventAction, eventLabel)
	return client.Send(ctx, v)
}
