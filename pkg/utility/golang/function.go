package golang

import (
	"fmt"
	"io"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
)

// Function definition
type Function struct {
	common.Strings
	Name    string
	Params  []common.KeyString
	Returns []common.KeyString
}

// NewFunction return new instance
func NewFunction(name string) *Function {
	return &Function{Name: name}
}

// Return statement
func (f *Function) Return(s ...string) {
	f.Append("return " + strings.Join(s, ", "))
}

func (f *Function) Write(w io.Writer) (err error) {
	fmt.Fprintf(w, "func %s", f.Name)
	fmt.Fprint(w, "(")
	for i, param := range f.Params {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprint(w, param.SimpleFormat(" "))
	}
	fmt.Fprint(w, ") ")
	if len(f.Returns) > 0 {
		fmt.Fprint(w, "(")
		for i, ret := range f.Returns {
			if i > 0 {
				fmt.Fprint(w, ", ")
			}
			fmt.Fprint(w, ret.SimpleFormat(" "))
		}
		fmt.Fprint(w, ") ")
	}
	fmt.Fprintln(w, "{")
	for _, s := range f.Strings {
		fmt.Fprint(w, s)
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w, "}")
	return
}
