package typrls

import (
	"io"
	"time"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// DefaultSummarizer default summary
	DefaultSummarizer = &GitSummarizer{
		ExcludePrefix: []string{"merge", "bump", "revision", "generate", "wip"},
	}
	// DefaultTagger default tag
	DefaultTagger = &StdTagger{}
	// Stdout standard output
	Stdout io.Writer = typgo.Stdout
	// Now is current time
	Now = time.Now
)
