package typtmpl

import (
	"io"
)

var _ Template = (*GitIgnore)(nil)

const gitignore = `/bin
/release
/.typical-tmp
/vendor 
*.envrc
*.env
*.test
*.out
`

// GitIgnore writer
type GitIgnore struct{}

// Execute GitIgnore
func (*GitIgnore) Execute(w io.Writer) (err error) {
	_, err = w.Write([]byte(gitignore))
	return
}
