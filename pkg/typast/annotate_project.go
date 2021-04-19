package typast

import (
	"errors"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// AnnotateProject task
	AnnotateProject struct {
		Sources    []string
		Annotators []Annotator
	}
	// Annotator responsible to annotate
	Annotator interface {
		Annotate(*Context) error
	}
	// AnnotateFn annotate function
	AnnotateFn    func(*Context) error
	annotatorImpl struct {
		fn AnnotateFn
	}
	// Context of annotation
	Context struct {
		*typgo.Context
		*Summary
		Dirs []string
	}
)

//
// AnnotateProject
//

var _ typgo.Tasker = (*AnnotateProject)(nil)
var _ typgo.Action = (*AnnotateProject)(nil)

// Task to annotate
func (a AnnotateProject) Task() *typgo.Task {
	return &typgo.Task{
		Name:    "annotate",
		Aliases: []string{"a"},
		Usage:   "Annotate the project and generate code",
		Action:  a,
	}
}

// Execute annotation
func (a AnnotateProject) Execute(c *typgo.Context) error {
	if len(a.Sources) == 0 {
		return errors.New("'Sources' is missing")
	}
	ac, err := a.CreateContext(c)
	if err != nil {
		return err
	}
	for _, annotator := range a.Annotators {
		if err := annotator.Annotate(ac); err != nil {
			return err
		}
	}
	return nil
}

// CreateContext create context
func (a AnnotateProject) CreateContext(c *typgo.Context) (*Context, error) {
	dirs, files := Walk(a.Sources)
	summary, err := Compile(files...)
	if err != nil {
		return nil, err
	}

	return &Context{
		Context: c,
		Summary: summary,
		Dirs:    dirs,
	}, nil
}

//
// annotatorImpl
//

// NewAnnotator return new instance of annotator
func NewAnnotator(fn AnnotateFn) Annotator {
	return &annotatorImpl{fn: fn}
}

func (a *annotatorImpl) Annotate(c *Context) (err error) {
	return a.fn(c)
}
