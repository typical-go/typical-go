package runnerkit

import (
	"context"
	"os"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// WriteTemplate runner
func WriteTemplate(target, template string, data interface{}, permission os.FileMode) Runner {
	return &writeTemplate{
		target:     target,
		template:   template,
		data:       data,
		permission: permission,
	}
}

// WriteTemplate to write template to file
type writeTemplate struct {
	target     string
	template   string
	data       interface{}
	permission os.FileMode
}

func (w *writeTemplate) Run(ctx context.Context) (err error) {
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
