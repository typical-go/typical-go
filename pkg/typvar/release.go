package typvar

import "github.com/typical-go/typical-go/pkg/git"

var (
	// Rls contain release data
	Rls = struct {
		Alpha         bool
		Tag           string
		GitLogs       []*git.Log
		Files         []string
		ExclMsgPrefix []string
	}{
		ExclMsgPrefix: []string{
			"merge", "bump", "revision", "generate", "wip",
		},
	}
)
