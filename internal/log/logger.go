package log

import (
	"sync"

	"go.uber.org/zap"
)

// Logger is a wrapper around zap.SugaredLogger to provide a singleton logger instance.
type Logger struct {
	*zap.SugaredLogger
}

// Variables for the singleton logger instance and synchronization.
var (
	logger *Logger
	Once   sync.Once
)

// BuildLoggerFunc is a function type that builds a zap.Logger. It allows for dependency injection in tests.
type BuildLoggerFunc func(cfg zap.Config) (*zap.Logger, error)

// DefaultBuildLogger is the default function that builds a zap.Logger.
var DefaultBuildLogger BuildLoggerFunc = func(cfg zap.Config) (*zap.Logger, error) {
	return cfg.Build()
}

// NewLogger initializes and returns a singleton Logger instance with the specified log level.
// The function panics if the provided log level is invalid.
func NewLogger(level string, buildLogger BuildLoggerFunc) *Logger {
	Once.Do(func() {
		// Parse the provided log level
		lvl, err := zap.ParseAtomicLevel(level)
		if err != nil {
			panic(err)
		}

		// Create a new zap logger configuration
		cfg := zap.NewProductionConfig()
		cfg.Level = lvl
		cfg.Encoding = "console"
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}

		// Build the zap logger with the configured settings
		l, err := buildLogger(cfg)
		if err != nil {
			panic(err)
		}

		// Create a SugaredLogger from the zap logger
		logger = &Logger{l.Sugar()}
	})

	return logger
}
