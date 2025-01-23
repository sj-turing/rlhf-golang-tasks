package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkWeatherProcessing(b *testing.B) {
	cities := []string{"New York", "Los Angeles", "Toronto", "Vancouver", "London"}
	for i := 0; i < b.N; i++ {
		data := fetchWeatherData(cities)
		_ = calculateAverageTemperature(data)
	}
}
