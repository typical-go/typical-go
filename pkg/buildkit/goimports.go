package buildkit

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typvar"
)

// GoImports target file
func GoImports(ctx context.Context, target string) (err error) {
	goimports := fmt.Sprintf("%s/bin/goimports", typvar.TypicalTmp)

	if _, err = os.Stat(goimports); os.IsNotExist(err) {
		if err = NewGoBuild(goimports, "golang.org/x/tools/cmd/goimports").Command().Run(ctx); err != nil {
			return
		}
	}

	cmd := exec.CommandContext(ctx, goimports, "-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
