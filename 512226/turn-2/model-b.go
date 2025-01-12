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

// UserInterest represents a user's interest in specific products during Black Friday.
type UserInterest struct {
	UserID     string
	SessionID  string
	EventTime  time.Time
	ProductID  string
	CategoryID string
	Price      float64
}

// UserInterestManager manages the storage and analysis of user interests.
type UserInterestManager struct {
	interests []*UserInterest
	mutex     sync.Mutex
}

// NewUserInterestManager creates a new UserInterestManager instance.
func NewUserInterestManager() *UserInterestManager {
	return &UserInterestManager{
		interests: make([]*UserInterest, 0),
	}
}

// TrackUserInterest records a user's interest in a mobile phone product during Black Friday.
func (uim *UserInterestManager) TrackUserInterest(userID string, requestURL string) {
	uim.mutex.Lock()
	defer uim.mutex.Unlock()

	u, err := url.Parse(requestURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	queryParams := u.Query()
	productID := queryParams.Get("productID")
	categoryID := queryParams.Get("categoryID")
	priceStr := queryParams.Get("price")

	if productID == "" || categoryID == "" || priceStr == "" {
		log.Printf("Missing required query parameters: productID, categoryID, or price")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Printf("Error parsing price: %v", err)
		return
	}

	userInterest := &UserInterest{
		UserID:     userID,
		SessionID:  uuid.NewString(),
		EventTime:  time.Now(),
		ProductID:  productID,
		CategoryID: categoryID,
		Price:      price,
	}

	uim.interests = append(uim.interests, userInterest)
}

// AnalyzeUserInterests analyzes user interests and identifies patterns during Black Friday.
func (uim *UserInterestManager) AnalyzeUserInterests() {
	uim.mutex.RLock()
	defer uim.mutex.RUnlock()

	fmt.Printf("-- User Interests During Black Friday --\n")
	for _, interest := range uim.interests {
		fmt.Printf("UserID: %s, SessionID: %s, ProductID: %s, CategoryID: %s, Price: %.2f\n",
			interest.UserID, interest.SessionID, interest.ProductID, interest.CategoryID, interest.Price)
	}
	fmt.Printf("--\n\n")

	// Insert your code here to analyze user interests, such as finding popular products or segmenting users based on interests.
}

// Middleware for tracking user interests.
func UserInterestMiddleware(uim *UserInterestManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("user-id") // Replace with actual user ID logic
		uim.TrackUserInterest(userID, r.URL.String())
	})
}

func main() {
	uim := NewUserInterestManager()

	// Register middleware
	http.Handle("/", UserInterestMiddleware(uim))

	// Start a goroutine to periodically analyze user interests
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			uim.AnalyzeUserInterests()
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
