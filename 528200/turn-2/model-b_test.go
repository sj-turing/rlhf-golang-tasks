// filename: main_test.go
package main

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

const (
	numCities      = 10000 // Generate a large dataset for performance testing
	numIterations  = 10    // Execute a few iterations for better statistics
	maxTemperature = 100   // Maximum random temperature
)

func BenchmarkCalculateAverageTemperature(b *testing.B) {
	// Initialize a random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a fixed number of weather data for performance testing
	weatherData := make(chan WeatherData, numCities)
	for i := 0; i < numCities; i++ {
		city := cities[rand.Intn(len(cities))]
		weatherData <- WeatherData{
			City:        city,
			Country:     cityMappedToCountry[city],
			Temperature: float32(rand.Intn(maxTemperature)),
		}
	}
	close(weatherData)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		// Compute the average temperature for each country using a fixed weather data
		calculateAverageTemperature(ctx, weatherData)
	}
}
