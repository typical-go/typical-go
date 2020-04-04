package typcfg

import (
	"fmt"
	"io"
	"os"
)

// Write configuration to file
func Write(m *ConfigManager, dest string) (err error) {
	var f *os.File
	if f, err = os.Create(dest); err != nil {
		return
	}
	defer f.Close()

	return write(m, f)
}

func write(m *ConfigManager, w io.Writer) (err error) {
	for _, cfg := range m.Configurations() {
		for _, field := range RetrieveFields(cfg) {
			var v interface{}
			if field.IsZero {
				v = field.Default
			} else {
				v = field.Value
			}
			if _, err = fmt.Fprintf(w, "%s=%v\n", field.Name, v); err != nil {
				return
			}
		}
	}
	return
}
