package typgo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

type (
	// BuildToolContext context related with build-tool
	BuildToolContext struct {
		Name       string
		Descriptor *Descriptor
		Stdout     io.Writer
	}
)

// NewBuildToolContext return new instance of begin context
func NewBuildToolContext(d *Descriptor, name string) *BuildToolContext {
	return &BuildToolContext{
		Name:       name,
		Descriptor: d,
		Stdout:     os.Stdout,
	}
}

// DummyBuildToolContext return dummy BuildToolContext
func DummyBuildToolContext() (*BuildToolContext, *strings.Builder) {
	var out strings.Builder
	return &BuildToolContext{
		Name: "dummy",
		Descriptor: &Descriptor{
			ProjectName:    "some-project",
			ProjectVersion: "0.0.1",
		},
		Stdout: &out,
	}, &out
}

// Info log text
func (c *BuildToolContext) Info(text string) {
	c.printHeader()
	fmt.Fprintln(c.Stdout, text)
}

// Infof formatted text
func (c *BuildToolContext) Infof(format string, a ...interface{}) {
	c.printHeader()
	fmt.Fprintf(c.Stdout, format, a...)
}

func (c *BuildToolContext) printHeader() {
	color.New(ProjectNameColor).Fprint(c.Stdout, c.Descriptor.ProjectName)
	fmt.Fprint(c.Stdout, ":")
	color.New(PrepareColor).Fprint(c.Stdout, c.Name)
	fmt.Fprint(c.Stdout, "> ")
}
