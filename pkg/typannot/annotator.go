package typannot

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Annotators is extra information from source code
	Annotators []Annotator
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

var _ typgo.Action = (Annotators)(nil)

// Execute annotation
func (a Annotators) Execute(c *typgo.Context) error {
	ac, err := CreateContext(c)
	if err != nil {
		return err
	}
	for _, annotator := range a {
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
