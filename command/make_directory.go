package command

import (
	"os"
)

type makeDirectory struct {
	path string
}

// MakeDirectory execute `mkdir` in linux bash
func MakeDirectory(path string) Command {
	return &makeDirectory{path: path}
}

func (md makeDirectory) Run() error {
	return os.MkdirAll(md.path, os.ModePerm)
}
