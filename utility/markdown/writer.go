package markdown

import (
	"fmt"
	"strings"
)

type Writer struct {
	Builder strings.Builder
}

func (w Writer) Write(s string) (int, error) {
	return w.Builder.WriteString(s)
}

func (w Writer) Heading1(s string) {
	w.Write("# " + s)
}

func (w Writer) Heading2(s string) {
	w.Write("## " + s)
}

func (w Writer) Heading3(s string) {
	w.Write("### " + s)
}

func (w Writer) OrderedList(list []string) {
	for i, s := range list {
		w.Write(fmt.Sprintf("%d. %s", i, s))
	}
}

func (w Writer) Code(lang, code string) {
	w.Write("```` " + lang)
	w.Write(code)
	w.Write("```")
}

func (w Writer) BashCode(code string) {
	w.Code("bash", code)
}
