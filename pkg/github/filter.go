package github

import "strings"

// MessageFilter responsible to filter the commit message
type MessageFilter interface {
	MessageFilter(string) string
}

// NoPrefixFilter is filter no-prefix
type NoPrefixFilter struct {
	prefixes []string
}

// NoPrefix return filter with no-prefix
func NoPrefix(prefixes ...string) *NoPrefixFilter {
	return &NoPrefixFilter{
		prefixes: prefixes,
	}
}

// DefaultNoPrefix return filter with no-prefix
func DefaultNoPrefix() *NoPrefixFilter {
	return NoPrefix(
		"merge",
		"bump",
		"revision",
		"generate",
		"wip",
	)
}

// Append to return NoPrefixFilter with appended prefix
func (f *NoPrefixFilter) Append(prefixes ...string) *NoPrefixFilter {
	f.prefixes = append(f.prefixes, prefixes...)
	return f
}

// MessageFilter to filter the messages
func (f *NoPrefixFilter) MessageFilter(msg string) string {
	if f.exclude(msg) {
		return ""
	}
	return msg
}

func (f *NoPrefixFilter) exclude(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range f.prefixes {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
