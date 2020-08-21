package tmplkit

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

// Write template
func Write(w io.Writer, tmplText string, tmplData interface{}) error {
	name := fmt.Sprintf("%T", tmplData)
	tmpl, err := template.New(name).Parse(tmplText)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, tmplData)
}

// WriteFile write template to file
func WriteFile(target, tmplText string, tmplData interface{}) error {
	f, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return Write(f, tmplText, tmplData)
}
