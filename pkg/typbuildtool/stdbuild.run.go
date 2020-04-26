package typbuildtool

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typfactory"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

// Run the project locally
func (b *StdBuild) Run(c *CliContext) (err error) {
	c.Info("Standard-Build: Build the project")
	binary := fmt.Sprintf("%s/%s", c.BuildTool.binFolder, c.Core.Name)
	srcDir := fmt.Sprintf("%s/%s", c.BuildTool.cmdFolder, c.Core.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		appMain := &typfactory.AppMain{
			DescPkg: c.Core.ProjectPkg + "/typical",
		}

		if err = typfactory.WriteFile(src, 0777, appMain); err != nil {
			return fmt.Errorf("%s: %w", srcDir, err)
		}
	}

	gobuild := buildkit.NewGoBuild(binary, src).
		WithStdout(b.stdout).
		WithStderr(b.stderr)

	if err = gobuild.Execute(c.Context); err != nil {
		return fmt.Errorf("GoBuild: %w", err)
	}

	cmd := exec.CommandContext(c.Context, binary, c.Args().Slice()...)
	cmd.Stdout = b.stdout
	cmd.Stderr = b.stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", binary, err)
	}

	return
}
