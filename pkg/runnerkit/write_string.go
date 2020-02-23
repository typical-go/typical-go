package runnerkit

import (
	"context"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// WriteString return new instance of WriteString
func WriteString(target, content string, permission os.FileMode) Runner {
	return &writeString{
		target:     target,
		content:    content,
		permission: permission,
	}
}

type writeString struct {
	target     string
	content    string
	permission os.FileMode
}

func (w writeString) Run(ctx context.Context) (err error) {
	log.Infof("Write File: %s", w.target)
	return ioutil.WriteFile(w.target, []byte(w.content), w.permission)
}
