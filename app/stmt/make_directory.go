package stmt

import (
	"os"
)

type Mkdir struct {
	Path string
}

func (md Mkdir) Run() error {
	return os.MkdirAll(md.Path, 0700)
}
