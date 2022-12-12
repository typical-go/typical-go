package typgen

import (
	"path/filepath"
	"reflect"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Annotation that contain extra additional information
	Annotation struct {
		Name   string            `json:"name"`
		Params reflect.StructTag `json:"params"`
		Decl   *Decl             `json:"decl"`
	}
	// Decl stand of declaration
	Decl struct {
		File *File
		Type
	}
	// Type declaration type
	Type interface {
		GetName() string
		GetDocs() []string
	}
)

//
// Annotation
//

// Package of annotation
func (d *Annotation) Package() string {
	if d.Decl != nil && d.Decl.File != nil {
		return d.Decl.File.Name
	}
	return ""
}

func (d *Annotation) Path() string {
	if d.Decl != nil && d.Decl.File != nil {
		return d.Decl.File.Path
	}
	return ""
}

func (d *Annotation) Dir() string {
	path := d.Path()
	if path == "" {
		return ""
	}
	return filepath.Dir(path)
}

func (d *Annotation) PackagePath() string {
	return typgo.ProjectPkg + "/" + d.Dir()
}
