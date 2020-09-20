package typast

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var _ CheckFn = EqualFunc
var _ CheckFn = EqualInterface
var _ CheckFn = EqualStruct

// EqualFunc return true if annot have tagName and public function
func EqualFunc(annot *Annot, tagName string) bool {
	funcDecl, ok := annot.Type.(*FuncDecl)
	return ok && strings.EqualFold(annot.TagName, tagName) &&
		IsPublic(annot) && !funcDecl.IsMethod()
}

// EqualInterface return true if annot have tagName and public interface
func EqualInterface(annot *Annot, tagName string) bool {
	_, ok := annot.Type.(*InterfaceDecl)
	return ok && strings.EqualFold(annot.TagName, tagName) && IsPublic(annot)
}

// EqualStruct return true if annot have tagName and public interface
func EqualStruct(annot *Annot, tagName string) bool {
	_, ok := annot.Type.(*StructDecl)
	return ok && strings.EqualFold(annot.TagName, tagName) && IsPublic(annot)
}

// IsPublic return true if decl is public access
func IsPublic(typ Type) bool {
	rune, _ := utf8.DecodeRuneInString(typ.GetName())
	return unicode.IsUpper(rune)
}
