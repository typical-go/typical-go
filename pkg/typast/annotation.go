package typast

import (
	"errors"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Annotation struct {
		Filter    Filter
		Processor Processor
	}
	Processor interface {
		Process(*typgo.Context, Directives) error
	}
	NewProcessor func(*typgo.Context, Directives) error
)

//
// Annotation
//

var _ Processor = (*Annotation)(nil)
var _ Annotator = (*Annotation)(nil)

func (a *Annotation) Process(c *typgo.Context, directives Directives) error {
	if a.Processor == nil {
		return errors.New("mising annotation processor")
	}
	var filtered []*Directive
	for _, dir := range directives {
		if a.isAllowed(dir) {
			filtered = append(filtered, dir)
		}
	}
	return a.Processor.Process(c, filtered)
}

func (a *Annotation) Annotate() Processor {
	return a
}

func (a *Annotation) isAllowed(d *Directive) bool {
	if a.Filter == nil {
		return true
	}
	return a.Filter.IsAllowed(d)
}

//
// NewProcessor
//

var _ Processor = (NewProcessor)(nil)

func (p NewProcessor) Process(c *typgo.Context, d Directives) error {
	return p(c, d)
}
