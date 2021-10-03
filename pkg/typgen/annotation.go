package typgen

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Annotation interface {
		TagName() string
		IsAllowed(d *Directive) bool
		Process(*Context) error
	}
	Context struct {
		*typgo.Context
		*InitFile
		Annot Annotation
		Dirs  []*Directive
	}
)

func Filter(dirs []*Directive, annot Annotation) []*Directive {
	var filtered []*Directive
	tagName := annot.TagName()
	for _, dir := range dirs {
		if strings.EqualFold(tagName, dir.TagName) && annot.IsAllowed(dir) {
			filtered = append(filtered, dir)
		}
	}
	return filtered
}

func IsFunc(d *Directive) bool {
	funcDecl, ok := d.Type.(*Function)
	return ok && !funcDecl.IsMethod()
}

func IsStruct(d *Directive) bool {
	_, ok := d.Type.(*Struct)
	return ok
}

func IsInterface(d *Directive) bool {
	_, ok := d.Type.(*Interface)
	return ok
}

func IsPublic(d *Directive) bool {
	rune, _ := utf8.DecodeRuneInString(d.GetName())
	return unicode.IsUpper(rune)
}
