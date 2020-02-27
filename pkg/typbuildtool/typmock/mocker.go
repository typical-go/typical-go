package typmock

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typast"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Mocker responsible to mock
type Mocker interface {
	Mock(context.Context, *Context) error
}

// Context of mock
type Context struct {
	*typcore.TypicalContext
	*typast.Store
}
