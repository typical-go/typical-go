package typlog

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var (

	// DefaultName is default name in log signature
	DefaultName string = "TYPICAL"

	// DefaultColor is default name in log signature
	DefaultColor color.Attribute = color.FgHiBlue
)

// Logger is simple logger
type Logger struct {
	Name  string
	Color color.Attribute
	Out   io.Writer
}

// Info level message
func (s *Logger) Info(args ...interface{}) {
	s.infoSign()
	s.print(args...)
}

// Infof is same with Info but with format
func (s *Logger) Infof(format string, args ...interface{}) {
	s.infoSign()
	s.printf(format, args...)
}

// Warn level log message
func (s *Logger) Warn(args ...interface{}) {
	s.warnSign()
	s.print(args...)
}

// Warnf is same with warn but with format
func (s *Logger) Warnf(format string, args ...interface{}) {
	s.warnSign()
	s.printf(format, args...)
}

func (s *Logger) print(args ...interface{}) {
	fmt.Fprintln(s, args...)
}

func (s *Logger) printf(format string, args ...interface{}) {
	fmt.Fprintf(s, format, args...)
	fmt.Fprintln(s)
}

func (s *Logger) infoSign() {
	s.typicalSign()
	fmt.Fprint(s, "[")
	color.New(color.FgCyan).Fprint(s, "INFO")
	fmt.Fprint(s, "] ")
}

func (s *Logger) warnSign() {
	s.typicalSign()
	fmt.Fprint(s, "[")
	color.New(color.FgYellow).Fprint(s, "WARN")
	fmt.Fprint(s, "] ")
}

func (s *Logger) typicalSign() {
	if s.Name == "" {
		s.Name = DefaultName
	}

	if s.Color == 0 {
		s.Color = DefaultColor
	}

	fmt.Fprint(s, "[")
	color.New(s.Color).Fprint(s, s.Name)
	fmt.Fprint(s, "]")
}

func (s Logger) Write(p []byte) (n int, err error) {
	if s.Out == nil {
		s.Out = os.Stdout
	}
	return s.Out.Write(p)
}
