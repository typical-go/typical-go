package typrls

import (
	"fmt"
	"strings"
)

type (
	// Summarizer responsible to create release summary
	Summarizer interface {
		Summarize(*Context) (string, error)
	}
	// SummarizeFn summary function
	SummarizeFn    func(*Context) (string, error)
	summarizerImpl struct {
		fn SummarizeFn
	}
	// ChangeSummarizer summary from git change log
	ChangeSummarizer struct {
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

func (s *summarizerImpl) Summarize(c *Context) (string, error) {
	return s.fn(c)
}

//
// ChangeSummary
//

var _ (Summarizer) = (*ChangeSummarizer)(nil)

// Summarize by git change logs
func (s *ChangeSummarizer) Summarize(c *Context) (string, error) {
	var changes []string
	for _, log := range c.Git.Logs {
		if !s.HasPrefix(log.Message) {
			changes = append(changes, fmt.Sprintf("%s %s", log.ShortCode, log.Message))
		}
	}
	return strings.Join(changes, "\n"), nil
}

// HasPrefix return true if eligible to excluded by prefix
func (s *ChangeSummarizer) HasPrefix(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range s.ExcludePrefix {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
