package typmock

import (
	"fmt"
	"log"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

type (
	// MockCmd mock
	MockCmd struct{}
)

var _ typgo.Cmd = (*MockCmd)(nil)

// Command to utility
func (d *MockCmd) Command(c *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:        "mock",
		Usage:       "Generate mock class",
		UsageText:   "mock [package_names]",
		Description: "If package_names is missing then check every package",
		Action: c.ActionFn(func(c *typgo.Context) error {
			ac, err := typannot.CreateContext(c)
			if err != nil {
				return err
			}
			return Annotate(ac)
		}),
	}
}

// Annotate mock
func Annotate(c *typannot.Context) error {
	mockgen := fmt.Sprintf("%s/bin/mockgen", typgo.TypicalTmp)
	if _, err := os.Stat(mockgen); os.IsNotExist(err) {
		if err := c.Execute(&execkit.GoBuild{
			Output:      mockgen,
			MainPackage: "github.com/golang/mock/mockgen",
		}); err != nil {
			return err
		}
	}

	mockery := NewMockery(typgo.ProjectPkg)
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckInterface(MockTag) {
			mockery.Put(CreateMock(annot))
		}
	}
	targetMap := mockery.Map
	args := c.Args()
	if args.Len() > 0 {
		targetMap = mockery.Filter(args.Slice()...)
	}

	for key, targets := range targetMap {
		mockPkg := fmt.Sprintf("%s_mock", key)

		c.Execute(&execkit.Command{Name: "rm", Args: []string{"-rf", mockPkg}})

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", typgo.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, t.MockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			if err := c.Execute(&execkit.Command{
				Name: mockgen,
				Args: []string{
					"-destination", dest,
					"-package", t.MockPkg,
					srcPkg,
					t.Source,
				},
				Stderr: os.Stderr,
			}); err != nil {
				log.Printf("Fail to mock '%s': %s", name, err.Error())
			}
		}
	}
	return nil
}
