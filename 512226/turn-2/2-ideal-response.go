package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MobilePhoneInterest struct {
	UserID        string
	SessionID     string
	LastEventTime time.Time
	PageViews     []string
	CountOfVisits map[string]map[string]map[int]uint16
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
			CountOfVisits: map[string]map[string]map[int]uint16{},
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
	brand := queryParams["brand"][0]
	model := queryParams["model"][0]
	budget, _ := strconv.Atoi(queryParams["budget"][0])

	if count, ok := interest.CountOfVisits[brand][model][budget]; ok {
		interest.CountOfVisits[brand][model][budget] = count + 1
	} else {
		if interest.CountOfVisits[brand] == nil {
			interest.CountOfVisits[brand] = map[string]map[int]uint16{}
		}
		if interest.CountOfVisits[brand][model] == nil {
			interest.CountOfVisits[brand][model] = map[int]uint16{}
		}
		interest.CountOfVisits[brand][model][budget] = 1
	}

	mim.interests[userID] = interest
}

func (mim *MobilePhoneInterestManager) AnalyzeInterests() {
	mim.mutex.RLock()
	defer mim.mutex.RUnlock()

	for _, interest := range mim.interests {
		for brand, modelBudget := range interest.CountOfVisits {
			for model, budget := range modelBudget {
				for b, count := range budget {
					fmt.Printf("'%s' user visited '%s' brand's '%s' model in budget of %d with %d times\n", interest.UserID, brand, model, b, count)
				}
			}
		}

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
