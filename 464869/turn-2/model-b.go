package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// LogImpler the interface provide the methods to send log data by mentioned method
type LogImpler interface {
	// ... (existing methods)
	Debugc(ctx context.Context, args ...interface{})
	Infoc(ctx context.Context, args ...interface{})
	Warnc(ctx context.Context, args ...interface{})
	Errorc(ctx context.Context, args ...interface{})
	Fatalc(ctx context.Context, args ...interface{})
}

// logger struct
type Logger struct {
	out      io.Writer
	mu       sync.Mutex
	dataPool sync.Pool
	Level    Level
}

func New(outType io.Writer, level string) LogImpler {
	return &Logger{
		out: outType,
		dataPool: sync.Pool{
			New: func() interface{} {
				return &Fields{}
			},
		},
		Level: ParseLevel(level),
	}
}

// NewContext adds request-id and session-id to the context
func NewContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	// generate request-id and session-id if they are not already present in the context
	reqID := ctx.Value("request_id")
	if reqID == nil {
		reqID = uuid.New().String()
		ctx = context.WithValue(ctx, "request_id", reqID)
	}
	sesID := ctx.Value("session_id")
	if sesID == nil {
		sesID = uuid.New().String()
		ctx = context.WithValue(ctx, "session_id", sesID)
	}
	return ctx
}

// Implement logging methods with context
func (l *Logger) logc(ctx context.Context, level string, msg string) {
	reqID := ctx.Value("request_id").(string)
	sesID := ctx.Value("session_id").(string)

	serialized, err := l.formatterc(reqID, sesID, level, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}
	// ... (existing code)
}

func (l *Logger) Debugc(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(DebugLevel) {
		l.logc(ctx, "debug", fmt.Sprint(args...))
	}
}

func (l *Logger) Infoc(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(InfoLevel) {
		l.logc(ctx, "info", fmt.Sprint(args...))
	}
}

// ... (existing methods)
