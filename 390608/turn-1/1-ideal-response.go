package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{
		Transport: &retryRoundTripper{
			defaultRoundTripper: http.DefaultTransport,
			maxRetries:          5,
			delay:               1,
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9000", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("statusCode: ", res.StatusCode)
	log.Println("response Body: ", string(body))
}

type retryRoundTripper struct {
	defaultRoundTripper http.RoundTripper
	maxRetries          int
	delay               int
}

func (rr retryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var attempts int
	var backOff time.Duration
	for {
		res, err := rr.defaultRoundTripper.RoundTrip(r)
		attempts++

		log.Printf("%d attempt\n", attempts)

		// max retries exceeded
		if attempts == rr.maxRetries {
			return res, err
		}

		// validating successful response
		if err == nil && res.StatusCode < http.StatusInternalServerError {
			return res, err
		}

		backOff = time.Duration(rr.delay) * time.Second
		rr.delay *= 2 // backingOff by power of 2

		// delay and retry
		select {
		case <-r.Context().Done():
			return res, r.Context().Err()
		case <-time.After(backOff):
		}
	}
}
