package main

import (
	"fmt"
	"os"

	"464869/turn-1/1-ideal-response/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct which matches the configuration file data
type Config struct {
	LogLevel       string `yaml:"logLevel" env:"LOG_LEVEL" env-default:"info"`
	OutputFileName string `yaml:"outputFileName" env:"OUTPUT_FILE_NAME"`
}

func main() {
	var cfg Config

	// reading a config.yml file using cleanenv
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Printf("error while reading config file: %v\n", err)
		os.Exit(1)
	}

	// opening a file to write log data into it
	logFile, err := os.OpenFile(cfg.OutputFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v, path:%s\n", err, cfg.OutputFileName)
		os.Exit(1)
	}

	// declaring a new logger
	log := logger.New(logFile, cfg.LogLevel)

	// printing messages
	log.Debug("Debug message")
	log.Info("Info message")
	log.Warn("Warn message")
	log.Error("Error message")
	log.Fatal("Fatal message")
}
