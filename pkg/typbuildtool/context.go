package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Context of buildtool
type Context struct {
	*typcore.Context
	*TypicalBuildTool
	Cli *cli.Context

	ast *typast.Ast
}

// Ast contain detail of AST analysis
func (c *Context) Ast() *typast.Ast {
	if c.ast == nil {
		var err error
		if c.ast, err = typast.Walk(c.ProjectFiles); err != nil {
			c.Errorf("PreconditionContext: %w", err.Error())
		}
	}
	return c.ast
}
