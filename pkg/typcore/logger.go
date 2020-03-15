package typcore

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// Logger responsible to log any useful information
type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
}

// TypicalLogger is simple logger
type TypicalLogger struct {
	w io.Writer
}

// NewLogger return new instance of TypicalLogger
func NewLogger() *TypicalLogger {
	return &TypicalLogger{
		w: os.Stdout,
	}
}

// Info logs level message
func (s *TypicalLogger) Info(args ...interface{}) {
	s.signature()
	fmt.Fprintln(s.w, args...)
}

// Infof is same with Info with formatted
func (s *TypicalLogger) Infof(format string, args ...interface{}) {
	s.signature()
	fmt.Fprintf(s.w, format, args...)
	fmt.Fprintln(s.w)
}

func (s *TypicalLogger) signature() {
	fmt.Print("[")
	color.New(color.FgHiBlue).Print("TYPICAL")
	fmt.Print("] ")
}
