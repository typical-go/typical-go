package typtest

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Tester responsible to test the project
type Tester interface {
	Test(context.Context, *Context) error
}

// Context of test
type Context struct {
	*typcore.TypicalContext
}
