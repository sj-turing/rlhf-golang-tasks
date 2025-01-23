package main

import (
	"context"
	"testing"
	"time"
)

func TestGetWeatherData(t *testing.T) {
	t.Parallel()

	timeoutSeconds := 1
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	weatherDataChan := getWeatherData(ctx)

forloop:
	for {
		select {
		case data, ok := <-weatherDataChan:
			if ok {
				if data.Country == "" || data.City == "" {
					t.Fatalf("got empty data")
				}
			} else {
				break forloop
			}
		case <-time.After(time.Duration(timeoutSeconds+1) * time.Second):
			t.Fatal("got timeout while calling the getWeatherData")
		}
	}
	cancel()
}

func BenchmarkGetWeatherData(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for i := 0; i < b.N; i++ {
		select {
		case <-getWeatherData(ctx):
		case <-time.After(3 * time.Second):
			b.Fatal("getWeatherData is taking more than 2 second to generate data")
		}
	}
}
