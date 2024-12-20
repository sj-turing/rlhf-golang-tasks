package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// LogLevel defines the different log levels.
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
	Critical
)

var logLevelStrings = []string{"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}

func (l LogLevel) String() string {
	return logLevelStrings[l]
}

func (l LogLevel) Compare(o LogLevel) int {
	return int(l - o)
}

type Logger struct {
	minLevel        LogLevel
	outputs         []io.Writer
	outputLevels    []LogLevel
	outputTemplates []*bytes.Buffer
	outputMux       *sync.Mutex
}

func NewLogger(minLevel LogLevel) *Logger {
	return &Logger{
		minLevel:        minLevel,
		outputs:         []io.Writer{os.Stdout},
		outputLevels:    []LogLevel{Info},
		outputTemplates: []*bytes.Buffer{bytes.NewBufferString("%s: %v\n")},
		outputMux:       &sync.Mutex{},
	}
}

func (l *Logger) logMessage(level LogLevel, v ...interface{}) {
	if level.Compare(l.minLevel) < 0 {
		return
	}

	l.outputMux.Lock()
	defer l.outputMux.Unlock()

	for i, output := range l.outputs {
		if level.Compare(l.outputLevels[i]) >= 0 {
			template := l.outputTemplates[i]
			message := fmt.Sprintf(template.String(), level.String(), v...)
			if _, err := output.Write([]byte(message)); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to output: %v\n", err)
			}
		}
	}
}

func (l *Logger) SetMinimumLevel(level LogLevel) {
	l.outputMux.Lock()
	defer l.outputMux.Unlock()
	l.minLevel = level
}

func (l *Logger) AddOutput(output io.Writer, level LogLevel, template string) {
	l.outputMux.Lock()
	defer l.outputMux.Unlock()

	templateBuf := bytes.NewBufferString(template)
	l.outputs = append(l.outputs, output)
	l.outputLevels = append(l.outputLevels, level)
	l.outputTemplates = append(l.outputTemplates, templateBuf)
}

func (l *Logger) RemoveOutput(output io.Writer) {
	l.outputMux.Lock()
	defer l.outputMux.Unlock()

	outputIndex := -1
	for i, o := range l.outputs {
		if fmt.Sprint(o) == fmt.Sprint(output) {
			outputIndex = i
			break
		}
	}

	if outputIndex != -1 {
		l.outputs = append(l.outputs[:outputIndex], l.outputs[outputIndex+1:]...)
		l.outputLevels = append(l.outputLevels[:outputIndex], l.outputLevels[outputIndex+1:]...)
		l.outputTemplates = append(l.outputTemplates[:outputIndex], l.outputTemplates[outputIndex+1:]...)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.logMessage(Debug, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.logMessage(Info, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.logMessage(Warning, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.logMessage(Error, v...)
}

func (l *Logger) Critical(v ...interface{}) {
	l.logMessage(Critical, v...)
}

func (l *Logger) Log(level LogLevel, v ...interface{}) {
	l.logMessage(level, v...)
}

type Config struct {
	MinimumLevel LogLevel `yaml:"min_level"`
	Outputs      []Output `yaml:"outputs"`
}

type Output struct {
	Filename string   `yaml:"filename"`
	Level    LogLevel `yaml:"level"`
	Template string   `yaml:"template"`
}

func ParseConfig(filename string) (*Config, error) {
	// Simplified placeholder for parsing logic. In practice, use a library like yaml or json.
	// Example using yaml for demonstration purposes.
	config := &Config{}
	// Assume the file has a valid structure (implement actual parsing here).
	config.MinimumLevel = Debug
	config.Outputs = []Output{
		{Filename: "info.log", Level: Info, Template: "%s: %v\n"},
		{Filename: "error.log", Level: Error, Template: "[%s][%s]: %v\n"},
	}
	return config, nil
}
