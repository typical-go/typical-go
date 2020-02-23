package runnerkit

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

// Mkdir runner
func Mkdir(path string) Runner {
	return &mkdir{
		path: path,
	}
}

type mkdir struct {
	path string
}

func (m mkdir) Run(ctx context.Context) error {
	log.Infof("Make directory: %s", m.path)
	return os.MkdirAll(m.path, 0700)
}
