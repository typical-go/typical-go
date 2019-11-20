package coll

import (
	"strings"
)

// Strings is slice of string
type Strings []string

// Append item
func (s *Strings) Append(item ...string) *Strings {
	*s = append(*s, item...)
	return s
}

// Join elements
func (s *Strings) Join(sep string) string {
	return strings.Join([]string(*s), sep)
}

// IsEmpty return true is no element
func (s *Strings) IsEmpty() bool {
	return len(*s) < 1
}
