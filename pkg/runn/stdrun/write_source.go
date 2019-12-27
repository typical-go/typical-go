package stdrun

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// WriteSource to write source to file
type WriteSource struct {
	target string
	source source
}

type source interface {
	Write(io.Writer) error
}

// NewWriteSource return new instance of WriteSource
func NewWriteSource(target string, source source) *WriteSource {
	return &WriteSource{
		target: target,
		source: source,
	}
}

// Run to write source
func (w WriteSource) Run() (err error) {
	log.Infof("Write File: %s", w.target)
	var f *os.File
	if f, err = os.Create(w.target); err != nil {
		return
	}
	defer f.Close()
	return w.source.Write(f)
}
