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
		Before  Compiler
		After   Compiler
		Source  string
		Output  string
		Ldflags fmt.Stringer
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
	if s.Source == "" {
		s.Source = fmt.Sprintf("./cmd/%s", c.Descriptor.Name)
	}

	if s.Output == "" {
		s.Output = fmt.Sprintf("bin/%s", c.Descriptor.Name)
	}

	if s.Ldflags == nil {
		s.Ldflags = execkit.BuildVars{
			"github.com/typical-go/typical-go/pkg/typapp.Name":    c.Descriptor.Name,
			"github.com/typical-go/typical-go/pkg/typapp.Version": c.Descriptor.Version,
		}
	}

	return c.Execute(&execkit.GoBuild{
		Output:  s.Output,
		Source:  s.Source,
		Ldflags: s.Ldflags,
	})
}
