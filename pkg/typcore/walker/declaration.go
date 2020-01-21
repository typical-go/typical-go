package walker

import (
	"go/ast"

	log "github.com/sirupsen/logrus"
)

// DeclType is declaration type
type DeclType string

const (
	FunctionType  = DeclType("Function")
	InterfaceType = DeclType("Interface")
	StructType    = DeclType("Struct")
	GenericType   = DeclType("Generic")
)

// DeclFunc to handle declaration
type DeclFunc func(*Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *Declaration, ann *Annotation) error

// Declaration information
type Declaration struct {
	Filename    string
	File        *ast.File
	Doc         string
	Annotations Annotations
	Type        DeclType
	SourceName  string
	SourceObj   interface{}
}

// Declarations is list of declaration
type Declarations []*Declaration

// Append event
func (d *Declarations) Append(e ...*Declaration) *Declarations {
	*d = append(*d, e...)
	return d
}

// Each to handle each declaration
func (d *Declarations) Each(fn DeclFunc) (err error) {
	for _, decl := range *d {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (d *Declarations) EachAnnotation(name string, declType DeclType, fn AnnotationFunc) (err error) {
	return d.Each(func(decl *Declaration) (err error) {
		annotation := decl.Annotations.Get(name)
		if annotation != nil {
			if decl.Type == declType {
				if err = fn(decl, annotation); err != nil {
					return
				}
			} else {
				log.Warnf("[%s] has no effect to %s:%s", name, declType, decl.SourceName)
			}
		}
		return
	})

}
