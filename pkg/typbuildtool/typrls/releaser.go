package typrls

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Releaser responsible to release
type Releaser interface {
	Release(*Context) (err error)
}

// Context of release
type Context struct {
	*typcore.TypicalContext
	Cli   *cli.Context
	Alpha bool
}
