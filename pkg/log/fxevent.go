package log

import (
	"os"

	"github.com/vnworkday/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FxEventParams struct {
	fx.In
	Logger *zap.Logger
}

func NewFxEvent(params FxEventParams) fxevent.Logger {
	if config.IsLocal() {
		return &fxevent.ConsoleLogger{
			W: os.Stdout,
		}
	}

	return &fxevent.ZapLogger{
		Logger: params.Logger.WithOptions(
			zap.IncreaseLevel(zapcore.WarnLevel),
		),
	}
}
