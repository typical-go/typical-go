package typgen

import (
	"reflect"
)

type (
	// Directive that contain extra additional information
	Directive struct {
		TagName  string            `json:"tag_name"`
		TagParam reflect.StructTag `json:"tag_param"`
		*Decl    `json:"decl"`
	}
	// Directives []*Directiveß
	// Decl stand of declaration
	Decl struct {
		File *File
		Type
	}
	// Type declaratio type
	Type interface {
		GetName() string
		GetDocs() []string
	}
)

//
// Directives
//

// Package of annotation
func (d *Directive) Package() string {
	if d.Decl != nil && d.Decl.File != nil {
		return d.Decl.File.Name
	}
	return ""
}

func (d *Directive) Path() string {
	if d.Decl != nil && d.Decl.File != nil {
		return d.Decl.File.Path
	}
	return ""
}
