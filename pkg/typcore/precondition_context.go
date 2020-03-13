package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typast"
)

// PreconditionContext is context of Precondition
type PreconditionContext struct {
	*Context
	ast *typast.Ast
}

// Ast contain detail of AST analysis
func (c *PreconditionContext) Ast() *typast.Ast {
	if c.ast == nil {
		var err error
		if c.ast, err = typast.Walk(c.ProjectFiles); err != nil {
			log.Errorf("PreconditionContext: %w", err.Error())
		}
	}
	return c.ast
}
