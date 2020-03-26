package typlog

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// TypicalLogger is simple logger
type TypicalLogger struct {
	io.Writer
}

// New instance of TypicalLogger
func New() *TypicalLogger {
	return &TypicalLogger{
		Writer: os.Stdout,
	}
}

// WithWriter return TypicalLogger with new writer
func (s *TypicalLogger) WithWriter(w io.Writer) *TypicalLogger {
	s.Writer = w
	return s
}

// Info level message
func (s *TypicalLogger) Info(args ...interface{}) {
	s.infoSign()
	s.print(args...)
}

// Infof is same with Info but with format
func (s *TypicalLogger) Infof(format string, args ...interface{}) {
	s.infoSign()
	s.printf(format, args...)
}

// Warn level log message
func (s *TypicalLogger) Warn(args ...interface{}) {
	s.warnSign()
	s.print(args...)
}

// Warnf is same with warn but with format
func (s *TypicalLogger) Warnf(format string, args ...interface{}) {
	s.warnSign()
	s.printf(format, args...)
}

func (s *TypicalLogger) print(args ...interface{}) {
	fmt.Fprintln(s, args...)
}

func (s *TypicalLogger) printf(format string, args ...interface{}) {
	fmt.Fprintf(s, format, args...)
	fmt.Fprintln(s)
}

func (s *TypicalLogger) infoSign() {
	s.typicalSign()
	fmt.Fprint(s, "[")
	color.New(color.FgCyan).Fprint(s, "INFO")
	fmt.Fprint(s, "] ")
}

func (s *TypicalLogger) warnSign() {
	s.typicalSign()
	fmt.Fprint(s, "[")
	color.New(color.FgYellow).Fprint(s, "WARN")
	fmt.Fprint(s, "] ")
}

func (s TypicalLogger) typicalSign() {
	fmt.Fprint(s, "[")
	color.New(color.FgHiBlue).Fprint(s, "TYPICAL")
	fmt.Fprint(s, "]")
}
