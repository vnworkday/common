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
		registerMock(RegisterWithName(NewMockDep, "dep1")),
		registerMock(RegisterWithName(NewMockDep, "dep2")),
		registerMock(RegisterWithName(NewMockService)),
	)

	app.RequireStart()
	defer app.RequireStop()
}

func TestRegisterServicesInGroup(t *testing.T) {
	t.Parallel()

	app := fxtest.New(t,
		registerMock(RegisterWithGroup(NewMockDep, "deps", new(MockDep))),
		registerMock(RegisterWithGroup(NewMockDep, "deps", new(MockDep))),
		registerMock(RegisterWithName(NewMockServiceGroup)),
	)

	app.RequireStart()
	defer app.RequireStop()
}
