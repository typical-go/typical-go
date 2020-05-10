package typgo

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

type (
	// Preconditioner responsible to precondition
	Preconditioner interface {
		Precondition(c *PrecondContext) error
	}

	// PrecondContext is context of preconditioning
	PrecondContext struct {
		*BuildTool
		typtmpl.Precond
		typlog.Logger
		Ctx      context.Context
		astStore *typast.ASTStore
	}
)

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
