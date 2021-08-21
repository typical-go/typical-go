package typapp_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCtorAnnot_Annotateß(t *testing.T) {
	typgo.ProjectPkg = "github.com/user/project"

	defer os.RemoveAll("internal")

	ctorAnnot := &typapp.CtorAnnot{}
	var out strings.Builder
	c := &typgo.Context{Logger: typgo.Logger{Stdout: &out}}
	defer c.PatchBash([]*typgo.MockCommand{})(t)

	directives := []*typgen.Directive{
		{
			TagName: "@ctor",
			Decl: &typgen.Decl{
				Type: &typgen.FuncDecl{Name: "NewObject"},
				File: typgen.File{Package: "pkg", Path: "project/pkg/file.go"},
			},
		},
		{
			TagName:  "@ctor",
			TagParam: `name:"obj2"`,
			Decl: &typgen.Decl{
				File: typgen.File{Package: "pkg2", Path: "project/pkg2/file.go"},
				Type: &typgen.FuncDecl{Name: "NewObject2"},
			},
		},
	}

	require.NoError(t, ctorAnnot.Process(c, directives))

	b, _ := ioutil.ReadFile("internal/generated/ctor/ctor.go")
	require.Equal(t, "package ctor\n\n/* DO NOT EDIT. This is code generated file from '@ctor' annotation. */\nimport (\n\t\"github.com/typical-go/typical-go/pkg/typapp\"\n\ta \"github.com/user/project/project/pkg\"\n\tb \"github.com/user/project/project/pkg2\"\n)\n\nfunc init(){\n\ttypapp.Provide(\"\", a.NewObject)\n\ttypapp.Provide(\"obj2\", b.NewObject2)\n}\n\n", string(b))

	require.Equal(t, "> Generate @ctor to internal/generated/ctor/ctor.go\n> go build -o /bin/goimports golang.org/x/tools/cmd/goimports\n", out.String())

}

func TestCtorAnnot_Annotate_Predefined(t *testing.T) {

	defer os.RemoveAll("folder2")

	ctorAnnot := &typapp.CtorAnnot{
		TagName: "@some-tag",
		Target:  "folder2/dest2/some-target",
	}
	var out strings.Builder
	c := &typgo.Context{Logger: typgo.Logger{Stdout: &out}}
	defer c.PatchBash([]*typgo.MockCommand{})(t)

	directives := []*typgen.Directive{
		{
			TagName: "@some-tag",
			Decl: &typgen.Decl{
				File: typgen.File{Package: "pkg"},
				Type: &typgen.FuncDecl{Name: "NewObject"},
			},
		},
	}

	require.NoError(t, ctorAnnot.Process(c, directives))

	_, err := ioutil.ReadFile("folder2/dest2/some-target")
	require.NoError(t, err)
	require.Equal(t, "> Generate @ctor to folder2/dest2/some-target\n> go build -o /bin/goimports golang.org/x/tools/cmd/goimports\n", out.String())
}
