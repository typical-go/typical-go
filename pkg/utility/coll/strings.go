package coll

// Strings is slice of string
type Strings []string

// Append item
func (s *Strings) Append(item ...string) *Strings {
	*s = append(*s, item...)
	return s
}
