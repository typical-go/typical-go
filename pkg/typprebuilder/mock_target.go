package typprebuilder

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"
	"github.com/typical-go/typical-go/pkg/utility/golang"
)

type mockTarget struct {
	ApplicationImports golang.Imports
	MockTargets        []string
}

func (g mockTarget) generate(target string) (err error) {
	defer debugkit.ElapsedTime("Generate mock target")()
	src := golang.NewSourceCode(typenv.Dependency.Package)
	src.Imports = g.ApplicationImports
	for _, mockTarget := range g.MockTargets {
		src.Init.Append(fmt.Sprintf("typical.Context.MockTargets.Append(\"%s\")", mockTarget))
	}
	if err = src.WriteToFile(target); err != nil {
		return
	}
	return bash.GoImports(target)
}
