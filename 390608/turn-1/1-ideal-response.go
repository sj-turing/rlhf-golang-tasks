package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	// create a custom http client with retryRoundTripper
	c := &http.Client{
		Transport: &retryRoundTripper{
			defaultRoundTripper: http.DefaultTransport, // default http client
			maxRetries:          5,                     // number of retries
			delay:               1,                     // used for exponential backoff
		},
	}

	// creating a new request to call client
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

	// reading resposne body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// printing logs
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
		// calling the http client
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

		backOff = time.Duration(rr.delay) * time.Second // assuming delay in second
		rr.delay *= 2                                   // backingOff by power of 2

		// delay and retry
		select {
		case <-r.Context().Done():
			return res, r.Context().Err()
		case <-time.After(backOff): // waiting till backOff duration
		}
	}
}
