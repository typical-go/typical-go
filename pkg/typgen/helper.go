package typgen

import (
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	DefaultParent = "internal/generated"
)

func CreateTargetDir(path string, suffix string) string {
	dir := filepath.Dir(path)
	if dir == "." {
		return DefaultParent
	}
	dir = strings.ReplaceAll(dir, "internal/", "")
	return DefaultParent + "/" + dir + "_" + suffix
}

func IsFunc(d *Annotation) bool {
	funcDecl, ok := d.Decl.Type.(*Function)
	return ok && !funcDecl.IsMethod()
}

func IsStruct(d *Annotation) bool {
	_, ok := d.Decl.Type.(*Struct)
	return ok
}

func IsInterface(d *Annotation) bool {
	_, ok := d.Decl.Type.(*Interface)
	return ok
}

func IsPublic(d *Annotation) bool {
	rune, _ := utf8.DecodeRuneInString(d.Decl.GetName())
	return unicode.IsUpper(rune)
}

func PackageName(path string) string {
	return filepath.Base(filepath.Dir(path))
}
