package typast

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

type (
	// AnnotateCmd annotate cmd
	AnnotateCmd struct {
		Destination string // By default is "internal/generated/typical"
		Annotators  []Annotator
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
		Dirs        []string
		Destination string
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
		Name:    "annotate",
		Aliases: []string{"a"},
		Usage:   "Annotate the project and generate code",
		Action:  sys.Action(a),
	}
}

// Execute annotation
func (a *AnnotateCmd) Execute(c *typgo.Context) error {
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
func (a *AnnotateCmd) CreateContext(c *typgo.Context) (*Context, error) {
	dirs, files := Walk(c.BuildSys.AppLayouts)
	summary, err := Compile(files...)
	if err != nil {
		return nil, err
	}
	destination := a.getDestination()
	os.MkdirAll(destination, 0777)
	return &Context{
		Context:     c,
		Summary:     summary,
		Dirs:        dirs,
		Destination: destination,
	}, nil
}

func (a *AnnotateCmd) getDestination() string {
	if a.Destination == "" {
		a.Destination = "internal/generated/typical"
	}
	return a.Destination
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
