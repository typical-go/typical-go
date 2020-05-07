package typmock

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typast"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/urfave/cli/v2"
)

// Utility to generate mock class
func Utility() typbuild.Utility {
	return typbuild.NewUtility(commands)
}

func commands(c *typbuild.Context) []*cli.Command {
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

func generateMock(c *typbuild.CliContext) (err error) {
	var (
		store *typast.ASTStore
	)

	if store, err = typast.CreateASTStore(c.Core.AppFiles...); err != nil {
		return
	}

	mockery := NewMockery(c.Core.ProjectPkg)

	mocks := typannot.GetMock(store)
	for _, mock := range mocks {
		mockery.Put(mock)
	}

	targetMap := mockery.TargetMap(c.Cli.Args().Slice()...)
	if len(targetMap) < 1 {
		return
	}

	mockgen := fmt.Sprintf("%s/bin/mockgen", c.Core.TypicalTmp)
	if err = installIfNotExist(c.Context, mockgen); err != nil {
		return
	}

	for pkg, targets := range targetMap {
		mockPkg := fmt.Sprintf("mock_%s", pkg)

		fmt.Printf("\nRemove package '%s'\n", mockPkg)
		os.RemoveAll(mockPkg)

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", c.Core.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, mockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			cmd := &buildkit.Command{
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

			if err = cmd.Run(c.Context); err != nil {
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
