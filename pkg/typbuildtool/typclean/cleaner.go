package typclean

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(context.Context, *Context) error
}

// Context of clean
type Context struct {
	*typcore.TypicalContext
}
