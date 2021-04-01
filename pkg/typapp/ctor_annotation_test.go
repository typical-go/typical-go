package typapp_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCtorAnnot_Annotate(t *testing.T) {
	typgo.ProjectPkg = "github.com/user/project"

	defer os.RemoveAll("internal")

	ctorAnnot := &typapp.CtorAnnot{}
	var out strings.Builder
	c := &typgo.Context{Logger: typgo.Logger{Stdout: &out}}
	defer c.PatchBash([]*typgo.MockBash{})(t)
	ctx := &typast.Context{
		Context: c,
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

	require.NoError(t, ctorAnnot.Annotate(ctx))

	b, _ := ioutil.ReadFile("internal/generated/ctor/ctor.go")
	require.Equal(t, `package ctor

/* DO NOT EDIT. This file generated due to '@ctor' annotation*/

import (
	 "github.com/typical-go/typical-go/pkg/typapp"
	a "github.com/user/project/project/pkg"
	b "github.com/user/project/project/pkg2"
)

func init() { 
	typapp.Provide("", a.NewObject)
	typapp.Provide("obj2", b.NewObject2)
}`, string(b))

	require.Equal(t, "> Generate @ctor to internal/generated/ctor/ctor.go\n> go build -o /bin/goimports golang.org/x/tools/cmd/goimports\n", out.String())

}

func TestCtorAnnot_Annotate_Predefined(t *testing.T) {

	defer os.RemoveAll("folder2")

	ctorAnnot := &typapp.CtorAnnot{
		TagName:  "@some-tag",
		Target:   "folder2/dest2/some-target",
		Template: "some-template",
	}
	var out strings.Builder
	c := &typgo.Context{Logger: typgo.Logger{Stdout: &out}}
	defer c.PatchBash([]*typgo.MockBash{})(t)
	ctx := &typast.Context{
		Context: c,
		Summary: &typast.Summary{
			Annots: []*typast.Annot{
				{
					TagName: "@some-tag",
					Decl: &typast.Decl{
						File: typast.File{Package: "pkg"},
						Type: &typast.FuncDecl{Name: "NewObject"},
					},
				},
			},
		},
	}

	require.NoError(t, ctorAnnot.Annotate(ctx))

	b, _ := ioutil.ReadFile("folder2/dest2/some-target")
	require.Equal(t, `some-template`, string(b))
	require.Equal(t, "> Generate @ctor to folder2/dest2/some-target\n> go build -o /bin/goimports golang.org/x/tools/cmd/goimports\n", out.String())
}

func TestCtor_Stringer(t *testing.T) {
	ctor := &typapp.Ctor{Name: "some-name", Def: "some-def"}
	require.Equal(t, "{Name=some-name Def=some-def}", ctor.String())
}
