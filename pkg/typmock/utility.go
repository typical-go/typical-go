package typmock

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typvar"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

// Utility to generate mock class
type Utility struct{}

var _ typgo.Utility = (*Utility)(nil)

// Commands to utility
func (*Utility) Commands(c *typgo.BuildCli) ([]*cli.Command, error) {
	return []*cli.Command{
		{
			Name:        "mock",
			Usage:       "Generate mock class",
			UsageText:   "mock [package_names]",
			Description: "If package_names is missing then check every package",
			Action:      c.ActionFn("mock", mock),
		},
	}, nil
}

func mock(c *typgo.Context) (err error) {

	mockery := createMockery(c)

	mockgen := fmt.Sprintf("%s/bin/mockgen", typvar.TypicalTmp)
	if err = installIfNotExist(c, mockgen); err != nil {
		return
	}

	targetMap := mockery.TargetMap

	if c.Args().Len() > 0 {
		targetMap = mockery.TargetMap.Filter(c.Args().Slice()...)
	}

	for key, targets := range targetMap {
		mockPkg := fmt.Sprintf("%s_mock", key)

		fmt.Printf("\nRemove all: %s\n", mockPkg)
		os.RemoveAll(mockPkg)

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", typvar.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, t.MockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			if err = c.Execute(&execkit.Command{
				Name: mockgen,
				Args: []string{
					"-destination", dest,
					"-package", t.MockPkg,
					srcPkg,
					t.Source,
				},
				Stderr: os.Stderr,
			}); err != nil {
				c.Warnf("Fail to mock '%s': %s", name, err.Error())
			}
		}
	}
	return
}

func installIfNotExist(c *typgo.Context, mockgen string) (err error) {
	if _, err = os.Stat(mockgen); os.IsNotExist(err) {
		return c.Execute(&execkit.GoBuild{
			Out:    mockgen,
			Source: "github.com/golang/mock/mockgen",
		})
	}
	return
}
