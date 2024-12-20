package main  
import (  
    "fmt"
    "os"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

// LogLevel represents the different log levels.
type LogLevel string

const (
    // TraceLevel logs are typically voluminous, and are usually disabled in production.
    TraceLevel LogLevel = "trace"
    // DebugLevel logs are usually enabled during development. Commonly used for debugging.
    DebugLevel LogLevel = "debug"
    // InfoLevel logs are the default logging priority.
    InfoLevel LogLevel = "info"
    // WarnLevel logs are more important than Info, but not an error situation that prevents the application from continuing to run.
    WarnLevel LogLevel = "warn"
    // ErrorLevel logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs.
    ErrorLevel LogLevel = "error"
    // FatalLevel logs and then calls `logger.Exit(1)`. It is equivalent to `logger.Panic(err)`.
    FatalLevel LogLevel = "fatal"
    // PanicLevel logs the message, then panics.
    PanicLevel LogLevel = "panic"
)

// logger is a custom logger struct that wraps logrus.Logger.
type logger struct {
    *logrus.Logger
}

// newLogger creates a new logger instance with the specified configuration.
func newLogger(cfg *logConfig) *logger {
    l := logrus.New()

    // Set the log level based on the configuration.
    switch cfg.Level {
    case TraceLevel:
        l.SetLevel(logrus.TraceLevel)
    case DebugLevel:
        l.SetLevel(logrus.DebugLevel)
    case InfoLevel:
        l.SetLevel(logrus.InfoLevel)
    case WarnLevel:
        l.SetLevel(logrus.WarnLevel)
    case ErrorLevel:
        l.SetLevel(logrus.ErrorLevel)
    case FatalLevel:
        l.SetLevel(logrus.FatalLevel)
    case PanicLevel:
        l.SetLevel(logrus.PanicLevel)
    default:
        l.SetLevel(logrus.InfoLevel) // Default to InfoLevel if an invalid level is provided.
    }

    // Set the formatter to JSON for structured logs.
    l.SetFormatter(&logrus.JSONFormatter{})

    // Create a new file writer for logging to the file.
    file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        l.Fatalf("Failed to open log file: %v", err)
    }

    // Set the output to the file writer.
    l.SetOutput(file)

    return &logger{l}
}

// logConfig represents the configuration for the logger.
type logConfig struct {
    Level   LogLevel `mapstructure:"level"`
    FilePath string   `mapstructure:"file_path"`
}

// readConfig reads the configuration from a YAML file.
func readConfig() *logConfig {
    viper.SetConfigName("config")  // name of config file (without extension)
    viper.AddConfigPath(".")       // optionally look for config in the working directory
    err := viper.ReadInConfig()  // Find and read the config file
    if err != nil {  // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    var cfg logConfig
    err = viper.Unmarshal(&cfg)
    if err != nil {
        panic(fmt.Errorf("Fatal error unmarshaling config: %s \n", err))
    }

    return &cfg
