package buildkit

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// GoImports utility
type GoImports struct {
	typicalTmp string
	target     string
}

// NewGoImports return new instance of GoImports
func NewGoImports(typicalTmp, target string) *GoImports {
	return &GoImports{
		typicalTmp: typicalTmp,
		target:     target,
	}
}

// Execute goimports
func (g *GoImports) Execute(ctx context.Context) (err error) {
	goimports := fmt.Sprintf("%s/bin/goimports", g.typicalTmp)

	if _, err = os.Stat(goimports); os.IsNotExist(err) {
		if err = NewGoBuild(goimports, "golang.org/x/tools/cmd/goimports").Execute(ctx); err != nil {
			return
		}
	}

	cmd := exec.CommandContext(ctx, goimports, "-w", g.target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
