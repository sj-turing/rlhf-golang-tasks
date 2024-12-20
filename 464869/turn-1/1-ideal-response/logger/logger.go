package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Level type for introducing possible log levels
type Level uint32

const (
	FatalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// ParseLevel parses string and returns valid Log Level
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

// LogImpler the interface provide the methods to send log data by type
type LogImpler interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
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

func (l *Logger) log(level string, msg string) {
	serialized, err := l.formatter(level, msg)
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

func (l *Logger) Debug(args ...interface{}) {
	if l.IsLevelEnabled(DebugLevel) {
		l.log("debug", fmt.Sprint(args...))
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.IsLevelEnabled(DebugLevel) {
		l.log("debug", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Info(args ...interface{}) {
	if l.IsLevelEnabled(InfoLevel) {
		l.log("info", fmt.Sprint(args...))
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.IsLevelEnabled(InfoLevel) {
		l.log("info", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warn(args ...interface{}) {
	if l.IsLevelEnabled(WarnLevel) {
		l.log("warn", fmt.Sprint(args...))
	}
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.IsLevelEnabled(WarnLevel) {
		l.log("warn", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(args ...interface{}) {
	if l.IsLevelEnabled(ErrorLevel) {
		l.log("error", fmt.Sprint(args...))
	}
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.IsLevelEnabled(ErrorLevel) {
		l.log("error", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Fatal(args ...interface{}) {
	if l.IsLevelEnabled(FatalLevel) {
		l.log("fatal", fmt.Sprint(args...))
		os.Exit(1)
	}
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.IsLevelEnabled(FatalLevel) {
		l.log("fatal", fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}

const (
	defaultTimestampFormat = time.RFC3339
)

type Fields struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Level   string `json:"level"`
}

func (f *Fields) Reset() {
	f = nil
}

func (l *Logger) formatter(level string, msg string) ([]byte, error) {
	f := l.dataPool.Get().(*Fields)
	f.Reset()
	defer l.dataPool.Put(f)

	f.Level = level
	f.Message = msg
	f.Time = time.Now().Format(defaultTimestampFormat)

	return json.Marshal(f)
}
