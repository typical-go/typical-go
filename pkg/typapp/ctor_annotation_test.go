package typapp_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/typical-go/typical-go/pkg/oskit"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCtorAnnotation_Annotate(t *testing.T) {
	typgo.ProjectPkg = "github.com/user/project"

	var out strings.Builder
	defer oskit.PatchStdout(&out)()
	defer os.RemoveAll("internal")
	defer typgo.PatchBash([]*typgo.RunExpectation{})(t)

	ctorAnnot := &typapp.CtorAnnotation{}
	ctx := &typast.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
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

	b, _ := ioutil.ReadFile("internal/generated/constructor/constructor.go")
	require.Equal(t, `package constructor

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

	require.Equal(t, "Generate @ctor to internal/generated/constructor/constructor.go\n", out.String())

}

func TestCtorAnnotation_Annotate_Predefined(t *testing.T) {
	var out strings.Builder
	defer typgo.PatchBash([]*typgo.RunExpectation{})(t)
	defer os.RemoveAll("folder2")
	defer oskit.PatchStdout(&out)()

	ctorAnnot := &typapp.CtorAnnotation{
		TagName:  "@some-tag",
		Target:   "folder2/dest2/some-target",
		Template: "some-template",
	}
	ctx := &typast.Context{
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
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
	require.Equal(t, "Generate @ctor to folder2/dest2/some-target\n", out.String())
}

func TestCtorAnnotation_Annotate_RemoveTargetWhenNoAnnotation(t *testing.T) {
	defer os.RemoveAll("folder4")

	ctorAnnot := &typapp.CtorAnnotation{
		Target: "folder4/pkg4/some-target",
	}
	ctx := &typast.Context{
		Context: &typgo.Context{},
		Summary: &typast.Summary{},
	}

	ioutil.WriteFile("folder4/pkg4/some-target", []byte("some-content"), 0777)
	require.NoError(t, ctorAnnot.Annotate(ctx))
	_, err := os.Stat("folder4/pkg4/some-target")
	require.True(t, os.IsNotExist(err))
}

func TestCtor_Stringer(t *testing.T) {
	ctor := &typapp.Ctor{Name: "some-name", Def: "some-def"}
	require.Equal(t, "{Name=some-name Def=some-def}", ctor.String())
}
