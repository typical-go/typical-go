package typmock

import (
	"errors"

	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	GoMock struct {
		Layout typgen.Layouts
		Walker typgen.Walker
	}
)

var (
	DefaultMockTag = "@mock"

	_ typgo.Tasker      = (*GoMock)(nil)
	_ typgen.Annotation = (*GoMock)(nil)
)

// Task to mock
func (d *GoMock) Task() *typgo.Task {
	return &typgo.Task{
		Name:   "mock",
		Usage:  "Generate mock class",
		Action: d,
	}
}

func (g *GoMock) Execute(c *typgo.Context) error {
	filePaths := g.walk()
	if len(filePaths) < 1 {
		return errors.New("walker couldn't find any filepath")
	}
	dirs, err := typgen.Compile(filePaths...)
	if err != nil {
		return err
	}
	if err := typgen.ExecuteAnnotation(c, g, dirs); err != nil {
		return err
	}

	return nil
}

func (g *GoMock) walk() []string {
	if g.Walker == nil {
		return typgen.Layouts{"internal"}.Walk()
	}
	return g.Walker.Walk()
}

func (g *GoMock) Process(c *typgen.Context) error {
	var mocks Mocks
	for _, annot := range c.Dirs {
		mocks = append(mocks, CreateMock(annot))
	}

	for _, t := range mocks {
		err := MockGen(c.Context, t.MockPkg, t.Dest, t.Pkg, t.Source)
		if err != nil {
			c.Infof("Fail to mock '%s': %s\n", t.Pkg, err.Error())
		}
	}
	return nil
}

func (g *GoMock) TagName() string {
	return DefaultMockTag
}

func (g *GoMock) IsAllowed(d *typgen.Directive) bool {
	return typgen.IsPublic(d) && typgen.IsInterface(d)
}
