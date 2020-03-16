package typbuildtool

import (
	"fmt"
	"io"
	"os"

	"github.com/typical-go/typical-go/pkg/exor"

	"github.com/typical-go/typical-go/pkg/typbuildtool/internal/tmpl"
)

// StdBuilder is standard builder
type StdBuilder struct {
	stdout       io.Writer
	stderr       io.Writer
	preExecutors []exor.Executor
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
func (b *StdBuilder) Before(executor ...exor.Executor) *StdBuilder {
	b.preExecutors = executor
	return b
}

// Build the project
func (b *StdBuilder) Build(c *Context) (dist BuildDistribution, err error) {
	binary := fmt.Sprintf("%s/%s", c.BinFolder(), c.Name)
	srcDir := fmt.Sprintf("%s/%s", c.CmdFolder(), c.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)
	ctx := c.Cli.Context

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		data := &tmpl.AppMainData{
			TypicalPackage: c.ProjectPackage + "/typical",
		}
		if err = exor.NewWriteTemplate(src, tmpl.AppMain, data).Execute(ctx); err != nil {
			return
		}
	}

	if err = exor.Execute(ctx, b.preExecutors...); err != nil {
		return
	}

	gobuild := exor.NewGoBuild(binary, src).
		WithStdout(b.stdout).
		WithStderr(b.stderr)

	if err = gobuild.Execute(ctx); err != nil {
		return
	}

	return NewBuildDistribution(binary), nil
}

// Clean build result
func (b *StdBuilder) Clean(c *Context) (err error) {
	c.Infof("Remove All in '%s'", c.BinFolder())
	if err := os.RemoveAll(c.BinFolder()); err != nil {
		c.Error(err.Error())
	}
	return
}
