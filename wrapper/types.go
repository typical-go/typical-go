package wrapper

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Context of wrap
type Context struct {
	*typcore.Descriptor
	Ctx            context.Context
	TypicalTmp     string
	ProjectPackage string
}
