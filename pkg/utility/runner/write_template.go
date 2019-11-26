package runner

import (
	"html/template"
	"os"
)

// WriteTemplate to write template to file
type WriteTemplate struct {
	Target   string
	Template string
	Data     interface{}
}

// Run to write file
func (w WriteTemplate) Run() (err error) {
	var f *os.File
	var tmpl *template.Template
	if f, err = os.Create(w.Target); err != nil {
		return
	}
	if tmpl, err = template.New("").Parse(w.Template); err != nil {
		return
	}
	return tmpl.Execute(f, w.Data)
}
