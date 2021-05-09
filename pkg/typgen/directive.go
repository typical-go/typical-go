package typgen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Directive that contain extra additional information
	Directive struct {
		TagName  string            `json:"tag_name"`
		TagParam reflect.StructTag `json:"tag_param"`
		*Decl    `json:"decl"`
	}
	Directives []*Directive
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

//
// Directives
//

// AddDecl add declaration
func (s *Directives) AddDecl(file File, declType Type) {
	decl := &Decl{
		File: file,
		Type: declType,
	}
	// s.Decls = append(s.Decls, decl)
	*s = append(*s, retrieveAnnots(decl)...)
}

func retrieveAnnots(decl *Decl) []*Directive {
	var annots []*Directive
	for _, raw := range decl.GetDocs() {
		if strings.HasPrefix(raw, "//") {
			raw = strings.TrimSpace(raw[2:])
		}
		if strings.HasPrefix(raw, "@") {
			tagName, tagAttrs := ParseRawAnnot(raw)
			annots = append(annots, &Directive{
				TagName:  tagName,
				TagParam: reflect.StructTag(tagAttrs),
				Decl:     decl,
			})
		}
	}

	return annots
}
