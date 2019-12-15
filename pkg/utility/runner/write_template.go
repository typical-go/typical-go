package runner

import (
	"os"
	"text/template"
)

// WriteTemplate to write template to file
type WriteTemplate struct {
	Target     string
	Template   string
	Data       interface{}
	Permission os.FileMode
}

// Run to write file
func (w WriteTemplate) Run() (err error) {
	var f *os.File
	var tmpl *template.Template
	if w.Permission == 0 {
		w.Permission = 0666
	}
	if f, err = os.OpenFile(w.Target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, w.Permission); err != nil {
		return
	}
	if tmpl, err = template.New("").Parse(w.Template); err != nil {
		return
	}
	return tmpl.Execute(f, w.Data)
}
