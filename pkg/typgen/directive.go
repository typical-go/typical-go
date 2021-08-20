package typgen

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/typical-go/typical-go/pkg/typgo"
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

//
// Directives
//

// Package of annotation
func (d *Directive) Package() string {
	return fmt.Sprintf("%s/%s", typgo.ProjectPkg, filepath.Dir(d.Path))
}
