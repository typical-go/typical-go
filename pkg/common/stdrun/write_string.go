package stdrun

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// WriteString to write string to file
type WriteString struct {
	target     string
	content    string
	permission os.FileMode
}

// NewWriteString return new instance of WriteString
func NewWriteString(target, content string) *WriteString {
	return &WriteString{
		target:  target,
		content: content,
	}
}

// WithPermission to set permission and return WriteString
func (w *WriteString) WithPermission(permission os.FileMode) *WriteString {
	w.permission = permission
	return w
}

// Run to write file
func (w WriteString) Run() (err error) {
	log.Infof("Write File: %s", w.target)
	if w.permission == 0 {
		w.permission = 0666
	}
	return ioutil.WriteFile(w.target, []byte(w.content), w.permission)
}
