package typannot

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

type (
	// AnnotateCmd annotate cmd
	AnnotateCmd struct {
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
)

//
// AnnotateCmd
//

var _ typgo.Cmd = (*AnnotateCmd)(nil)
var _ typgo.Action = (*AnnotateCmd)(nil)

// Command annotate
func (a *AnnotateCmd) Command(sys *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:   "annotate",
		Usage:  "Annotate the project and generate code",
		Action: sys.ActionFn(a),
	}
}

// Execute annotation
func (a *AnnotateCmd) Execute(c *typgo.Context) error {
	ac, err := CreateContext(c)
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
