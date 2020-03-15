package exor

import (
	"context"
	"io/ioutil"
	"os"
)

// WriteString responsible to write string
type WriteString struct {
	target     string
	content    string
	permission os.FileMode
}

// NewWriteString return new instance of WriteString
func NewWriteString(target, content string) *WriteString {
	return &WriteString{
		target:     target,
		content:    content,
		permission: 0777,
	}
}

// WithPermission return new instance of WriteString
func (w *WriteString) WithPermission(permission os.FileMode) *WriteString {
	w.permission = permission
	return w
}

// Execute write string
func (w *WriteString) Execute(ctx context.Context) (err error) {
	return ioutil.WriteFile(w.target, []byte(w.content), w.permission)
}
