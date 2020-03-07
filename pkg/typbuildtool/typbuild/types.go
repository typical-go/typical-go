package typbuild

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Builder reponsible to build
type Builder interface {
	Build(c *Context) (bin string, err error)
}

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*Context) error
}

// Prebuilder responsible to prebuild
type Prebuilder interface {
	Prebuild(c *Context) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*Context) error
}

// Runner responsible to run the project
type Runner interface {
	Run(*RunContext) error
}

// Context of build
type Context struct {
	*typcore.TypicalContext
	*typast.Store
	Cli *cli.Context
}

// RunContext of run
type RunContext struct {
	*Context
	Binary string
}
