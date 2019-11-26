package runner

import (
	"os"
)

// Mkdir to make directory
type Mkdir struct {
	Path string
}

// Run to making the directory
func (md Mkdir) Run() error {
	return os.MkdirAll(md.Path, 0700)
}
