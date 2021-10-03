package typgen

import (
	"strings"
)

type (
	SourceCoder interface {
		SourceCode() string
	}
	SourceCoders []SourceCoder
	Comment      string
)

//
// SourceCoders
//

var _ SourceCoder = (SourceCoders)(nil)

func (s SourceCoders) SourceCode() string {
	var b strings.Builder
	for _, src := range s {
		b.WriteString(src.SourceCode())
		b.WriteString("\n")
	}

	return b.String()
}

//
// Comment
//

func (c Comment) SourceCode() string {
	return "// " + string(c)
}
