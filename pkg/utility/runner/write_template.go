package runner

import (
	"os"
	"text/template"
)

// WriteTemplate to write template to file
type WriteTemplate struct {
	target     string
	template   string
	data       interface{}
	permission os.FileMode
}

// NewWriteTemplate return new instance of WriteTemplate
func NewWriteTemplate(target, template string, data interface{}) *WriteTemplate {
	return &WriteTemplate{
		target:   target,
		template: template,
		data:     data,
	}
}

// WithPermission to set permission and return WritePermission
func (w *WriteTemplate) WithPermission(permission os.FileMode) *WriteTemplate {
	w.permission = permission
	return w
}

// Run to write file
func (w *WriteTemplate) Run() (err error) {
	var f *os.File
	var tmpl *template.Template
	if w.permission == 0 {
		w.permission = 0666
	}
	if f, err = os.OpenFile(w.target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, w.permission); err != nil {
		return
	}
	if tmpl, err = template.New("").Parse(w.template); err != nil {
		return
	}
	return tmpl.Execute(f, w.data)
}
