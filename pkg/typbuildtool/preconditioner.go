package typbuildtool

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typfactory"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *PreconditionContext) error
}

// PreconditionContext is context of preconditioning
type PreconditionContext struct {
	typfactory.Precond
	Core     *typcore.Context
	Ctx      context.Context
	astStore *typast.ASTStore
}

// ASTStore return the ast store
func (c *PreconditionContext) ASTStore() *typast.ASTStore {
	var err error
	if c.astStore == nil {
		c.astStore, err = typast.CreateASTStore(c.Core.AppFiles...)
		if err != nil {
			c.Warn(err.Error())
		}
	}
	return c.astStore
}

// Info logger
func (c *PreconditionContext) Info(args ...interface{}) {
	c.Core.Info(args...) // TODO: precondition label
}

// Infof logger
func (c *PreconditionContext) Infof(format string, args ...interface{}) {
	c.Core.Infof(format, args) // TODO: precondition label
}

// Warn logger
func (c *PreconditionContext) Warn(args ...interface{}) {
	c.Core.Warn(args...) // TODO: precondition label
}

// Warnf logger
func (c *PreconditionContext) Warnf(format string, args ...interface{}) {
	c.Core.Warnf(format, args...) // TODO: precondition label
}
