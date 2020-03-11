package exor

import (
	"context"
	"os"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// WriteTemplate responsible to write template to file
type WriteTemplate struct {
	target     string
	template   string
	data       interface{}
	permission os.FileMode
}

// NewWriteTemplate return new instance of runner
func NewWriteTemplate(target, template string, data interface{}) *WriteTemplate {
	return &WriteTemplate{
		target:     target,
		template:   template,
		data:       data,
		permission: 0777,
	}
}

// WithPermission return WriteTemplateRunner with new permission
func (w *WriteTemplate) WithPermission(permission os.FileMode) *WriteTemplate {
	w.permission = permission
	return w
}

// Execute write template
func (w *WriteTemplate) Execute(ctx context.Context) (err error) {
	var (
		f    *os.File
		tmpl *template.Template
	)

	log.Infof("Write File: %s", w.target)
	if f, err = os.OpenFile(w.target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, w.permission); err != nil {
		return
	}
	if tmpl, err = template.New("").Parse(w.template); err != nil {
		return
	}
	return tmpl.Execute(f, w.data)
}
