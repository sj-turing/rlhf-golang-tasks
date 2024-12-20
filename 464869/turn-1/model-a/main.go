package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	logpkg "464869/turn-1/model-a/logger"
)

func main() {
	config, err := logpkg.ParseConfig("logger.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing configuration:", err)
		return
	}

	logger := logpkg.NewLogger(config.MinimumLevel)

	// Add outputs from configuration
	for _, output := range config.Outputs {
		file, err := os.Create(output.Filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating file output:", err)
			continue
		}
		defer file.Close()
		logger.AddOutput(file, output.Level, output.Template)
	}

	// Log messages of different levels
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warning("Warning message")
	logger.Error("Error message")
	logger.Critical("Critical message")

	// Handle SIGINT to clean up
	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for _ = range sighup {
		fmt.Fprintln(os.Stderr, "Received signal, exiting...")
		return
	}
}
