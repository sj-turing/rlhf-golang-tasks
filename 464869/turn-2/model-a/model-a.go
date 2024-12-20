package main

import (
	"context"
	"fmt"
	"os"

	"464869/turn-2/model-a/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel       string `yaml:"logLevel" env:"LOG_LEVEL" env-default:"info"`
	OutputFileName string `yaml:"outputFileName" env:"OUTPUT_FILE_NAME"`
}

func main() {
	var cfg Config

	// Reading configuration
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Printf("error while reading config file: %v\n", err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile(cfg.OutputFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v, path:%s\n", err, cfg.OutputFileName)
		os.Exit(1)
	}

	log := logger.New(logFile, cfg.LogLevel)

	// Create a context with request and session IDs
	requestID := "request-123"
	sessionID := "session-456"
	ctx := context.WithValue(context.Background(), "requestID", requestID)
	ctx = context.WithValue(ctx, "sessionID", sessionID)

	log.Debug(ctx, "Debug message")
	log.Info(ctx, "Info message")
	log.Warn(ctx, "Warn message")
	log.Error(ctx, "Error message")
	log.Fatal(ctx, "Fatal message")
}
