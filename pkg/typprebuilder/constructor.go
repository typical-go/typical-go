package typprebuilder

import (
	"fmt"

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
	src := golang.NewSource("main")
	src.Imports = g.ApplicationImports
	if len(g.Constructors) < 1 {
		return
	}
	for _, constructor := range g.Constructors {
		src.Init.Append(fmt.Sprintf("typical.Context.Constructors.Append(%s)", constructor))
	}
	if err = filekit.Write(target, src); err != nil {
		return
	}
	return goimports(target)
}
