package typmock

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

// Utility to generate mock class
func Utility() typbuildtool.Utility {
	return typbuildtool.NewUtility(commands)
}

func commands(c *typbuildtool.Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Action: func(cliCtx *cli.Context) (err error) {
				return mock(c.BuildContext(cliCtx))
			},
		},
	}
}

func mock(c *typbuildtool.BuildContext) (err error) {
	ctx := c.Cli.Context
	store := NewMockStore()
	if err = c.Ast().EachAnnotation("mock", typast.InterfaceType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		store.Put(buildkit.CreateGoMock(c.TypicalTmp, c.ProjectPkg, decl))
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
