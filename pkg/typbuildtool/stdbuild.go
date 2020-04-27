package typbuildtool

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

var (
	_ Cleaner  = (*StdBuild)(nil)
	_ Tester   = (*StdBuild)(nil)
	_ Releaser = (*StdBuild)(nil)
	_ Runner   = (*StdBuild)(nil)
)

// StdBuild is standard build module for go project
type StdBuild struct {
	stdout         io.Writer
	stderr         io.Writer
	stdin          io.Reader
	coverProfile   string
	releaseTargets []ReleaseTarget // TODO: move to BuildTool
	releaseFolder  string

	testTimeout time.Duration
}

// StandardBuild return new instance of Module
func StandardBuild() *StdBuild {
	return &StdBuild{
		stdout:        os.Stdout,
		stderr:        os.Stderr,
		stdin:         os.Stdin,
		coverProfile:  "cover.out",
		releaseFolder: "release",
		testTimeout:   20 * time.Second,
	}
}

// Validate the releaser
func (b *StdBuild) Validate() (err error) {
	for _, target := range b.releaseTargets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}

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

// Clean build result
func (b *StdBuild) Clean(c *CliContext) (err error) {
	c.Info("Standard-Build: Clean the project")
	c.Infof("Remove All in '%s'", c.BuildTool.binFolder)
	if err := os.RemoveAll(c.BuildTool.binFolder); err != nil {
		c.Warn(err.Error())
	}
	return
}

// Release this project
func (b *StdBuild) Release(c *ReleaseContext) (files []string, err error) {
	for _, target := range b.releaseTargets {
		c.Infof("Build release for %s", target)
		var binary string
		if binary, err = b.releaseBuild(c, target); err != nil {
			err = fmt.Errorf("Failed build release: %w", err)
			return
		}
		files = append(files, binary)
	}

	return
}

func (b *StdBuild) releaseBuild(c *ReleaseContext, target ReleaseTarget) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{c.Core.Name, c.Tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(c.Context,
		"go", "build",
		"-o", fmt.Sprintf("%s/%s", b.releaseFolder, binary),
		"-ldflags", "-w -s",
		fmt.Sprintf("./%s/%s", c.BuildTool.cmdFolder, c.Core.Name),
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = b.stdout
	cmd.Stderr = b.stderr

	err = cmd.Run()
	return
}
