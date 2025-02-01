package main

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logData := map[string]interface{}{
			"timestamp":   time.Now().Format(time.RFC3339),
			"request_ip":  r.RemoteAddr,
			"http_method": r.Method,
			"url":         r.RequestURI,
			"user_agent":  r.UserAgent(),
		}

		// Log query parameters
		logData["query_parameters"] = r.URL.Query()

		logAuditTrail(logData) // Function to handle the actual logging

		next.ServeHTTP(w, r)
	})
}

func trackTransactionLifecycle(r *http.Request, transactionStatus string) {
	transactionDetails := map[string]string{
		"userID": getUserIDFromSession(r),
		"status": transactionStatus,
	}
	if transactionStatus == "initiated" {
		// Track transaction start
	} else if transactionStatus == "failed" {
		// Track failed transaction completion
	}
	logAuditTrail(transactionDetails)
}

func trackSubscriptionBrowsing(r *http.Request) {
	browsingDetails := map[string]string{
		"userID":   getUserIDFromSession(r),
		"activity": "browsing_subscription",
	}
	logAuditTrail(browsingDetails)
}
