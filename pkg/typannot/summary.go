package typannot

import (
	"reflect"
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
		Name    string
		Path    string
		Package string
		Type    interface{}
	}
	// Annot that contain extra additional information
	Annot struct {
		TagName  string            `json:"tag_name"`
		TagParam reflect.StructTag `json:"tag_param"`
		*Decl    `json:"decl"`
	}
	// FuncType function type
	FuncType struct{}
	// InterfaceType interface type
	InterfaceType struct{}
	// StructType struct type
	StructType struct {
		Fields []*Field
	}
	// Field information
	Field struct {
		Name string
		Type string
		reflect.StructTag
	}
)

// FindAnnotByFunc find annotation by function
func (c *Summary) FindAnnotByFunc(tagName string) []*Annot {
	return c.findAnnot(func(a *Annot) bool {
		_, ok := a.Type.(*FuncType)
		return strings.EqualFold(tagName, a.TagName) && ok
	})
}

// FindAnnotByStruct find annotation by struct
func (c *Summary) FindAnnotByStruct(tagName string) []*Annot {
	return c.findAnnot(func(a *Annot) bool {
		_, ok := a.Type.(*StructType)
		return strings.EqualFold(tagName, a.TagName) && ok
	})
}

// FindAnnotByInterface find annotation by interface
func (c *Summary) FindAnnotByInterface(tagName string) []*Annot {
	return c.findAnnot(func(a *Annot) bool {
		_, ok := a.Type.(*InterfaceType)
		return strings.EqualFold(tagName, a.TagName) && ok
	})
}

func (c *Summary) findAnnot(checkFn func(*Annot) bool) []*Annot {
	var annots []*Annot
	for _, annot := range c.Annots {
		if checkFn(annot) {
			annots = append(annots, annot)
		}
	}
	return annots
}
