package typbuildtool

import (
	"fmt"
	"io"
	"os"

	"github.com/typical-go/typical-go/pkg/runnerkit"
	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typbuildtool/internal/tmpl"
)

// StdBuilder is standard builder
type StdBuilder struct {
	stdout       io.Writer
	stderr       io.Writer
	preExecutors []runnerkit.Runner
}

// NewBuilder return new instance of standard builder
func NewBuilder() *StdBuilder {
	return &StdBuilder{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

// WithStdout return StdBuilder with new stdout
func (b *StdBuilder) WithStdout(stdout io.Writer) *StdBuilder {
	b.stdout = stdout
	return b
}

// WithStderr return StdBuilder with new stderr
func (b *StdBuilder) WithStderr(stderr io.Writer) *StdBuilder {
	b.stderr = stderr
	return b
}

// Before build execution
func (b *StdBuilder) Before(executor ...runnerkit.Runner) *StdBuilder {
	b.preExecutors = executor
	return b
}

// Build the project
func (b *StdBuilder) Build(c *BuildContext) (dist BuildDistribution, err error) {
	binary := fmt.Sprintf("%s/%s", c.BinFolder, c.Name)
	srcDir := fmt.Sprintf("%s/%s", c.CmdFolder, c.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)
	ctx := c.Cli.Context

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		data := &tmpl.AppMainData{
			TypicalPackage: typcore.TypicalPackage(c.ProjectPackage),
		}
		if err = runnerkit.NewWriteTemplate(src, tmpl.AppMain, data).Run(ctx); err != nil {
			return
		}
	}

	for _, executor := range b.preExecutors {
		if err = executor.Run(ctx); err != nil {
			return
		}
	}

	cmd := buildkit.NewGoBuild(binary, src).Command(ctx)
	cmd.Stdout = b.stdout
	cmd.Stderr = b.stderr

	if err = cmd.Run(); err != nil {
		return
	}

	return NewBuildDistribution(binary), nil
}
