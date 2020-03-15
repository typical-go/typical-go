package typcore

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Logger responsible to log any useful information
type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
}

// TypicalLogger is simple logger
type TypicalLogger struct {
	io.Writer
}

// NewLogger return new instance of TypicalLogger
func NewLogger() *TypicalLogger {
	return &TypicalLogger{
		Writer: os.Stdout,
	}
}

// Info logs level message
func (s *TypicalLogger) Info(args ...interface{}) {
	s.signature("info")
	fmt.Fprintln(s, args...)
}

// Infof is same with Info with formatted
func (s *TypicalLogger) Infof(format string, args ...interface{}) {
	s.signature("info")
	fmt.Fprintf(s, format, args...)
	fmt.Fprintln(s)
}

func (s *TypicalLogger) Error(args ...interface{}) {
	s.signature("error")
	fmt.Fprintln(s, args...)
}

func (s *TypicalLogger) signature(level string) {
	fmt.Fprint(s, "[")
	color.New(color.FgHiBlue).Fprint(s, "TYPICAL")
	fmt.Fprint(s, "]")

	if level != "" {
		fmt.Fprintf(s, "[%s]", strings.ToUpper(level))
	}

	fmt.Fprint(s, " ")
}
