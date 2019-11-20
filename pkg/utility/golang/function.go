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
	Params  []Param
	Returns []Param
}

// Param short from parameter
type Param struct {
	Name string
	Type string
}

// Return statement
func (f *Function) Return(s ...string) {
	f.Append("return " + strings.Join(s, ", "))
}

func (f *Function) Write(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("func %s", f.Name)))
	w.Write([]byte("("))
	if len(f.Params) > 0 {
		w.Write([]byte(paramsString(f.Params)))
	}
	w.Write([]byte(") "))
	if len(f.Returns) > 0 {
		w.Write([]byte("("))
		w.Write([]byte(paramsString(f.Returns)))
		w.Write([]byte(") "))
	}
	w.Write([]byte("{\n"))
	for _, s := range f.Strings {
		w.Write([]byte(s))
		w.Write([]byte("\n"))
	}
	w.Write([]byte("}\n"))
}

func paramsString(params []Param) string {
	var b strings.Builder
	for i, param := range params {
		if i > 0 {
			b.Write([]byte(", "))
		}
		b.Write([]byte(fmt.Sprintf("%s %s", param.Name, param.Type)))
	}
	return b.String()
}
