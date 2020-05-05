package typbuildtool

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *PreconditionContext) error
}

// PreconditionContext is context of preconditioning
type PreconditionContext struct {
	typtmpl.Precond
	Logger   typlog.Logger
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

// SetASTStore to set ast store
func (c *PreconditionContext) SetASTStore(astStore *typast.ASTStore) {
	c.astStore = astStore
}

// Info logger
func (c *PreconditionContext) Info(args ...interface{}) {
	c.Logger.Info(args...)
}

// Infof logger
func (c *PreconditionContext) Infof(format string, args ...interface{}) {
	c.Logger.Infof(format, args...)
}

// Warn logger
func (c *PreconditionContext) Warn(args ...interface{}) {
	c.Logger.Warn(args...)
}

// Warnf logger
func (c *PreconditionContext) Warnf(format string, args ...interface{}) {
	c.Logger.Warnf(format, args...)
}
