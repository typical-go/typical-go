package typapp_test

import (
	"fmt"
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCtorAnnotation_CreateCtors(t *testing.T) {
	ctx := &typast.Context{
		Summary: &typast.Summary{
			Annots: []*typast.Annot{
				{
					TagName: "@ctor",
					Decl: &typast.Decl{
						Type: &typast.FuncDecl{Name: "NewObject"},
						File: typast.File{Package: "pkg", Path: "project/pkg/file.go"},
					},
				},
				{
					TagName:  "@ctor",
					TagParam: `name:"obj2"`,
					Decl: &typast.Decl{
						File: typast.File{Package: "pkg2", Path: "project/pkg2/file.go"},
						Type: &typast.FuncDecl{Name: "NewObject2"},
					},
				},
			},
		},
	}

	typgo.ProjectPkg = "github.com/typical-go/typical-go"

	ctors, imports := typapp.FindAnnotFunc(ctx, "@ctor")
	fmt.Printf("%+v\n", ctors)
	fmt.Printf("%+v\n", imports)
}
