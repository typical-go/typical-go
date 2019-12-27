package typrls

import "strings"

// Filter the commit message
type Filter interface {
	Filter([]string) []string
}

// StandardFilter is the default filter
type StandardFilter struct {
	Ignorings []string
}

// Filter the messages
func (f *StandardFilter) Filter(msgs []string) (filtereds []string) {
	filtereds = []string{}
	for _, msg := range msgs {
		msg = CleanMessage(msg)
		if !f.ignore(msg) {
			filtereds = append(filtereds, msg)
		}
	}
	return
}

func (f *StandardFilter) ignore(changelog string) bool {
	text := MessageText(changelog)
	lower := strings.ToLower(text)
	for _, word := range f.Ignorings {
		if strings.HasPrefix(lower, strings.ToLower(word)) {
			return true
		}
	}
	return false
}

// CleanMessage to clean message
func CleanMessage(message string) string {
	iCoAuthor := strings.Index(message, "Co-Authored-By")
	if iCoAuthor > 0 {
		message = message[0:strings.Index(message, "Co-Authored-By")]
	}
	message = strings.TrimSpace(message)
	return message
}

// MessageText to return text from message
func MessageText(changelog string) string {
	if len(changelog) < 7 {
		return ""
	}
	return strings.TrimSpace(changelog[7:])
}
