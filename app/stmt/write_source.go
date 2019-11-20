package stmt

import (
	"github.com/typical-go/typical-go/pkg/utility/golang"
)

// WriteSource to write source to file
type WriteSource struct {
	Target string
	Source golang.SourceCode
}

// Run to write source
func (w WriteSource) Run() (err error) {
	return w.Source.WriteToFile(w.Target)
}
