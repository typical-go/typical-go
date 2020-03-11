package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *PreconditionContext) error
}

// PreconditionContext is context of Precondition
type PreconditionContext struct {
	*typcore.Context
	Cli *cli.Context
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
