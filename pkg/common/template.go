package common

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

// ExecuteTmpl parse and execute template
func ExecuteTmpl(w io.Writer, text string, data interface{}) error {
	name := fmt.Sprintf("%T", data)
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// ExecuteTmplToFile parse and execute template to file
func ExecuteTmplToFile(target, text string, data interface{}) error {
	f, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return ExecuteTmpl(f, text, data)
}
