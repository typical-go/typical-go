package typtmpl

import (
	"io"
	"os"
)

// Template responsible to write
type Template interface {
	Execute(io.Writer) error
}

// WriteFile to write file
func WriteFile(filename string, perm os.FileMode, tmpl Template) (err error) {
	var f *os.File
	if f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm); err != nil {
		return
	}
	defer f.Close()

	return tmpl.Execute(f)
}
