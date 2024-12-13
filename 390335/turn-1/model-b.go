package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Custom error type to carry metadata
type DataProcessingError struct {
	error
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Context map[string]interface{} `json:"context,omitempty"`
}

func (e *DataProcessingError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.error, e.Code)
}

func newDataProcessingError(err error, code int, message string, context map[string]interface{}) *DataProcessingError {
	return &DataProcessingError{
		error:   err,
		Code:    code,
		Message: message,
		Context: context,
	}
}

// Simulate a data source
func fetchData(ctx context.Context) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fetchData")
	defer span.Finish()

	ext.SpanKindRPCClient.Set(span)
	span.SetTag("component", "data_source")

	time.Sleep(time.Second)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return []byte(`{"id": 1, "value": "data"}`), nil
}

// Simulate data processing
func processData(ctx context.Context, data []byte) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "processData")
	defer span.Finish()

	ext.SpanKindRPCServer.Set(span)
	span.SetTag("component", "data_processor")

	var record map[string]interface{}
	if err := json.Unmarshal(data, &record); err != nil {
		return newDataProcessingError(err, 1001, "Failed to unmarshal data", map[string]interface{}{"data": string(data)})
	}

	if _, ok := record["id"]; !ok {
		return newDataProcessingError(errors.New("missing 'id' in data"), 1002, "Data validation failed", record)
	}

	time.Sleep(time.Second)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return nil
}

// Simulate data storage
func storeData(ctx context.Context, data []byte) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "storeData")
	defer span.Finish()

	ext.SpanKindRPCServer.Set(span)
	span.SetTag("component", "data_store")

	time.Sleep(time.Second)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return nil
}

func main() {
	tracer := opentracing.GlobalTracer()

	logger := log.NewStdoutLogger()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := fetchData(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	if err := processData(ctx, data); err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	if err := storeData(ctx, data); err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	level.Info(logger).Log("msg", "Data pipeline completed successfully")
}
