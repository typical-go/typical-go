package typannot

import (
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
)

// IsTag return true if annotation same type and tag names
func IsTag(a *typast.Annot, typ typast.DeclType, tagNames ...string) bool {
	if a.Decl.Type == typ {
		for _, tagName := range tagNames {
			if strings.EqualFold(tagName, a.TagName) {
				return true
			}
		}
	}
	return false
}

// IsFuncTag return true if annotation function has tag names
func IsFuncTag(a *typast.Annot, tagNames ...string) bool {
	return IsTag(a, typast.Function, tagNames...)
}
