package typgo

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

type (
	// PrepareContext context related with build-tool
	PrepareContext struct {
		Name       string
		Descriptor *Descriptor
		Stdout     io.Writer
	}
)

// NewPrepareContext return new instance of begin context
func NewPrepareContext(d *Descriptor, name string) *PrepareContext {
	return &PrepareContext{
		Name:       name,
		Descriptor: d,
		Stdout:     os.Stdout,
	}
}

// Info log text
func (c *PrepareContext) Info(text string) {
	c.printHeader()
	fmt.Fprintln(c.Stdout, text)
}

// Infof formatted text
func (c *PrepareContext) Infof(format string, a ...interface{}) {
	c.printHeader()
	fmt.Fprintf(c.Stdout, format, a...)
}

func (c *PrepareContext) printHeader() {
	if c.Descriptor != nil {
		color.New(ProjectNameColor).Fprint(c.Stdout, c.Descriptor.ProjectName)

	}
	if c.Name != "" {
		fmt.Fprint(c.Stdout, ":")
		color.New(PrepareColor).Fprint(c.Stdout, c.Name)
	}
	fmt.Fprint(c.Stdout, "> ")
}
