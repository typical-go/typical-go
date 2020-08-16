package typrls

import (
	"io"
	"os"
)

var (
	// DefaultValidator default validation
	DefaultValidator = Validators{
		&NoGitChangeValidation{},
		&AlreadyReleasedValidation{},
		&UncommittedValidation{},
	}
	// DefaultSummarizer default summary
	DefaultSummarizer = &ChangeSummarizer{
		ExcludePrefix: []string{"merge", "bump", "revision", "generate", "wip"},
	}
	// DefaultTagger default tag
	DefaultTagger = &StdTagger{}
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)
