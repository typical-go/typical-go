package typgen

import (
	"errors"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Annotation struct {
		Filter    Filter
		ProcessFn ProcessFn
	}
	Processor interface {
		Process(*typgo.Context, []*Directive) error
	}
	ProcessFn  func(*typgo.Context, []*Directive) error
	Processors []Processor
)

//
// Annotation
//

var _ Processor = (*Annotation)(nil)

func (a *Annotation) Process(c *typgo.Context, directives []*Directive) error {
	if a.ProcessFn == nil {
		return errors.New("mising annotation processor")
	}
	var filtered []*Directive
	for _, dir := range directives {
		if a.isAllowed(dir) {
			filtered = append(filtered, dir)
		}
	}
	return a.ProcessFn(c, filtered)
}

func (a *Annotation) isAllowed(d *Directive) bool {
	if a.Filter == nil {
		return true
	}
	return a.Filter.IsAllowed(d)
}

//
// ProcessFn
//

var _ Processor = (ProcessFn)(nil)

func (p ProcessFn) Process(c *typgo.Context, d []*Directive) error {
	return p(c, d)
}

//
// Processor
//

var _ Processor = (Processors)(nil)

func (p Processors) Process(c *typgo.Context, d []*Directive) error {
	for _, processor := range p {
		if err := processor.Process(c, d); err != nil {
			return err
		}
	}
	return nil
}
