package walker

import (
	"regexp"
	"strings"
)

// Notations is tag in godoc
type Notations []string

// ParseNotations to parse notation from document tag
func ParseNotations(doc string) (n Notations) {
	r, _ := regexp.Compile("\\[(.*?)\\]")
	for _, s := range r.FindAllString(doc, -1) {
		n = append(n, s[1:len(s)-1])
	}
	return
}

// Contain to check is tag avaible
func (n Notations) Contain(name string) bool {
	for _, notation := range n {
		if strings.ToLower(name) == strings.ToLower(notation) {
			return true
		}
	}
	return false
}
