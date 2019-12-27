package stdrun

import (
	"os"

	log "github.com/sirupsen/logrus"
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
func (m Mkdir) Run() error {
	log.Infof("Make directory: %s", m.path)
	return os.MkdirAll(m.path, 0700)
}
