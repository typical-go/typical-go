package typfactory

import (
	"io"
)

const gitignore = `/bin
/release
/.typical-tmp
/vendor 
.envrc
.env
*.test
*.out
`

// GitIgnore writer
type GitIgnore struct{}

func (*GitIgnore) Write(w io.Writer) (err error) {
	_, err = w.Write([]byte(gitignore))
	return
}
