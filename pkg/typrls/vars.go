package typrls

import (
	"io"
	"os"
)

var (
	// DefaultSummarizer default summary
	DefaultSummarizer = &GitSummarizer{
		ExcludePrefix: []string{"merge", "bump", "revision", "generate", "wip"},
	}
	// DefaultTagger default tag
	DefaultTagger = &StdTagger{}
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)
