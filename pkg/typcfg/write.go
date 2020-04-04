package typcfg

import (
	"fmt"
	"io"
	"os"
)

// Write configuration to file
func Write(dest string, m Configurer) (err error) {
	var (
		fields []*Field
		f      *os.File
	)

	for _, cfg := range m.Configure() {
		for _, field := range Fields(cfg) {
			fields = append(fields, field)
		}
	}

	if _, err = os.Stat(dest); os.IsNotExist(err) {
		if f, err = os.Create(dest); err != nil {
			return
		}
		defer f.Close()
		return write(f, fields)
	}

	return

}

func write(w io.Writer, fields []*Field) (err error) {
	for _, field := range fields {
		var v interface{}
		if field.IsZero {
			v = field.Default
		} else {
			v = field.Value
		}
		fmt.Fprintf(w, "%s=%v\n", field.Name, v)
	}

	return
}
