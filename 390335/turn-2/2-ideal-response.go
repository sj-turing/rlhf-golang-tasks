package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Custom error type
type DataProcessingError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Context string `json:"context"`
}

func (e *DataProcessingError) Error() string {
	return fmt.Sprintf("DataProcessingError: %d - %s (Context: %s)", e.Code, e.Message, e.Context)
}

// DataStream represents a data stream
type DataStream struct {
	data chan int
}

// NewDataStream initializes a new data stream
func NewDataStream() *DataStream {
	return &DataStream{
		data: make(chan int),
	}
}

// Start produces data
func (ds *DataStream) Start() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ds.data <- i
		}
		close(ds.data)
	}()
}

// Process processes data and handles errors
func Process(ctx context.Context, data <-chan int, logger *zerolog.Logger) {
	for val := range data {
		if err := doWork(ctx, val); err != nil {
			logger.Error().Err(err).Msg("Error processing data")
		}
	}
}

// doWork simulates work that can fail
func doWork(ctx context.Context, value int) error {
	if rand.Float32() < 0.2 {
		return &DataProcessingError{
			Code:    500,
			Message: "failed to process value",
			Context: ctx.Value("context").(string),
		}
	}
	return nil
}

// ErrorHandler handles errors
func ErrorHandler(w http.ResponseWriter, r *http.Request, logger *zerolog.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	err := r.Context().Err()
	if err == nil {
		err = errors.New("unknown error")
	}

	logger.Error().Err(err).Msg("Internal server error")

	response := map[string]string{"error": err.Error()}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error().Err(err).Msg("Failed to encode error response")
	}
}

func main() {

	// Set up logging
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	ctx := context.WithValue(context.Background(), "context", "main pipeline")

	dataStream := NewDataStream()
	dataStream.Start()

	go Process(ctx, dataStream.data, &logger)

	// Simulate a simple HTTP server with error handling
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			http.Error(w, "timeout", http.StatusGatewayTimeout)
		default:
			fmt.Fprintln(w, "Hello, World!")
		}
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("simulated panic")
	})

	http.HandleFunc("/internal-error", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "context", "/internal-error endpoint")
		r = r.WithContext(ctx)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	})

	http.HandleFunc("/recover", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		panic("another simulated panic")
	})

	logger.Info().Msg("Starting HTTP server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
