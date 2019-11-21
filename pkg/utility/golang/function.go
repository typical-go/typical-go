package golang

import (
	"fmt"
	"io"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

// Function definition
type Function struct {
	coll.Strings
	Name    string
	Params  []coll.KeyString
	Returns []coll.KeyString
}

// Return statement
func (f *Function) Return(s ...string) {
	f.Append("return " + strings.Join(s, ", "))
}

func (f *Function) Write(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("func %s", f.Name)))
	w.Write([]byte("("))
	for i, param := range f.Params {
		if i > 0 {
			w.Write([]byte(", "))
		}
		w.Write([]byte(param.SimpleFormat(" ")))
	}
	w.Write([]byte(") "))
	if len(f.Returns) > 0 {
		w.Write([]byte("("))
		for i, ret := range f.Returns {
			if i > 0 {
				w.Write([]byte(", "))
			}
			w.Write([]byte(ret.SimpleFormat(" ")))
		}
		w.Write([]byte(") "))
	}
	w.Write([]byte("{\n"))
	for _, s := range f.Strings {
		w.Write([]byte(s))
		w.Write([]byte("\n"))
	}
	w.Write([]byte("}\n"))
}
