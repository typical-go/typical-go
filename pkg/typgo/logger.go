package typgo

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

type (
	// Logger logger
	Logger struct {
		Headers []LogHeader
		Stdout  io.Writer
	}
	LogHeader struct {
		Text  string
		Color color.Attribute
	}
)

func LogHeaders(taskNames ...string) []LogHeader {
	var headers []LogHeader
	for _, taskName := range taskNames {
		headers = append(headers, LogHeader{
			Text:  taskName,
			Color: ColorSet.Task,
		})
	}
	return headers
}

// Warn log text
func (c Logger) Warn(a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	color.New(ColorSet.Warn).Fprintln(c.Stdout, a...)
}

// Warnf formatted text
func (c Logger) Warnf(format string, a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	color.New(ColorSet.Warn).Fprintf(c.Stdout, format, a...)
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
	color.New(ColorSet.Bash).Fprintln(c.Stdout, bash)
}

func (c Logger) printHeader() {
	for i, header := range c.Headers {
		if i > 0 {
			fmt.Fprint(c.Stdout, ":")
		}
		color.New(header.Color).Fprint(c.Stdout, header.Text)
	}
	fmt.Fprint(c.Stdout, "> ")
}
