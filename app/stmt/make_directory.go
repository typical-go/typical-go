package stmt

import (
	"os"
)

type MakeDirectory struct {
	Path string
}

func (md MakeDirectory) Run() error {
	return os.MkdirAll(md.Path, os.ModePerm)
}
