package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MobilePhoneInterest struct {
	UserID        string
	SessionID     string
	LastEventTime time.Time
	PageViews     []string
	Brand         string
	Model         string
	Budget        int
	CountOfVisits uint32
}

type MobilePhoneInterestManager struct {
	interests map[string]*MobilePhoneInterest
	mutex     sync.RWMutex
}

func NewMobilePhoneInterestManager() *MobilePhoneInterestManager {
	return &MobilePhoneInterestManager{
		interests: make(map[string]*MobilePhoneInterest),
	}
}

func (mim *MobilePhoneInterestManager) TrackInterest(userID, sessionID string, requestUrl string) {
	mim.mutex.Lock()
	defer mim.mutex.Unlock()

	interest, ok := mim.interests[userID]
	if !ok {
		interest = &MobilePhoneInterest{
			UserID:        userID,
			SessionID:     sessionID,
			LastEventTime: time.Now(),
			PageViews:     []string{},
			CountOfVisits: 0,
		}
	}

	interest.PageViews = append(interest.PageViews, requestUrl)
	interest.LastEventTime = time.Now()

	// Extract and store URL parameters
	u, err := url.Parse(requestUrl)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	queryParams := u.Query()
	interest.Brand = queryParams["brand"][0]
	interest.Model = queryParams["model"][0]
	interest.Budget, _ = strconv.Atoi(queryParams["budget"][0])

	interest.CountOfVisits++

	mim.interests[userID] = interest
}

func (mim *MobilePhoneInterestManager) AnalyzeInterests() {
	mim.mutex.RLock()
	defer mim.mutex.RUnlock()

	for _, interest := range mim.interests {
		fmt.Printf("UserID: %s, Brand: %s, Model: %s, Budget: %d, Visits: %d\n", interest.UserID, interest.Brand, interest.Model, interest.Budget, interest.CountOfVisits)
	}

	fmt.Printf("--\n\n")
}

func UserInterestMiddleware(mim *MobilePhoneInterestManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("user-id")
		if userID == "" {
			http.Error(w, "User ID not provided", http.StatusBadRequest)
			return
		}

		sessionID := r.Header.Get("session-id")
		if sessionID == "" {
			sessionID = uuid.NewString()
			w.Header().Set("session-id", sessionID)
		}

		mim.TrackInterest(userID, sessionID, r.URL.String())
	})
}

func main() {
	mim := NewMobilePhoneInterestManager()

	// Register middleware
	http.Handle("/", UserInterestMiddleware(mim))

	// Start a goroutine to periodically analyze user interests
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			mim.AnalyzeInterests()
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
