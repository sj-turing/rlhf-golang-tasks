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
)

type Level uint32

const (
	FatalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func ParseLevel(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	default:
		return InfoLevel
	}
}

type LogImpler interface {
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})

	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
}

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

func (l *Logger) IsLevelEnabled(level Level) bool {
	return l.Level >= level
}

func (l *Logger) log(ctx context.Context, level string, msg string) {
	serialized, err := l.formatter(ctx, level, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	serialized = append(serialized, '\n')
	if _, err := l.out.Write(serialized); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}
}

func (l *Logger) Debug(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(DebugLevel) {
		l.log(ctx, "debug", fmt.Sprint(args...))
	}
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	if l.IsLevelEnabled(DebugLevel) {
		l.log(ctx, "debug", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(InfoLevel) {
		l.log(ctx, "info", fmt.Sprint(args...))
	}
}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	if l.IsLevelEnabled(InfoLevel) {
		l.log(ctx, "info", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warn(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(WarnLevel) {
		l.log(ctx, "warn", fmt.Sprint(args...))
	}
}

func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	if l.IsLevelEnabled(WarnLevel) {
		l.log(ctx, "warn", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(ErrorLevel) {
		l.log(ctx, "error", fmt.Sprint(args...))
	}
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	if l.IsLevelEnabled(ErrorLevel) {
		l.log(ctx, "error", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Fatal(ctx context.Context, args ...interface{}) {
	if l.IsLevelEnabled(FatalLevel) {
		l.log(ctx, "fatal", fmt.Sprint(args...))
		os.Exit(1)
	}
}

func (l *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	if l.IsLevelEnabled(FatalLevel) {
		l.log(ctx, "fatal", fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}

const (
	defaultTimestampFormat = time.RFC3339
)

type Fields struct {
	Message   string `json:"message"`
	Time      string `json:"time"`
	Level     string `json:"level"`
	RequestID string `json:"requestId,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
}

func (f *Fields) Reset() {
	f = nil
}

func (l *Logger) formatter(ctx context.Context, level string, msg string) ([]byte, error) {
	f := l.dataPool.Get().(*Fields)
	f.Reset()
	defer l.dataPool.Put(f)

	f.Level = level
	f.Message = msg
	f.Time = time.Now().Format(defaultTimestampFormat)

	// Extract requestID and sessionID from context
	if requestID, ok := ctx.Value("requestID").(string); ok {
		f.RequestID = requestID
	}
	if sessionID, ok := ctx.Value("sessionID").(string); ok {
		f.SessionID = sessionID
	}

	return json.Marshal(f)
}
