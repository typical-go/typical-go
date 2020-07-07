package typtmpl

import (
	"io"
	"text/template"
)

// Template responsible to write
type Template interface {
	Execute(io.Writer) error
}

// Execute template
func Execute(name, text string, data interface{}, w io.Writer) (err error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return
	}
	return tmpl.Execute(w, data)
}
