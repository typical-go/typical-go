package typrls

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Tagger responsible to create tag
	Tagger interface {
		CreateTag(c *typgo.Context, alpha bool) string
	}
	// StdTagger standard tagger
	StdTagger struct {
		includeGitID bool // include git id in tagName e.g. v0.0.1+5339d71
		includeDate  bool // include current date in tagName e.g. v0.0.1_20201023
	}
)

//
// StdTagger
//

var _ Tagger = (*StdTagger)(nil)

// IncludeGitID tag contain GitDI
func (s *StdTagger) IncludeGitID() *StdTagger {
	s.includeGitID = true
	return s
}

// IncludeDate tag contain current date
func (s *StdTagger) IncludeDate() *StdTagger {
	s.includeDate = true
	return s
}

// CreateTag create tag
func (s *StdTagger) CreateTag(c *typgo.Context, alpha bool) string {
	tagName := "v0.0.1"
	if c.BuildSys.ProjectVersion != "" {
		tagName = fmt.Sprintf("v%s", c.BuildSys.ProjectVersion)
	}
	if s.includeGitID {
		latestGitID := latestGitID(c)
		if len(latestGitID) > 6 {
			latestGitID = latestGitID[:6]
		}
		tagName = tagName + "+" + latestGitID
	}
	if s.includeDate {
		tagName = tagName + "_" + Now().Format("20060102")
	}
	if alpha {
		tagName = tagName + "_alpha"
	}

	return tagName
}
