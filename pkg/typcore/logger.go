package typcore

import "fmt"

// Logger responsible to log any useful information
type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
}

// SimpleLogger is simple logger
type SimpleLogger struct{}

// Info logs level message
func (s *SimpleLogger) Info(args ...interface{}) {
	fmt.Print("[TYPICAL] ")
	fmt.Println(args...)
}

// Infof is same with Info with formatted
func (s *SimpleLogger) Infof(format string, args ...interface{}) {
	fmt.Print("[TYPICAL] ")
	fmt.Printf(format, args...)
	fmt.Println()
}
