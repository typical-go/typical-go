package golang

import (
	"fmt"
	"io"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
)

// Struct is plain old go object for struct
type Struct struct {
	Name        string
	Fields      common.StringDictionary
	Description string
}

// AddField to add field to struct
func (s *Struct) AddField(name, typ string) {
	s.Fields.Add(name, typ)
}

func (s Struct) Write(w io.Writer) (err error) {
	fmt.Fprintf(w, "// %s %s\n", s.Name, s.Description)
	fmt.Fprintf(w, "type %s struct{\n", s.Name)
	for _, kv := range s.Fields {
		fmt.Fprintf(w, "%s %s\n", kv.Key, kv.Value)
	}
	fmt.Fprintln(w, "}")
	return
}

func (s Struct) String() string {
	var builder strings.Builder
	s.Write(&builder)
	return builder.String()
}
