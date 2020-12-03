package typmock

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/typical-go/typical-go/pkg/oskit"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

var (
	// MockTag is tag for mock
	MockTag = "@mock"
)

type (
	// GenerateMock mock
	GenerateMock struct {
		Includes []string
		Excludes []string
	}
)

var _ typgo.Tasker = (*GenerateMock)(nil)
var _ typgo.Action = (*GenerateMock)(nil)

// Task to mock
func (d *GenerateMock) Task(c *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:        "mock",
		Usage:       "Generate mock class",
		UsageText:   "mock [package_names]",
		Description: "If package_names is missing then check every package",
		Action:      c.Action(d),
	}
}

// Execute mock command
func (d *GenerateMock) Execute(c *typgo.Context) error {
	dirs, err := filekit.FindDir(d.Includes, d.Excludes)
	if err != nil {
		return err
	}
	_, files := typast.Walk(dirs)
	summary, err := typast.Compile(files...)
	if err != nil {
		return err
	}
	return Annotate(c, summary)
}

// Annotate mock
func Annotate(c *typgo.Context, summary *typast.Summary) error {
	mockgen := fmt.Sprintf("%s/bin/mockgen", typgo.TypicalTmp)
	if _, err := os.Stat(mockgen); os.IsNotExist(err) {
		err := c.Execute(&typgo.GoBuild{Output: mockgen, MainPackage: "github.com/golang/mock/mockgen"})
		if err != nil {
			return err
		}
	}
	mockery := NewMockery(typgo.ProjectPkg)

	for _, annot := range summary.FindAnnot(MockTag, typast.EqualInterface) {

		mockery.Put(CreateMock(annot))
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
				fmt.Fprintf(oskit.Stdout, "Fail to mock '%s': %s\n", name, err.Error())
			}
		}
	}
	return nil
}