package typtmpl

import (
	"io"
	"os"
	"text/template"
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

// Execute template
func Execute(name, text string, data interface{}, w io.Writer) (err error) {
	var tmpl *template.Template
	if tmpl, err = template.New(name).Parse(text); err != nil {
		return
	}
	return tmpl.Execute(w, data)
}
