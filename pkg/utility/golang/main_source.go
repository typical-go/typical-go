package golang

import (
	"io"
)

// MainSource is main source code
type MainSource struct {
	Source
	MainFunc *Function
}

// NewMainSource return new instance of MainSource
func NewMainSource() *MainSource {
	return &MainSource{
		Source:   Source{Package: "main"},
		MainFunc: NewFunction("main"),
	}
}

// Write to apply the writer
func (s *MainSource) Write(w io.Writer) (err error) {
	if err = s.Source.Write(w); err != nil {
		return
	}
	if err = s.MainFunc.Write(w); err != nil {
		return
	}
	return
}

// Append codes
func (s *MainSource) Append(codes ...string) *MainSource {
	s.MainFunc.Append(codes...)
	return s
}
