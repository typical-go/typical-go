package typgo

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typvar"
)

// GoImports target file
func GoImports(ctx context.Context, target string) (err error) {
	goimports := fmt.Sprintf("%s/bin/goimports", typvar.TypicalTmp)

	if _, err = os.Stat(goimports); os.IsNotExist(err) {
		execute(ctx, &buildkit.GoBuild{
			Out:    goimports,
			Source: "golang.org/x/tools/cmd/goimports",
		})
	}

	return execute(ctx, &execkit.Command{
		Name:   goimports,
		Args:   []string{"-w", target},
		Stderr: os.Stderr,
	})
}
