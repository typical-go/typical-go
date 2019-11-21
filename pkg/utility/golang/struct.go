package golang

import (
	"fmt"
	"io"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

// Struct is plain old go object for struct
type Struct struct {
	Name        string
	Fields      coll.KeyStrings
	Description string
}

// AddField to add field to struct
func (s *Struct) AddField(name, typ string) {
	s.Fields.Append(coll.KeyString{Key: name, String: typ})
}

func (s Struct) Write(w io.Writer) (err error) {
	fmt.Fprintf(w, "// %s %s\n", s.Name, s.Description)
	fmt.Fprintf(w, "type %s struct{\n", s.Name)
	for _, field := range s.Fields {
		fmt.Fprintln(w, field.SimpleFormat(" "))
	}
	fmt.Fprintln(w, "}")
	return
}

func (s Struct) String() string {
	var builder strings.Builder
	s.Write(&builder)
	return builder.String()
}
