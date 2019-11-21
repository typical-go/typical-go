package common

import (
	"io"

	"github.com/typical-go/typical-go/pkg/utility/filekit"
)

// WriteSource to write source to file
type WriteSource struct {
	Target string
	Source interface {
		Write(io.Writer) error
	}
}

// Run to write source
func (w WriteSource) Run() (err error) {
	return filekit.Write(w.Target, w.Source)
}
