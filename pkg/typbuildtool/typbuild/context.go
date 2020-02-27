package typbuild

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Context of build
type Context struct {
	*typcore.TypicalContext
	*typast.Store
}
