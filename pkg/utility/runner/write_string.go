package runner

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
	if w.Permission == 0 {
		w.Permission = 0666
	}
	return ioutil.WriteFile(w.Target, []byte(w.Content), w.Permission)
}
