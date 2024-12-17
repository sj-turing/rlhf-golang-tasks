package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port  int    `json:"port"`
	DBURL string `json:"db_url"`
	Debug bool   `json:"debug,omitempty"` // Example of an optional field
}

func loadConfig() *Config {
	// Read the configuration file
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}

	// Overwrite config with environment variables
	if port := os.Getenv("PORT"); port != "" {
		config.Port, err = strconv.Atoi(port)
		if err != nil {
			log.Warnf("Invalid PORT environment variable: %v, using config value", err)
		} else {
			fmt.Println("Overwriting PORT from environment:", port)
		}
	}

	if dbURL := os.Getenv("DB_URL"); dbURL != "" {
		config.DBURL = dbURL
		fmt.Println("Overwriting DB_URL from environment:", dbURL)
	}

	return &config
}

func main() {
	config := loadConfig()

	// Display the final configuration
	fmt.Println("Final Configuration:")
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("DB URL: %s\n", config.DBURL)
	fmt.Printf("Debug: %t\n", config.Debug)

	// Example of using the configuration
	// Start server, connect to DB, etc.
}
