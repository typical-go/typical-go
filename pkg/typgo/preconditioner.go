package typgo

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *PrecondContext) error
}

// PrecondContext is context of preconditioning
type PrecondContext struct {
	*Context
	typtmpl.Precond
	Logger   typlog.Logger
	Ctx      context.Context
	astStore *typast.ASTStore
}

// ASTStore return the ast store
func (c *PrecondContext) ASTStore() *typast.ASTStore {
	var err error
	if c.astStore == nil {
		c.astStore, err = typast.CreateASTStore(c.AppFiles...)
		if err != nil {
			c.Warn(err.Error())
		}
	}
	return c.astStore
}

// SetASTStore to set ast store
func (c *PrecondContext) SetASTStore(astStore *typast.ASTStore) {
	c.astStore = astStore
}

// Info logger
func (c *PrecondContext) Info(args ...interface{}) {
	c.Logger.Info(args...)
}

// Infof logger
func (c *PrecondContext) Infof(format string, args ...interface{}) {
	c.Logger.Infof(format, args...)
}

// Warn logger
func (c *PrecondContext) Warn(args ...interface{}) {
	c.Logger.Warn(args...)
}

// Warnf logger
func (c *PrecondContext) Warnf(format string, args ...interface{}) {
	c.Logger.Warnf(format, args...)
}
