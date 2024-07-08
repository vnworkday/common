//go:build tools

package tools

// This file imports packages that are used when running go generate, or used by tools that analyze Go code.
import (
	_ "golang.org/x/tools/cmd/goimports"
)
