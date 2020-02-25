package typbuild

import (
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Context of build
type Context struct {
	*typcore.TypicalContext
	*prebld.DeclStore
}
