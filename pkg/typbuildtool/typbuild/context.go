package typbuild

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Context of build
type Context struct {
	*typcore.TypicalContext
	*typast.Ast
	Cli *cli.Context
}

// RunContext of run
type RunContext struct {
	*Context
	Binary string
}
