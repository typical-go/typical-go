package common

import (
	"os"
	"text/template"
)

// WriteTemplate to write template to file
func WriteTemplate(filename, text string, data interface{}) (err error) {
	var (
		f    *os.File
		tmpl *template.Template
	)
	if f, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		return
	}
	if tmpl, err = template.New("").Parse(text); err != nil {
		return
	}
	return tmpl.Execute(f, data)
}
