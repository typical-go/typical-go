package common

import (
	"sort"
	"strings"
)

// Strings is slice of string
type Strings []string

// NewStrings return new string
func NewStrings(item ...string) *Strings {
	var s Strings
	return s.Append(item...)
}

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

// Sort the slice
func (s *Strings) Sort() *Strings {
	sort.Strings(*s)
	return s
}

// Slice of string
func (s *Strings) Slice() []string {
	return *s
}
