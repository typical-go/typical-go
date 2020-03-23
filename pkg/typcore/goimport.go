package typcore

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

// GoImport to execute goimports to the target
func GoImport(ctx context.Context, c *Context, target string) (err error) {
	goimports := fmt.Sprintf("%s/bin/goimports", c.TmpFolder)

	if _, err = os.Stat(goimports); os.IsNotExist(err) {
		c.Infof("Install goimports")
		if err = buildkit.NewGoBuild(goimports, "golang.org/x/tools/cmd/goimports").Execute(ctx); err != nil {
			return
		}
	}

	cmd := exec.CommandContext(ctx, goimports, "-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
