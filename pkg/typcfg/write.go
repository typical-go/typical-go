package typcfg

import (
	"fmt"
	"os"
)

// Write configuration to file
func Write(dest string, c Configurer) (err error) {
	var (
		fields []*Field
		f      *os.File
		m      map[string]string
	)

	for _, cfg := range c.Configurations() {
		for _, field := range cfg.Fields() {
			fields = append(fields, field)
		}
	}

	if f, err = os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return
	}
	defer f.Close()

	m = Read(f)

	stat, _ := f.Stat()
	if stat.Size() > 0 {
		fmt.Fprintln(f)
	}

	for _, field := range fields {
		if _, ok := m[field.Name]; !ok {
			fmt.Fprintf(f, "%s=%v\n", field.Name, field.GetValue())
		}
	}

	return

}
