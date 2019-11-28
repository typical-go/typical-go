package typprebuilder

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/typical-go/typical-go/pkg/utility/coll"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"
	"github.com/typical-go/typical-go/pkg/utility/filekit"
	"github.com/typical-go/typical-go/pkg/utility/golang"
)

type constructor struct {
	ApplicationImports coll.KeyStrings
	Constructors       []string
}

func (g constructor) generate(target string) (err error) {
	defer debugkit.ElapsedTime("Generate constructor")()
	src := golang.NewSource(typenv.Dependency)
	src.Imports = g.ApplicationImports
	for _, constructor := range g.Constructors {
		src.Init.Append(fmt.Sprintf("typical.Context.Constructors.Append(%s)", constructor))
	}
	if err = filekit.Write(target, src); err != nil {
		return
	}
	return bash.GoImports(target)
}
