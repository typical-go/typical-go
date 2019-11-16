package typprebuilder

import (
	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"

	"github.com/typical-go/typical-go/pkg/typprebuilder/golang"
	"github.com/typical-go/typical-go/pkg/typenv"
)

type constructor struct {
	ApplicationImports golang.Imports
	Constructors       []string
}

func (g constructor) generate(target string) (err error) {
	defer debugkit.ElapsedTime("Generate constructor")()
	src := golang.NewSourceCode(typenv.Dependency.Package)
	src.Imports = g.ApplicationImports
	src.AddConstructors(g.Constructors...)
	if err = src.Cook(target); err != nil {
		return
	}
	return bash.GoImports(target)
}
