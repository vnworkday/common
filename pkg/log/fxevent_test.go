package log

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

//nolint:paralleltest
func TestFxEventLoggerUsesConsoleLoggerInLocalEnvironment(t *testing.T) {
	restoreConfig := NewMockConfig(true)
	defer restoreConfig()

	app := fxtest.New(t,
		fx.Provide(NewMockLogger),
		fx.Invoke(func(l fxevent.Logger) {
			if _, ok := l.(*fxevent.ConsoleLogger); !ok {
				t.Errorf("Expected fxevent.Logger to be ConsoleLogger in local environment")
			}
		}),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return NewFxEvent(FxEventParams{Logger: logger})
		}),
	)

	app.RequireStart()
	defer app.RequireStop()
}

//nolint:paralleltest
func TestFxEventLoggerUsesZapLoggerInNonLocalEnvironment(t *testing.T) {
	restoreConfig := NewMockConfig(false)
	defer restoreConfig()

	app := fxtest.New(t,
		fx.Provide(NewMockLogger),
		fx.Invoke(func(l fxevent.Logger) {
			if _, ok := l.(*fxevent.ZapLogger); !ok {
				t.Errorf("Expected fxevent.Logger to be ZapLogger in non-local environment")
			}
		}),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return NewFxEvent(FxEventParams{Logger: logger})
		}),
	)

	app.RequireStart()
	defer app.RequireStop()
}
