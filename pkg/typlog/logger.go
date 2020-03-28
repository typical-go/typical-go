package typlog

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var (
	// Output for logger
	output io.Writer = os.Stdout

	// DefaultName is default name in log signature
	DefaultName string = "TYPICAL"

	// DefaultColor is default name in log signature
	DefaultColor color.Attribute = color.FgHiBlue
)

// Logger is simple logger
type Logger struct {
	name  string
	color color.Attribute
}

// SetOutput to set logger output
func SetOutput(w io.Writer) func() {
	output = w
	return func() {
		output = os.Stdout
	}
}

// SetLogSignature to change signature. It return method to reset the signature.
func (s *Logger) SetLogSignature(name string, color color.Attribute) func() {
	s.name = name
	s.color = color
	return s.ResetLogSignature
}

// ResetLogSignature to reset signature
func (s *Logger) ResetLogSignature() {
	s.name = DefaultName
	s.color = DefaultColor
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
	if s.name == "" {
		s.ResetLogSignature()
	}

	fmt.Fprint(s, "[")
	color.New(s.color).Fprint(s, s.name)
	fmt.Fprint(s, "]")
}

func (s Logger) Write(p []byte) (n int, err error) {
	return output.Write(p)
}
