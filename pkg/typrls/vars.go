package typrls

import (
	"io"
	"os"
)

var (
	// DefaultValidation default validation
	DefaultValidation = Validators{
		&NoGitChangeValidation{},
		&AlreadyReleasedValidation{},
		&UncommittedValidation{},
	}
	// DefaultSummary default summary
	DefaultSummary = &ChangeSummary{
		ExcludePrefix: []string{"merge", "bump", "revision", "generate", "wip"},
	}
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)
