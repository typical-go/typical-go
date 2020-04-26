package typmock

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/buildkit"
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
			Name:        "mock",
			Usage:       "Generate mock class",
			UsageText:   "mock [package_names]",
			Description: "If package_names is missing then check every package",
			Action: func(cliCtx *cli.Context) (err error) {
				return generateMock(c.CliContext(cliCtx))
			},
		},
	}
}

func generateMock(c *typbuildtool.CliContext) (err error) {
	ctx := c.Cli.Context
	mockery := NewMockery(c.ProjectPkg)

	if err = mockery.Walk(c.Ast()); err != nil {
		return
	}

	targetMap := mockery.TargetMap(c.Cli.Args().Slice()...)

	if len(targetMap) < 1 {
		c.Info("Nothing to mock")
		return
	}

	mockgen := fmt.Sprintf("%s/bin/mockgen", c.TypicalTmp)

	if err = installIfNotExist(ctx, mockgen); err != nil {
		return
	}

	for pkg, targets := range targetMap {
		mockPkg := fmt.Sprintf("mock_%s", pkg)

		c.Infof("Remove package: %s", mockPkg)
		os.RemoveAll(mockPkg)

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", c.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, mockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			cmd := exec.CommandContext(ctx, mockgen,
				"-destination", dest,
				"-package", mockPkg,
				srcPkg,
				t.Source,
			)
			cmd.Stderr = os.Stderr

			c.Infof("Mock '%s'", name)
			if err = cmd.Run(); err != nil {
				c.Warnf("Fail to mock '%s': %w", name, err)
			}
		}
	}
	return
}

func installIfNotExist(ctx context.Context, mockgen string) (err error) {
	if _, err = os.Stat(mockgen); os.IsNotExist(err) {
		return buildkit.
			NewGoBuild(mockgen, "github.com/golang/mock/mockgen").
			Execute(ctx)
	}
	return
}
