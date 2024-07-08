package ioc

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func registerMock(constructor any) fx.Option {
	return fx.Provide(constructor)
}

func TestRegisterServiceWithDifferentNames(t *testing.T) {
	t.Parallel()

	app := fxtest.New(t,
		registerMock(FxRegisterWithName(NewMockDep, "dep1")),
		registerMock(FxRegisterWithName(NewMockDep, "dep2")),
		registerMock(FxRegisterWithName(NewMockService)),
	)

	app.RequireStart()
	defer app.RequireStop()
}

func TestRegisterServicesInGroup(t *testing.T) {
	t.Parallel()

	app := fxtest.New(t,
		registerMock(FxRegisterWithGroup(NewMockDep, "deps", new(MockDep))),
		registerMock(FxRegisterWithGroup(NewMockDep, "deps", new(MockDep))),
		registerMock(FxRegisterWithName(NewMockServiceGroup)),
	)

	app.RequireStart()
	defer app.RequireStop()
}
