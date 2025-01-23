package main

import (
	"math/rand"
	"testing"
)

const (
	numCities      = 10000 // Generate a large dataset for performance testing
	maxTemperature = 100   // Maximum random temperature
)

func simulateWeatherData() <-chan WeatherData {

	// Generate a fixed number of weather data for performance testing
	weatherDataChan := make(chan WeatherData, numCities)
	for i := 0; i < numCities; i++ {
		city := cities[rand.Intn(len(cities))]
		weatherDataChan <- WeatherData{
			City:        city,
			Country:     cityMappedToCountry[city],
			Temperature: float32(rand.Intn(maxTemperature)),
		}
	}
	close(weatherDataChan)
	return weatherDataChan
}

func BenchmarkCalculateAverageTemperature(b *testing.B) {
	weatherDataChan := simulateWeatherData()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Compute the average temperature for each country using a fixed weather data
		for data := range calculateAverageTemperature(weatherDataChan) {
			if data.Country == "" || data.AvgTemperature == 0.0 {
				b.Fatalf("received empty data")
			}
		}
	}
}
