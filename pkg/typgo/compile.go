package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/execkit"
)

type (
	// Compiler responsible to compile
	Compiler interface {
		Compile(*Context) error
	}
	// Compilers for composite compile
	Compilers []Compiler
	// CompileFn compile function
	CompileFn    func(*Context) error
	compilerImpl struct {
		fn CompileFn
	}
	// StdCompile is standard compile
	StdCompile struct {
		Before Compiler
		After  Compiler
	}
)

//
// stdCompiler
//

var _ (Compiler) = (*StdCompile)(nil)

// NewCompiler return new instance of Compiler
func NewCompiler(fn CompileFn) Compiler {
	return &compilerImpl{fn: fn}
}

func (i *compilerImpl) Compile(c *Context) error {
	return i.fn(c)
}

//
// Commpilers
//

var _ (Compiler) = (Compilers)(nil)

// Compile array of compiler
func (s Compilers) Compile(c *Context) error {
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
func (s *StdCompile) Compile(c *Context) error {
	if s.Before != nil {
		if err := s.Before.Compile(c); err != nil {
			return err
		}
	}
	if err := s.compile(c); err != nil {
		return err
	}
	if s.After != nil {
		if err := s.After.Compile(c); err != nil {
			return err
		}
	}
	return nil
}

func (s *StdCompile) compile(c *Context) error {
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
