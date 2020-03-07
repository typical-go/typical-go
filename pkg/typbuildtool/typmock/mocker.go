package typmock

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/urfave/cli/v2"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Mocker responsible to mock
type Mocker interface {
	Mock(*Context) error
}

// Context of mock
type Context struct {
	*typcore.TypicalContext
	*typast.Store
	Cli *cli.Context
}
