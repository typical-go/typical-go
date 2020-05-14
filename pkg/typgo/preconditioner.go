package typgo

import (
	"context"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
)

type (
	// Preconditioner responsible to precondition
	Preconditioner interface {
		Precondition(c *PrecondContext) error
	}

	// PrecondContext is context of preconditioning
	PrecondContext struct {
		*Descriptor
		typtmpl.Precond
		typlog.Logger
		Ctx      context.Context
		ASTStore *typast.ASTStore
	}
)

func createPrecondContext(ctx context.Context, d *Descriptor) *PrecondContext {
	var (
		err      error
		astStore *typast.ASTStore
	)

	logger := typlog.Logger{
		Name:  "PRECOND",
		Color: typlog.DefaultColor,
	}

	appDirs, appFiles := WalkLayout(d.Layouts)

	if astStore, err = typast.CreateASTStore(appFiles...); err != nil {
		logger.Warn(err.Error())
	}

	return &PrecondContext{
		Precond: typtmpl.Precond{
			Imports: retrImports(appDirs),
			Package: "main",
		},
		Logger:     logger,
		Descriptor: d,
		Ctx:        ctx,
		ASTStore:   astStore,
	}
}

func retrImports(dirs []string) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typgo",
	}
	for _, dir := range dirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", typvar.ProjectPkg, dir))
		}
	}
	return imports
}
