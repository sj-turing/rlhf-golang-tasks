// filename: main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
)

// UserJourney represents a single user's journey through the e-commerce site.
type UserJourney struct {
	UserID              string                       // Unique identifier for the user
	SessionID           string                       // Unique identifier for the current session
	LastEventTime       time.Time                    // Timestamp of the last activity
	PageViews           []string                     // List of URLs visited
	CountOfVisitedPages map[string]map[string]uint32 // Map of parameters and their values
}

// UserJourneyManager manages the storage and analysis of user journeys.
type UserJourneyManager struct {
	journeys map[string]*UserJourney // Map of UserID to UserJourney
	mutex    sync.RWMutex
}

// NewUserJourneyManager creates a new UserJourneyManager instance.
func NewUserJourneyManager() *UserJourneyManager {
	return &UserJourneyManager{
		journeys: make(map[string]*UserJourney),
	}
}

// TrackUserJourney records a new page view for the given user.
func (ujm *UserJourneyManager) TrackUserJourney(userID string, requestUrl string) {
	ujm.mutex.Lock()
	defer ujm.mutex.Unlock()

	userJourney, ok := ujm.journeys[userID]
	if !ok {
		userJourney = &UserJourney{
			UserID:              userID,
			SessionID:           uuid.NewString(),
			LastEventTime:       time.Now(),
			PageViews:           []string{},
			CountOfVisitedPages: map[string]map[string]uint32{},
		}
	}

	userJourney.PageViews = append(userJourney.PageViews, requestUrl)
	userJourney.LastEventTime = time.Now()

	// Extract and store URL parameters
	u, err := url.Parse(requestUrl)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	// http://localhost:8080?category=kids&productName=t-shirts

	queryParams := u.Query()
	category := queryParams["category"][0]
	productName := queryParams["productName"][0]

	if count, ok := userJourney.CountOfVisitedPages[category][productName]; ok {
		userJourney.CountOfVisitedPages[category][productName] = count + 1
	} else {
		if userJourney.CountOfVisitedPages[category] == nil {
			userJourney.CountOfVisitedPages[category] = map[string]uint32{}
		}
		userJourney.CountOfVisitedPages[category][productName] = 1
	}

	// assumming there are only two query params as category and productName

	ujm.journeys[userID] = userJourney
}

// AnalyzeUserJourneys analyzes user journeys and identifies patterns.
func (ujm *UserJourneyManager) AnalyzeUserJourneys() {
	ujm.mutex.RLock()
	defer ujm.mutex.RUnlock()

	for _, userJourney := range ujm.journeys {
		for category, productNames := range userJourney.CountOfVisitedPages {
			for productName, count := range productNames {
				fmt.Printf("'%s' User visited '%s' category with '%s' productName for '%d' times\n", userJourney.UserID, category, productName, count)
			}
		}
	}

	fmt.Printf("--\n\n")

	// Insert your code here to forward user journey data to external tracking platforms,
	// providing a graphical representation of user behavior patterns
	// you can track user journey session-wise as well with SessionID
}

// Middleware for tracking user journeys.
func UserJourneyMiddleware(ujm *UserJourneyManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate or retrieve user ID (e.g., from cookies or authentication)
		userID := r.Header.Get("user-id") // Replace with actual user ID logic

		ujm.TrackUserJourney(userID, r.URL.String())
	})
}

func main() {
	ujm := NewUserJourneyManager()

	// Register middleware
	http.Handle("/", UserJourneyMiddleware(ujm))

	// Start a goroutine to periodically analyze user journeys
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			ujm.AnalyzeUserJourneys()
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
