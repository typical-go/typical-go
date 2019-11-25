package common

import (
	"io/ioutil"
	"os"
)

// WriteString to write string to file
type WriteString struct {
	Target     string
	Content    string
	Permission os.FileMode
}

// Run to write file
func (w WriteString) Run() (err error) {
	return ioutil.WriteFile(w.Target, []byte(w.Content), w.Permission)
}
