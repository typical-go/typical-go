package typwrap

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Wrapper responsible to wrap the project
type Wrapper interface {
	Wrap(*Context) error
}

// Context of wrap
type Context struct {
	*typcore.Descriptor
	Ctx            context.Context
	TypicalTmp     string
	ProjectPackage string
}
