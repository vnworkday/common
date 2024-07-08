package log

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestLoggerInitInDevelopmentMode(t *testing.T) {
	t.Parallel()

	restoreConfig := NewMockConfig(true)
	defer restoreConfig()

	logger, err := NewLogger()
	if err != nil {
		t.Fatalf("Failed to create logger in development mode: %v", err)
	}

	enabled := logger.Core().Enabled(zapcore.DebugLevel)
	if !enabled {
		t.Errorf("Expected logger to be enabled for DebugLevel, got %t", enabled)
	}
}

func TestLoggerInitInProductionMode(t *testing.T) {
	t.Parallel()

	restoreConfig := NewMockConfig(false)
	defer restoreConfig()

	logger, err := NewLogger()
	if err != nil {
		t.Fatalf("Failed to create logger in production mode: %v", err)
	}

	enabled := logger.Core().Enabled(zapcore.DebugLevel)
	if enabled {
		t.Errorf("Expected logger to be disabled for DebugLevel, got %t", enabled)
	}
}
