package typictx

// Github contain github information
type Github struct {
	Owner    string
	RepoName string
}

// Tagging setting
type Tagging struct {
	WithGitBranch       bool
	WithLatestGitCommit bool
}
