package typrls

import (
	"time"
)

var (
	// DefaultSummarizer default summary
	DefaultSummarizer = &GitSummarizer{
		ExcludePrefix: []string{"merge", "bump", "revision", "generate", "wip"},
	}
	// DefaultTagger default tag
	DefaultTagger = &StdTagger{}
	// Now is current time
	Now = time.Now
)
