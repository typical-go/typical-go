package typast

import (
	"unicode"
	"unicode/utf8"
)

type (
	// Decl stand of declaration
	Decl struct {
		File
		Type
	}
	// Type declaratio type
	Type interface {
		GetName() string
		GetDocs() []string
	}
	// File information
	File struct {
		Path    string
		Package string
	}
)

// IsPublic return true if decl is public access
func IsPublic(typ Type) bool {
	rune, _ := utf8.DecodeRuneInString(typ.GetName())
	return unicode.IsUpper(rune)
}
