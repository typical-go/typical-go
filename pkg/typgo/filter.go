package typgo

import "strings"

// Excluder responsible exclude
type Excluder interface {
	Exclude(string) bool
}

// PrefixExcluder is filter no-prefix
type PrefixExcluder struct {
	prefixes []string
}

// ExcludePrefix to create string filter to exlude message
func ExcludePrefix(prefixes ...string) *PrefixExcluder {
	return &PrefixExcluder{
		prefixes: prefixes,
	}
}

// Exclude the message
func (f *PrefixExcluder) Exclude(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range f.prefixes {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
