package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("json") // Supported types: json, yaml, toml

	// Set the path where the config file is located
	viper.AddConfigPath(".")

	// Read environment variables (automatic override)
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Show values and let the environment variables override them
	port := viper.GetString("port")
	dbURL := viper.GetString("db_url")
	logLevel := viper.GetString("log_level")

	fmt.Println("Configuration:")
	fmt.Println("Port:", port)
	fmt.Println("DB URL:", dbURL)
	fmt.Println("Log Level:", logLevel)

	// Optionally, use os variables for overriding
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if envDBUrl := os.Getenv("DB_URL"); envDBUrl != "" {
		dbURL = envDBUrl
	}

	fmt.Println("\nEffective Configuration:")
	fmt.Println("Effective Port:", port)
	fmt.Println("Effective DB URL:", dbURL)
	fmt.Println("Effective Log Level:", logLevel)
}
