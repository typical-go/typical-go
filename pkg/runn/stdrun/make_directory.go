package stdrun

import (
	"os"
)

// Mkdir to make directory
type Mkdir struct {
	path string
}

// NewMkdir return new instance of MkdirMkdir
func NewMkdir(path string) *Mkdir {
	return &Mkdir{
		path: path,
	}
}

// Run to making the directory
func (md Mkdir) Run() error {
	return os.MkdirAll(md.path, 0700)
}
