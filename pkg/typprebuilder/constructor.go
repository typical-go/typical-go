package typprebuilder

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"

	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/golang"
)

type constructor struct {
	ApplicationImports golang.Imports
	Constructors       []string
}

func (g constructor) generate(target string) (err error) {
	defer debugkit.ElapsedTime("Generate constructor")()
	src := golang.NewSourceCode(typenv.Dependency.Package)
	src.Imports = g.ApplicationImports
	for _, constructor := range g.Constructors {
		src.Init.Append(fmt.Sprintf("typical.Context.Constructors.Append(%s)", constructor))
	}
	if err = src.Cook(target); err != nil {
		return
	}
	return bash.GoImports(target)
}
