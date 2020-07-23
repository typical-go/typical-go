package typtmpl

import (
	"io"
	"os"
	"text/template"
)

type (
	// Template responsible to write
	Template interface {
		Execute(io.Writer) error
	}
)

// Execute template
func Execute(w io.Writer, tmpl Template) error {
	return tmpl.Execute(w)
}

// ExecuteToFile to execute template to file
func ExecuteToFile(target string, tmpl Template) error {
	f, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return Execute(f, tmpl)
}

// Parse template
func Parse(name, text string, data interface{}, w io.Writer) (err error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return
	}
	return tmpl.Execute(w, data)
}
