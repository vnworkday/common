package log

import (
	"github.com/vnworkday/config"
	"go.uber.org/zap"
	"os"
)

func NewMockConfig(local bool) func() {
	originalProfile := config.GetProfile()

	if local {
		_ = os.Setenv("profile", "local")
	} else {
		_ = os.Setenv("profile", "prod")
	}

	return func() {
		_ = os.Setenv("profile", originalProfile)
	}
}

func NewMockLogger() *zap.Logger {
	logger, _ := NewLogger()

	return logger
}
