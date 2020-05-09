package typbuild

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
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
	binary := fmt.Sprintf("%s/%s", typvar.BinFolder, c.Core.Name)
	srcDir := fmt.Sprintf("%s/%s", typvar.CmdFolder, c.Core.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		appMain := &typtmpl.AppMain{
			DescPkg: typvar.ProjectPkg + "/typical",
		}

		if err = typtmpl.WriteFile(src, 0777, appMain); err != nil {
			return fmt.Errorf("%s: %w", srcDir, err)
		}
	}

	gobuild := buildkit.NewGoBuild(binary, src).
		WithStdout(b.stdout).
		WithStderr(b.stderr).
		Command()

	gobuild.Print(os.Stdout)

	ctx := c.Cli.Context
	if err = gobuild.Run(ctx); err != nil {
		return fmt.Errorf("GoBuild: %w", err)
	}

	binExec := &buildkit.Command{
		Name:   binary,
		Args:   c.Cli.Args().Slice(),
		Stdout: b.stdout,
		Stderr: b.stderr,
	}

	binExec.Print(os.Stdout)
	fmt.Printf("\n\n")

	if err = binExec.Run(ctx); err != nil {
		return fmt.Errorf("%s: %w", binary, err)
	}

	return
}

// Test the project
func (b *StdBuild) Test(c *CliContext) (err error) {
	var (
		targets []string
	)

	for _, layout := range c.Core.Layouts {
		targets = append(targets, fmt.Sprintf("./%s/...", layout))
	}

	if len(targets) < 1 {
		c.Info("Nothing to test")
		return
	}

	gotest := buildkit.GoTest{
		Targets:      targets,
		Timeout:      b.testTimeout,
		CoverProfile: b.coverProfile,
		Race:         true,
	}

	cmd := gotest.Command()
	cmd.Stdout = b.stdout
	cmd.Stderr = b.stderr

	cmd.Print(os.Stdout)
	fmt.Println()

	ctx := c.Cli.Context

	return cmd.Run(ctx)
}

// Clean build result
func (b *StdBuild) Clean(c *CliContext) (err error) {
	c.Infof("Remove All in '%s'", typvar.BinFolder)
	if err := os.RemoveAll(typvar.BinFolder); err != nil {
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

	gobuild := buildkit.Command{
		Name: "go",
		Args: []string{
			"build",
			"-o", fmt.Sprintf("%s/%s", b.releaseFolder, binary),
			"-ldflags", "-w -s",
			fmt.Sprintf("./%s/%s", typvar.CmdFolder, c.Core.Name),
		},
		Stdout: b.stdout,
		Stderr: b.stderr,
		Env:    append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch),
	}

	ctx := c.Cli.Context
	err = gobuild.Run(ctx)
	return
}
