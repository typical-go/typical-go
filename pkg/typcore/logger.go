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
	Error(...interface{})
	Errorf(string, ...interface{})
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

// Info leveled message
func (s *TypicalLogger) Info(args ...interface{}) {
	s.infoSign()
	s.print(args...)
}

// Infof is same with Info with format
func (s *TypicalLogger) Infof(format string, args ...interface{}) {
	s.infoSign()
	s.printf(format, args...)
}

// Error leveled log message
func (s *TypicalLogger) Error(args ...interface{}) {
	s.errorSign()
	s.print(args...)
}

// Errorf is same with Info with format
func (s *TypicalLogger) Errorf(format string, args ...interface{}) {
	s.errorSign()
	fmt.Fprintf(s, format, args...)
	fmt.Fprintln(s)
}

func (s *TypicalLogger) print(args ...interface{}) {
	fmt.Fprintln(s, args...)
}

func (s *TypicalLogger) printf(format string, args ...interface{}) {
	fmt.Fprintf(s, format, args...)
	fmt.Println()
}

func (s *TypicalLogger) infoSign() {
	s.typicalSign()
	fmt.Fprint(s, "[")
	color.New(color.FgCyan).Fprint(s, "INFO")
	fmt.Fprint(s, "] ")
}

func (s *TypicalLogger) errorSign() {
	s.typicalSign()
	fmt.Fprint(s, "[")
	color.New(color.FgRed).Fprint(s, "ERRO")
	fmt.Fprint(s, "] ")
}

func (s TypicalLogger) typicalSign() {
	fmt.Fprint(s, "[")
	color.New(color.FgHiBlue).Fprint(s, "TYPICAL")
	fmt.Fprint(s, "]")
}
