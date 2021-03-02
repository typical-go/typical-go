package typgo

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

type (
	// Logger logger
	Logger struct {
		ProjectName string
		TaskNames   []string
		Stdout      io.Writer
	}
)

// Warn log text
func (c Logger) Warn(a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	color.New(color.FgYellow).Fprintln(c.Stdout, a...)
}

// Warnf formatted text
func (c Logger) Warnf(format string, a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	color.New(color.FgYellow).Fprintf(c.Stdout, format, a...)
}

// Info log text
func (c Logger) Info(a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	fmt.Fprintln(c.Stdout, a...)
}

// Infof formatted text
func (c Logger) Infof(format string, a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	fmt.Fprintf(c.Stdout, format, a...)
}

// Bash information
func (c Logger) Bash(bash *Bash) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	fmt.Fprintln(c.Stdout, bash)
}

func (c Logger) printHeader() {
	if c.ProjectName != "" {
		color.New(ProjectNameColor).Fprint(c.Stdout, c.ProjectName)
	}
	for _, name := range c.TaskNames {
		fmt.Fprint(c.Stdout, ":")
		color.New(TaskNameColor).Fprint(c.Stdout, name)
	}
	fmt.Fprint(c.Stdout, "> ")
}
