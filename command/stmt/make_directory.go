package stmt

import (
	"os"
)

// MakeDirectory execute `mkdir` in linux bash
func MakeDirectory(path string) Statement {
	return &makeDirectory{path: path}
}

type makeDirectory struct {
	path string
}

func (md makeDirectory) Run() error {
	return os.MkdirAll(md.path, os.ModePerm)
}
