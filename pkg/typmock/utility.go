package typmock

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typvar"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

// Utility to generate mock class
func Utility() typgo.Utility {
	return typgo.NewUtility(commands)
}

func commands(c *typgo.BuildTool) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "mock",
			Usage:       "Generate mock class",
			UsageText:   "mock [package_names]",
			Description: "If package_names is missing then check every package",
			Action:      c.ActionFunc("mock", generateMock),
		},
	}
}

func generateMock(c *typgo.Context) (err error) {
	var (
		store *typast.ASTStore
	)

	_, files := typgo.WalkLayout(c.BuildTool.Layouts)
	if store, err = typast.CreateASTStore(files...); err != nil {
		return
	}

	mockery := NewMockery(typvar.ProjectPkg)

	mocks := typannot.GetMock(store)
	for _, mock := range mocks {
		mockery.Put(mock)
	}

	targetMap := mockery.TargetMap(c.Cli.Args().Slice()...)
	if len(targetMap) < 1 {
		return
	}

	mockgen := fmt.Sprintf("%s/bin/mockgen", typvar.TypicalTmp)
	if err = installIfNotExist(c.Cli.Context, mockgen); err != nil {
		return
	}

	for pkg, targets := range targetMap {
		mockPkg := fmt.Sprintf("%s_mock", pkg)

		fmt.Printf("\nRemove package '%s'\n", mockPkg)
		os.RemoveAll(mockPkg)

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", typvar.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, mockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			cmd := &execkit.Command{
				Name: mockgen,
				Args: []string{
					"-destination", dest,
					"-package", mockPkg,
					srcPkg,
					t.Source,
				},
				Stderr: os.Stderr,
			}

			cmd.Print(os.Stdout)

			if err = cmd.Run(c.Cli.Context); err != nil {
				c.Warnf("Fail to mock '%s': %s", name, err.Error())
			}
		}
	}
	return
}

func installIfNotExist(ctx context.Context, mockgen string) (err error) {
	if _, err = os.Stat(mockgen); os.IsNotExist(err) {
		cmd := buildkit.
			NewGoBuild(mockgen, "github.com/golang/mock/mockgen").
			Command()
		return cmd.Run(ctx)
	}
	return
}
