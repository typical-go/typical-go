package typrls

import (
	"fmt"
	"strings"
)

type (
	// Summarizer responsible to create release summary
	Summarizer interface {
		Summarize(*Context) string
	}
	// SummarizeFn summary function
	SummarizeFn    func(*Context) string
	summarizerImpl struct {
		fn SummarizeFn
	}
	// GitSummarizer summary from git change log
	GitSummarizer struct {
		ExcludePrefix []string
	}
)

//
// summarizerImpl
//

// NewSummarizer return new instance of Summarizer
func NewSummarizer(fn SummarizeFn) Summarizer {
	return &summarizerImpl{fn: fn}
}

func (s *summarizerImpl) Summarize(c *Context) string {
	return s.fn(c)
}

//
// ChangeSummary
//

var _ (Summarizer) = (*GitSummarizer)(nil)

// Summarize by git change logs
func (s *GitSummarizer) Summarize(c *Context) string {
	var changes []string
	for _, log := range c.Git.Logs {
		if !s.HasPrefix(log.Message) {
			changes = append(changes, fmt.Sprintf("%s %s", log.ShortCode, log.Message))
		}
	}
	return strings.Join(changes, "\n")
}

// HasPrefix return true if eligible to excluded by prefix
func (s *GitSummarizer) HasPrefix(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range s.ExcludePrefix {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
