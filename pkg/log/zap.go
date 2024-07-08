package log

import (
	"github.com/vnworkday/config"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	if config.IsLocal() {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}
