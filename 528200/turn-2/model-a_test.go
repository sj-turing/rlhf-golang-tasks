// filename: main_test.go
package main

import (
	"testing"
)

// simulateWeatherData generates a fixed number of weather data entries
func simulateWeatherData(numEntries int) <-chan WeatherData {
	weatherChan := make(chan WeatherData, numEntries)
	go func() {
		defer close(weatherChan)
		for i := 0; i < numEntries; i++ {
			city := cities[i%len(cities)]
			weatherChan <- WeatherData{
				City:        city,
				Country:     cityMappedToCountry[city],
				Temperature: float32(rand.Intn(55)),
			}
		}
	}()
	return weatherChan
}

func BenchmarkCalculateAverageTemperature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Reset timer to exclude setup time
		b.StopTimer()
		numEntries := 10000 // Number of weather data entries for benchmarking
		weatherDataChan := simulateWeatherData(numEntries)
		b.StartTimer()

		// Run the function being benchmarked
		avgTempChan := calculateAverageTemperature(weatherDataChan)
		for range avgTempChan {
			// Drain the channel to ensure all processing is completed
		}
	}
}
