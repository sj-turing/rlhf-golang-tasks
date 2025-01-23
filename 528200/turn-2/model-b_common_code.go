package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type WeatherData struct {
	City        string
	Country     string
	Temperature float32
}

var cities = []string{"abc", "xyz", "pqr", "jkl", "mno", "qwe", "asd", "uip"}
var cityMappedToCountry = map[string]string{
	"abc": "country1",
	"pqr": "country1",
	"jkl": "country1",
	"asd": "country1",
	"mno": "country2",
	"qwe": "country2",
	"xyz": "country2",
	"uip": "country3",
}

// getWeatherData generates random weather data for city
func getWeatherData(ctx context.Context) <-chan WeatherData {
	weatherChan := make(chan WeatherData)

	go func() {
		defer close(weatherChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(10 * time.Millisecond)
				city := cities[rand.Intn(len(cities))]
				weatherChan <- WeatherData{
					City:        city,
					Country:     cityMappedToCountry[city],
					Temperature: float32(rand.Intn(55)),
				}
			}
		}
	}()

	return weatherChan
}

type Metadata struct {
	Country        string
	AvgTemperature float32
}

// calculateAverageTemperature calculates countrywise average temperature
func calculateAverageTemperature(weatherDataChan <-chan WeatherData) <-chan Metadata {

	// keeps country with its total temperature
	totalTemperatureByCountry := map[string]float32{}
	// keeps country with number of total collected temperature
	countByCountry := map[string]float32{}

	// reads from the weatherDataChan
	for wd := range weatherDataChan {
		if v, ok := totalTemperatureByCountry[wd.Country]; ok {
			totalTemperatureByCountry[wd.Country] = v + wd.Temperature
			countByCountry[wd.Country] += 1
		} else {
			totalTemperatureByCountry[wd.Country] = wd.Temperature
			countByCountry[wd.Country] = 1
		}
	}

	averageTempChan := make(chan Metadata, len(countByCountry))

	go func() {
		defer close(averageTempChan)
		for country, totalTemperature := range totalTemperatureByCountry {
			totalRecords := countByCountry[country]

			// calculating average temperature
			averageTempChan <- Metadata{Country: country, AvgTemperature: totalTemperature / totalRecords}
		}
	}()

	return averageTempChan
}

func main() {
	// waiting 2 seconds to wait the main function's goroutine
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// generate and returns a weather data
	weatherDataChan := getWeatherData(ctx)

	// print outs the countrywise calculated average temperature
	for avgTemp := range calculateAverageTemperature(weatherDataChan) {
		fmt.Printf("country = %s with average temperature = %.2f\n", avgTemp.Country, avgTemp.AvgTemperature)
	}
}
