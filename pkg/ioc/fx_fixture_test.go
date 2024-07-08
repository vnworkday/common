package ioc

import "go.uber.org/fx"

type MockDep interface {
	DoSomething()
}

type mockDep struct{}

func (m mockDep) DoSomething() {
	// do something
}

//nolint:ireturn
func NewMockDep() MockDep {
	return &mockDep{}
}

type MockService struct {
	Dep1 MockDep
	Dep2 MockDep
}

type MockServiceParams struct {
	fx.In

	Dep1 MockDep `name:"dep1"`
	Dep2 MockDep `name:"dep2"`
}

func NewMockService(params MockServiceParams) *MockService {
	return &MockService{
		Dep1: params.Dep1,
		Dep2: params.Dep2,
	}
}

type MockServiceGroup struct {
	Deps []*MockDep
}

type MockServiceGroupParams struct {
	fx.In

	Deps []*MockDep `group:"deps"`
}

func NewMockServiceGroup(params MockServiceGroupParams) *MockServiceGroup {
	return &MockServiceGroup{
		Deps: params.Deps,
	}
}
