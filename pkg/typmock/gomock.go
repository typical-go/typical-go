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
	DefaultMockAnnot = "@mock"

	_ typgo.Tasker     = (*GoMock)(nil)
	_ typgen.Annotator = (*GoMock)(nil)
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
	annots, err := typgen.Compile(filePaths...)
	if err != nil {
		return err
	}
	annots = typgen.Filter(annots, g)
	ctx := &typgen.Context{Context: c}
	for _, annot := range annots {
		if err := g.ProcessAnnot(ctx, annot); err != nil {
			return err
		}
	}

	return nil
}

func (g *GoMock) walk() []string {
	if g.Walker == nil {
		return typgen.Layouts{"internal"}.Walk()
	}
	return g.Walker.Walk()
}

func (g *GoMock) ProcessAnnot(c *typgen.Context, annot *typgen.Annotation) error {
	mock := CreateMock(annot)
	if err := mock.Generate(c.Context); err != nil {
		c.Infof("Fail to mock '%s': %s\n", mock.Package, err.Error())
	}
	return nil
}

func (g *GoMock) AnnotationName() string {
	return DefaultMockAnnot
}

func (g *GoMock) IsAllowed(d *typgen.Annotation) bool {
	return typgen.IsPublic(d) && typgen.IsInterface(d)
}
