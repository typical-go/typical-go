package typgo

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typlog"
)

type (
	// Prebuilder return new instance of prebuilder
	Prebuilder interface {
		Prebuild(*PrebuildContext) error
	}
	// PrebuildContext prebuild context
	PrebuildContext struct {
		typlog.Logger
		*BuildCli
		ctx context.Context
	}
	// Prebuilds composite prebuild
	Prebuilds []Prebuilder
	// PrebuildFn function
	PrebuildFn     func(*PrebuildContext) error
	prebuilderImpl struct {
		fn PrebuildFn
	}
)

//
// prebuilderImpl
//

// NewPrebuild return new instance of Prebuilder
func NewPrebuild(fn PrebuildFn) Prebuilder {
	return &prebuilderImpl{fn: fn}
}

func (p *prebuilderImpl) Prebuild(c *PrebuildContext) error {
	return p.fn(c)
}

//
// Prebuilds
//

var _ Prebuilder = (Prebuilds)(nil)

// Prebuild prebuilds
func (p Prebuilds) Prebuild(c *PrebuildContext) error {
	for _, prebuild := range p {
		if err := prebuild.Prebuild(c); err != nil {
			return err
		}
	}
	return nil
}

//
// PrebuildContext
//

// Ctx return golang context
func (c *PrebuildContext) Ctx() context.Context {
	return c.ctx
}
