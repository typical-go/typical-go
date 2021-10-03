package typgen

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Annotator interface {
		AnnotationName() string
		IsAllowed(d *Annotation) bool
		Process(*Context) error
	}
	Context struct {
		*typgo.Context
		*InitFile
		Annotator   Annotator
		Annotations []*Annotation
	}
)

func Filter(annotations []*Annotation, annotator Annotator) []*Annotation {
	var filtered []*Annotation
	name := annotator.AnnotationName()
	for _, annot := range annotations {
		if strings.EqualFold(name, annot.Name) && annotator.IsAllowed(annot) {
			filtered = append(filtered, annot)
		}
	}
	return filtered
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
