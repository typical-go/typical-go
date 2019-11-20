package common

import (
	"github.com/typical-go/typical-go/pkg/utility/golang"
)

// WriteSource to write source to file
type WriteSource struct {
	Target string
	golang.Source
}

// Run to write source
func (w WriteSource) Run() (err error) {
	return w.WriteToFile(w.Target)
}
