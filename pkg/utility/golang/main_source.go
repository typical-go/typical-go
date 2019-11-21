package golang

import "io"

// MainSource is main source code
type MainSource struct {
	*Source
	MainFunc *Function
}

// NewMainSource return new instance of MainSource
func NewMainSource() MainSource {
	return MainSource{
		Source:   NewSource("main"),
		MainFunc: NewFunction("main"),
	}
}

// Write to
func (s MainSource) Write(w io.Writer) (err error) {
	if err = s.Source.Write(w); err != nil {
		return
	}
	return s.MainFunc.Write(w)
}
