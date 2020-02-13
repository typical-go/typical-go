package stdrelease

import "strings"

// Filter the commit message
type Filter interface {
	Filter(string) string
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

// Filter the messages
func (f *NoPrefixFilter) Filter(msg string) string {
	msg = cleanMessage(msg)
	if f.exclude(msg) {
		return ""
	}
	return msg
}

func (f *NoPrefixFilter) exclude(changelog string) bool {
	var (
		text  = messageText(changelog)
		lower = strings.ToLower(text)
	)
	for _, prefix := range f.prefixes {
		if strings.HasPrefix(lower, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

func cleanMessage(message string) string {
	iCoAuthor := strings.Index(message, "Co-Authored-By")
	if iCoAuthor > 0 {
		message = message[0:strings.Index(message, "Co-Authored-By")]
	}
	message = strings.TrimSpace(message)
	return message
}

// MessageText to return text from message
func messageText(changelog string) string {
	if len(changelog) < 7 { // TODO: use token instead of fixed length
		return ""
	}
	return strings.TrimSpace(changelog[7:])
}
