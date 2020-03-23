package typbuildtool

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Build the project
func (b *StandardModule) Build(c *BuildContext) (dists []BuildDistribution, err error) {
	c.Info("Standard-Build: Build the project")
	binary := fmt.Sprintf("%s/%s", c.BinFolder(), c.Name)
	srcDir := fmt.Sprintf("%s/%s", c.CmdFolder(), c.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)
	ctx := c.Cli.Context

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		if err = typcore.WriteAppMain(ctx, src, c.ProjectPackage+"/typical"); err != nil {
			return
		}
	}

	gobuild := buildkit.NewGoBuild(binary, src).
		WithStdout(b.stdout).
		WithStderr(b.stderr)

	if err = gobuild.Execute(ctx); err != nil {
		return
	}

	return []BuildDistribution{
		&GoBinary{
			binary: binary,
			stdout: b.stdout,
			stderr: b.stderr,
			stdin:  b.stdin,
		},
	}, nil
}
