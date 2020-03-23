package typbuildtool

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
)

// Mock the project
func (b *StandardModule) Mock(c *BuildContext) (err error) {
	ctx := c.Cli.Context
	store := NewMockStore()
	if err = c.Ast().EachAnnotation("mock", typast.InterfaceType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		store.Put(buildkit.CreateGoMock(c.TypicalTmp, c.ProjectPackage, decl))
		return
	}); err != nil {
		return
	}

	for pkg, targets := range store.Map() {

		c.Infof("Remove package: %s", pkg)
		os.RemoveAll(pkg)

		for _, target := range targets {
			c.Infof("Mock '%s'", target)
			if err = target.Execute(ctx); err != nil {
				c.Warnf("Fail to mock '%s': %s", target, err.Error())
			}
		}
	}

	return
}
