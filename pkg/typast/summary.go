package typast

import (
	"strings"
)

type (
	// Summary responsible to store filename, declaration and annotation
	Summary struct {
		Paths  []string
		Decls  []*Decl
		Annots []*Annot
	}
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

// AddDecl add declaration
func (s *Summary) AddDecl(file File, declType Type) {
	decl := &Decl{
		File: file,
		Type: declType,
	}
	s.Decls = append(s.Decls, decl)
	s.Annots = append(s.Annots, retrieveAnnots(decl)...)
}

// FindAnnotByFunc find annotation by function
func (s *Summary) FindAnnotByFunc(tagName string) []*Annot {
	return s.findAnnot(func(a *Annot) bool {
		_, ok := a.Type.(*FuncDecl)
		return strings.EqualFold(tagName, a.TagName) && ok
	})
}

// FindAnnotByStruct find annotation by struct
func (s *Summary) FindAnnotByStruct(tagName string) []*Annot {
	return s.findAnnot(func(a *Annot) bool {
		_, ok := a.Type.(*StructDecl)
		return strings.EqualFold(tagName, a.TagName) && ok
	})
}

// FindAnnotByInterface find annotation by interface
func (s *Summary) FindAnnotByInterface(tagName string) []*Annot {
	return s.findAnnot(func(a *Annot) bool {
		_, ok := a.Type.(*InterfaceDecl)
		return strings.EqualFold(tagName, a.TagName) && ok
	})
}

func (s *Summary) findAnnot(checkFn func(*Annot) bool) []*Annot {
	var annots []*Annot
	for _, annot := range s.Annots {
		if checkFn(annot) {
			annots = append(annots, annot)
		}
	}
	return annots
}
