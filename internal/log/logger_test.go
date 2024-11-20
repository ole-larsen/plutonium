package log_test

import (
	"bytes"
	"fmt"

	"strings"
	"sync"
	"testing"

	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MockCore is a custom zapcore.Core used to capture log outputs in memory for testing.
type MockCore struct {
	zapcore.Core
	encoder zapcore.Encoder
	buffer  *bytes.Buffer
}

// With returns a new MockCore instance with additional fields added.
func (m *MockCore) With(fields []zapcore.Field) zapcore.Core {
	return &MockCore{
		Core:    m.Core.With(fields),
		encoder: m.encoder,
		buffer:  m.buffer,
	}
}

// Check determines whether the log entry should be logged at the specified level.
//
//nolint:gocritic // explanation: hugeParam: entry is heavy (136 bytes); consider passing it by pointer
func (m *MockCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return m.Core.Check(entry, checkedEntry)
}

// Write writes the log entry to the buffer.
//
//nolint:gocritic // explanation: hugeParam: entry is heavy (136 bytes); consider passing it by pointer
func (m *MockCore) Write(entry zapcore.Entry, _ []zapcore.Field) error {
	fmt.Fprintf(m.buffer, "%s: %s\n", entry.Level, entry.Message)
	return nil
}

// Sync synchronizes the buffer, if necessary.
func (m *MockCore) Sync() error {
	return nil
}

// createMockLogger creates a Logger instance with a MockCore for testing.
func createMockLogger(_ string, buffer *bytes.Buffer) *log.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	mockCore := &MockCore{
		encoder: zapcore.NewConsoleEncoder(encoderCfg),
		buffer:  buffer,
	}

	levelEnabler := zapcore.DebugLevel
	core := zapcore.NewCore(mockCore.encoder, zapcore.AddSync(mockCore.buffer), levelEnabler)

	zapLogger := zap.New(core)

	return &log.Logger{zapLogger.Sugar()}
}

func TestNewLogger(t *testing.T) {
	levels := []struct {
		name  string
		level string
		msg   string
	}{
		{"debug", "debug", "Test message at level: debug"},
		{"info", "info", "Test message at level: info"},
		{"warn", "warn", "Test message at level: warn"},
		{"error", "error", "Test message at level: error"},
		{"panic", "panic", "Test message at level: panic"},
		// {"fatal", "fatal", "Test message at level: fatal"},
	}

	for _, tt := range levels {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			log.Once = sync.Once{} // Reset the singleton

			// Create a logger
			logger := createMockLogger(tt.level, buffer)

			// Log the message
			switch tt.level {
			case "debug":
				logger.Debugf(tt.msg)
			case "info":
				logger.Infof(tt.msg)
			case "warn":
				logger.Warnf(tt.msg)
			case "error":
				logger.Errorf(tt.msg)
			case "panic":
				defer func() {
					if r := recover(); r != nil {
						// Handle panic gracefully
						require.Equal(t, tt.msg, r)
					}
				}()
				logger.Panicf(tt.msg)
			case "fatal":
				// Simulate fatal without stopping the test suite
				func() {
					defer func() {
						if r := recover(); r != nil {
							require.Equal(t, tt.msg, r)
						}
					}()
					logger.Fatalf(tt.msg)
				}()
			}

			output := buffer.String()
			if !strings.Contains(output, tt.msg) && tt.name != "fatal" {
				t.Errorf("Expected log output to contain '%s', got: '%s'", tt.msg, output)
			}
		})
	}
}

func TestLoggerSingleton(t *testing.T) {
	log.Once = sync.Once{} // Reset the once variable

	logger1 := log.NewLogger("info", log.DefaultBuildLogger)
	logger2 := log.NewLogger("debug", log.DefaultBuildLogger)

	if logger1 != logger2 {
		t.Error("Expected logger1 and logger2 to be the same instance")
	}
}

func TestInvalidLogLevel(t *testing.T) {
	// Capture any panics that occur during the execution of NewLogger
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			// Verify that the panic message contains an indication of an invalid log level
			require.True(t, ok)

			require.True(t, strings.Contains(err.Error(), "unrecognized level"), "Expected panic due to invalid log level, but got a different message")
		} else {
			// Fail the test if no panic occurs
			t.Error("Expected panic for invalid log level, but no panic occurred")
		}
	}()

	log.Once = sync.Once{} // Reset the once variable
	_ = log.NewLogger("invalid", log.DefaultBuildLogger)
}

func TestLoggerBuildError(t *testing.T) {
	// Define a build function that simulates an error
	buildErrorFunc := func(_ zap.Config) (*zap.Logger, error) {
		return nil, fmt.Errorf("build error")
	}

	// Capture any panics that occur during the execution of NewLogger
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			require.True(t, ok)
			require.True(t, strings.Contains(err.Error(), "build error"), "Expected panic due to build error, but got a different message")
		} else {
			// Fail the test if no panic occurs
			t.Error("Expected panic for build error, but no panic occurred")
		}
	}()

	log.Once = sync.Once{} // Reset the once variable
	_ = log.NewLogger("info", buildErrorFunc)
}
