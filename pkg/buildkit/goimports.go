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
		gobuild := &GoBuild{
			Out:    goimports,
			Source: "golang.org/x/tools/cmd/goimports",
		}
		if err = gobuild.Command().Run(ctx); err != nil {
			return
		}
	}

	cmd := exec.CommandContext(ctx, goimports, "-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
