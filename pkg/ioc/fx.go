package ioc

import (
	"fmt"

	"go.uber.org/fx"
)

// RegisterWithName registers a constructor with a name and parameters.
// This is useful for providing multiple instances of the same type.
//
// Parameters:
//   - constructor: The constructor to register, should be a function that returns an instance of the type to register.
//   - name: The unique name used to register with container.
//
// Returns:
//   - the constructor annotated with the name and parameters.
//
// Example:
//
//	 fx.Provide(
//		    ioc.RegisterWithName(NewService, "service1"),
//		    ioc.RegisterWithName(NewService, "service2"),
//	 )
func RegisterWithName(constructor any, name ...string) any {
	if len(name) == 0 {
		return constructor
	}

	annotations := make([]fx.Annotation, 0)

	id := name[0]

	if id != "" {
		annotations = append(annotations,
			fx.ResultTags(`name:"`+id+`"`),
		)
	}

	return fx.Annotate(
		constructor,
		annotations...,
	)
}

// RegisterWithGroup registers a constructor with a group name and type.
// This is useful for grouping constructors together.
//
// Parameters:
//   - groupCreation: the constructor to register. This should be a function that returns an instance of the group type
//     given in the groupType parameter.
//   - groupName: the name of the group to register the constructor with.
//   - groupType: the type of the group to register the constructor with. This should be an interface type.
//
// Returns:
//   - the constructor annotated with the group name and type.
//
// Example:
//
//	 fx.Provide(
//		    ioc.RegisterWithGroup(NewGroup, "my-group", new(MyGroup)),
//	 )
func RegisterWithGroup(groupCreation any, groupName string, groupType any) any {
	tag := fmt.Sprintf(`group:"%s"`, groupName)

	return fx.Annotate(
		groupCreation,
		fx.As(groupType),
		fx.ResultTags(tag),
	)
}
