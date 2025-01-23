package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WeatherData struct {
	City        string
	Country     string
	Temperature float64
}

func fetchWeatherData(cities []string) []WeatherData {
	data := []WeatherData{}
	for _, city := range cities {
		// Simulate data fetching with random temperatures
		data = append(data, WeatherData{
			City:        city,
			Country:     getCountryForCity(city),
			Temperature: rand.Float64() * 40, // Simulate temperatures
		})
	}
	return data
}

func getCountryForCity(city string) string {
	// Simplified example mapping
	cityCountryMap := map[string]string{
		"New York":    "USA",
		"Los Angeles": "USA",
		"Toronto":     "Canada",
		"Vancouver":   "Canada",
		"London":      "UK",
	}
	return cityCountryMap[city]
}

func calculateAverageTemperature(data []WeatherData) map[string]float64 {
	countryTempMap := make(map[string][]float64)
	for _, weather := range data {
		countryTempMap[weather.Country] = append(countryTempMap[weather.Country], weather.Temperature)
	}

	averageTempMap := make(map[string]float64)
	for country, temps := range countryTempMap {
		total := 0.0
		for _, temp := range temps {
			total += temp
		}
		averageTempMap[country] = total / float64(len(temps))
	}
	return averageTempMap
}

func main() {
	cities := []string{"New York", "Los Angeles", "Toronto", "Vancouver", "London"}
	data := fetchWeatherData(cities)
	averages := calculateAverageTemperature(data)
	fmt.Println("Average temperatures:", averages)
}
