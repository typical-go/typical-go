package typprebuilder

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/debugkit"
	"github.com/typical-go/typical-go/pkg/utility/filekit"
	"github.com/typical-go/typical-go/pkg/utility/golang"
)

type constructor struct {
	Constructors   []string
	ProjectPackage string
}

func (g constructor) generate(target string) (err error) {
	defer debugkit.ElapsedTime("Generate constructor")()
	src := golang.NewSource("main")
	if len(g.Constructors) < 1 {
		return
	}
	imports := make(map[string]struct{})
	imports[g.ProjectPackage+"/typical"] = struct{}{}
	for _, constructor := range g.Constructors {
		dotIndex := strings.Index(constructor, ".")
		if dotIndex >= 0 {
			pkg := constructor[:dotIndex]
			imports[g.ProjectPackage+"/"+pkg] = struct{}{}
		}
		src.Init.Append(fmt.Sprintf("typical.Context.Constructors.Append(%s)", constructor))
	}
	for key := range imports {
		src.Imports.Add("", key)
	}
	if err = filekit.Write(target, src); err != nil {
		return
	}
	return goimports(target)
}
