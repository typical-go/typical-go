package typgen

import (
	"strings"
)

type (
	Coder interface {
		Code() string
	}
	Coders  []Coder
	Comment string
)

//
// Coders
//

var _ Coder = (Coders)(nil)

func (s Coders) Code() string {
	var b strings.Builder
	for _, src := range s {
		b.WriteString(src.Code())
		b.WriteString("\n")
	}

	return b.String()
}

//
// Comment
//

func (c Comment) Code() string {
	return "// " + string(c)
}
