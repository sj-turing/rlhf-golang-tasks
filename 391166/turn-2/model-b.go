package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

const (
	// Set a reasonable timeout for the context
	streamTimeout = 5 * time.Second
)

func init() {
	// Configure structured logging
	zerolog.SetFormatter(zerolog.JSONFormatter())
	zerolog.TimeFieldFormat = time.RFC3339
}

func streamData(ctx context.Context, filePath string, logger *zerolog.Logger) error {
	logger = logger.WithCtx(ctx).Ctx

	file, err := os.Open(filePath)
	if err != nil {
		return logger.Err(err).Msg("failed to open file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Err(err).Msg("error closing file")
		}
	}()

	reader := bufio.NewReader(file)

	for {
		select {
		case <-ctx.Done():
			logger.Debug().Msg("streaming context cancelled")
			return ctx.Err()
		default:
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					logger.Debug().Msg("end of file reached")
					return nil
				}
				return logger.Err(err).Msg("error reading line")
			}

			if isPrefix {
				logger.Warn().Msg("incomplete line read")
			}

			processLine(ctx, line, logger)
		}
	}
}

func processLine(ctx context.Context, line []byte, logger *zerolog.Logger) {
	// Simulate processing logic
	logger.Debug().Str("line", string(line)).Msg("processing line")

	select {
	case <-ctx.Done():
		logger.Warn().Msg("processing cancelled")
		return
	case <-time.After(100 * time.Millisecond):
		// Simulate a processing error 10% of the time
		if rand.Float64() < 0.1 {
			logger.Err(errors.New("processing error simulated")).Msg("error processing line")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), streamTimeout)
	defer cancel()

	filePath := "large_data.txt"

	err := streamData(ctx, filePath, logger.With().Str("operation", "streamData"))
	if err != nil {
		logger.Error().Err(err).Msg("error streaming data")
	}
}
