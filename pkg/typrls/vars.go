package typrls

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
)
