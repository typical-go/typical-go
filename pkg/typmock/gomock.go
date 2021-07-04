package typmock

import (
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// MockTag is tag for mock
	MockTag = "@mock"
)

type (
	// GoMock mock
	GoMock struct {
		Walker typgen.Walker
	}
)

var _ typgo.Tasker = (*GoMock)(nil)
var _ typgen.Processor = (*GoMock)(nil)

// Task to mock
func (d *GoMock) Task() *typgo.Task {
	return &typgo.Task{
		Name:  "mock",
		Usage: "Generate mock class",
		Action: &typgen.Generator{
			Walker:    d.Walker,
			Processor: d,
		},
	}
}
func (d *GoMock) Process(c *typgo.Context, directives typgen.Directives) error {
	return d.Annotation().Process(c, directives)
}

func (d *GoMock) Annotation() *typgen.Annotation {
	return &typgen.Annotation{
		Filter: typgen.Filters{
			&typgen.TagNameFilter{MockTag},
			&typgen.PublicFilter{},
			&typgen.InterfaceFilter{},
		},
		ProcessFn: d.process,
	}
}

func (d *GoMock) process(c *typgo.Context, directives typgen.Directives) error {
	var mocks Mocks

	for _, annot := range directives {
		mocks = append(mocks, CreateMock(annot))
	}

	for _, t := range mocks {
		err := MockGen(c, t.MockPkg, t.Dest, t.Pkg, t.Source)
		if err != nil {
			c.Infof("Fail to mock '%s': %s\n", t.Pkg, err.Error())
		}
	}

	return nil
}
