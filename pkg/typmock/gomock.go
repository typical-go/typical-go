package typmock

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// MockTag is tag for mock
	MockTag = "@mock"
)

type (
	// GoMock mock
	GoMock struct {
		Sources []string
	}
)

var _ typgo.Tasker = (*GoMock)(nil)
var _ typgo.Action = (*GoMock)(nil)

// Task to mock
func (d *GoMock) Task() *typgo.Task {
	return &typgo.Task{
		Name:   "mock",
		Usage:  "Generate mock class",
		Action: d,
	}
}

// Execute mock command
func (d *GoMock) Execute(c *typgo.Context) error {
	_, files := typast.Walk(d.Sources)
	summary, err := typast.Compile(files...)
	if err != nil {
		return err
	}
	return Annotate(c, summary)
}

// Annotate mock
func Annotate(c *typgo.Context, summary *typast.Summary) error {
	mockery := NewMockery(typgo.ProjectPkg)

	for _, annot := range summary.Annots {
		if strings.EqualFold(annot.TagName, MockTag) && typast.IsInterface(annot) {
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

		c.Execute(&typgo.Bash{Name: "rm", Args: []string{"-rf", mockPkg}})

		for _, t := range targets {
			srcPkg := fmt.Sprintf("%s/%s", typgo.ProjectPkg, t.Dir)
			dest := fmt.Sprintf("%s%s/%s.go", t.Parent, t.MockPkg, strcase.ToSnake(t.Source))
			name := fmt.Sprintf("%s.%s", srcPkg, t.Source)

			err := MockGen(c, t.MockPkg, dest, srcPkg, t.Source)
			if err != nil {
				c.Infof("Fail to mock '%s': %s\n", name, err.Error())
			}
		}
	}
	return nil
}

// MockGen execute mockgen bash
func MockGen(c *typgo.Context, destPkg, dest, srcPkg, src string) error {
	mockgen, err := typgo.InstallTool(c.Ctx(), "mockgen", "github.com/golang/mock/mockgen")
	if err != nil {
		return err
	}

	return c.Execute(&typgo.Bash{
		Name: mockgen,
		Args: []string{
			"-destination", dest,
			"-package", destPkg,
			srcPkg,
			src,
		},
		Stderr: os.Stderr,
	})
}
