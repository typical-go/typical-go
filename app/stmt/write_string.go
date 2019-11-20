package stmt

import "io/ioutil"

// WriteString to write string to file
type WriteString struct {
	Target  string
	Content string
}

// Run to write file
func (w WriteString) Run() (err error) {
	return ioutil.WriteFile(w.Target, []byte(w.Content), 0644)
}
