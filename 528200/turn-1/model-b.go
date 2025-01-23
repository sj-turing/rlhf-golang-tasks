package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Weather struct {
	City        string
	Country     string
	Temperature float64
}

func getWeather(city string) Weather {
	// Simulate a blocking operation
	time.Sleep(time.Millisecond * 100)

	// Return sample weather data for each city
	switch city {
	case "New York":
		return Weather{City: city, Country: "USA", Temperature: 25.0}
	case "Paris":
		return Weather{City: city, Country: "France", Temperature: 20.0}
	case "Tokyo":
		return Weather{City: city, Country: "Japan", Temperature: 30.0}
	case "London":
		return Weather{City: city, Country: "UK", Temperature: 18.0}
	case "Rio de Janeiro":
		return Weather{City: city, Country: "Brazil", Temperature: 28.0}
	default:
		return Weather{City: city, Country: "Invalid", Temperature: math.NaN()}
	}
}

func calculateAverageTemperature(countryWeathers map[string][]Weather) map[string]float64 {
	averages := make(map[string]float64)
	for country, weathers := range countryWeathers {
		total := 0.0
		validCount := 0
		for _, weather := range weathers {
			if !math.IsNaN(weather.Temperature) {
				total += weather.Temperature
				validCount++
			}
		}
		if validCount > 0 {
			averages[country] = total / float64(validCount)
		}
	}
	return averages
}

func fetchAndCalculate(cities []string) map[string]float64 {
	var wg sync.WaitGroup
	countryWeathers := make(map[string][]Weather)

	// Concurrently fetch weather data for each city
	for _, city := range cities {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			weather := getWeather(city)
			countryWeathers[weather.Country] = append(countryWeathers[weather.Country], weather)
		}(city)
	}

	wg.Wait()
	return calculateAverageTemperature(countryWeathers)
}

func main() {
	cities := []string{"New York", "Paris", "Tokyo", "London", "Rio de Janeiro", "New York", "Paris", "Tokyo"}
	averages := fetchAndCalculate(cities)
	fmt.Println("Average temperatures:")
	for country, average := range averages {
		fmt.Printf("%s: %.2f\n", country, average)
	}
}
