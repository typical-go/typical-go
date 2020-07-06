package typgo

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// Compiler responsible to compile
	Compiler interface {
		Compile(*Context) error
	}

	// Compiles for composite compile
	Compiles []Compiler

	// CompileFn compile function
	CompileFn func(*Context) error

	compilerImpl struct {
		fn CompileFn
	}

	// StdCompile is standard compile
	StdCompile struct{}
)

var _ (Compiler) = (*StdCompile)(nil)
var _ (Compiler) = (Compiles)(nil)

//
// stdCompiler
//

// NewCompile return new instance of Compiler
func NewCompile(fn CompileFn) Compiler {
	return &compilerImpl{fn: fn}
}

func (i *compilerImpl) Compile(c *Context) error {
	return i.fn(c)
}

//
// Commpilers
//

// Compile array of compiler
func (s Compiles) Compile(c *Context) error {
	for _, compiler := range s {
		if err := compiler.Compile(c); err != nil {
			return err
		}
	}
	return nil
}

//
// StdCompile
//

// Compile standard go project
func (*StdCompile) Compile(c *Context) (err error) {
	src := fmt.Sprintf("%s/%s", CmdFolder, c.Descriptor.Name)

	return c.Execute(&execkit.GoBuild{
		Out:    AppBin(c.Descriptor.Name),
		Source: "./" + src,
		Ldflags: []string{
			execkit.BuildVar("github.com/typical-go/typical-go/pkg/typapp.Name", c.Descriptor.Name),
			execkit.BuildVar("github.com/typical-go/typical-go/pkg/typapp.Version", c.Descriptor.Version),
		},
	})
}

//
// Command
//

func cmdCompile(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "compile",
		Aliases: []string{"c"},
		Usage:   "Compile the project",
		Action:  c.ActionFn("COMPILE", compile),
	}
}

func compile(c *Context) error {
	if c.Compile == nil {
		return errors.New("compile is missing")
	}
	return c.Compile.Compile(c)
}
