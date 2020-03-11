package exor

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

// Mkdir execute the make dir
type Mkdir struct {
	path string
}

// NewMkdir runner
func NewMkdir(path string) *Mkdir {
	return &Mkdir{
		path: path,
	}
}

// Execute mkdir
func (m Mkdir) Execute(ctx context.Context) error {
	log.Infof("Make directory: %s", m.path)
	return os.MkdirAll(m.path, 0700)
}
