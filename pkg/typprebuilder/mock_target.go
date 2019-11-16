package typprebuilder

import (
	"github.com/typical-go/typical-go/pkg/typprebuilder/golang"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"
)

type mockTarget struct {
	ApplicationImports golang.Imports
	MockTargets        []string
}

func (g mockTarget) generate(target string) (err error) {
	defer debugkit.ElapsedTime("Generate mock target")()
	src := golang.NewSourceCode(typenv.Dependency.Package)
	src.Imports = g.ApplicationImports
	src.AddMockTargets(g.MockTargets...)
	if err = src.Cook(target); err != nil {
		return
	}
	return bash.GoImports(target)
}
