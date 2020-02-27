package typrun

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Runner responsible to run the project
type Runner interface {
	Run(context.Context, *Context) error
}

// Context of run
type Context struct {
	*typcore.TypicalContext
	Binary string
	Args   []string
}
