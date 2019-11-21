package golang

import (
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

func (s Struct) Write(w io.Writer) {
	writelnf(w, "// %s %s", s.Name, s.Description)
	writelnf(w, "type %s struct{", s.Name)
	for _, field := range s.Fields {
		writelnf(w, field.SimpleFormat(" "))
	}
	writeln(w, "}")
}

func (s Struct) String() string {
	var builder strings.Builder
	s.Write(&builder)
	return builder.String()
}
