package typgo

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
)

// GoImports target file
func GoImports(ctx context.Context, target string) (err error) {
	goimports := fmt.Sprintf("%s/bin/goimports", TypicalTmp)

	if _, err = os.Stat(goimports); os.IsNotExist(err) {
		execute(ctx, &execkit.GoBuild{
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
