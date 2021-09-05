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
				File: &typgen.File{Name: "pkg", Path: "project/pkg/file.go"},
			},
		},
		{
			TagName:  "@ctor",
			TagParam: `name:"obj2"`,
			Decl: &typgen.Decl{
				File: &typgen.File{Name: "pkg2", Path: "project/pkg2/file.go"},
				Type: &typgen.FuncDecl{Name: "NewObject2"},
			},
		},
	}

	require.NoError(t, ctorAnnot.Process(c, directives))

	b, _ := ioutil.ReadFile("internal/generated/ctor/ctor.go")
	require.Equal(t, `package ctor

/* DO NOT EDIT. This is code generated file. */
import (
	"github.com/typical-go/typical-go/pkg/typapp"
	a "github.com/user/project/project/pkg"
	b "github.com/user/project/project/pkg2"
)

func init(){
	typapp.Provide("", a.NewObject)
	typapp.Provide("obj2", b.NewObject2)
}

`, string(b))

	require.Equal(t, `> Generate @ctor to internal/generated/ctor/ctor.go
> go build -o /bin/goimports golang.org/x/tools/cmd/goimports
`, out.String())

}

func TestCtorAnnot_GenerateCode(t *testing.T) {
	typgo.ProjectPkg = "github.com/user/project"
	defer func() { typgo.ProjectPkg = "" }()
	testcases := []struct {
		TestName         string
		Context          *typgo.Context
		Directive        *typgen.Directive
		ExpectedImports  map[string]string
		ExpectedInitFunc []string
		ExpectedErr      string
	}{
		{
			TestName: "a function",
			Directive: &typgen.Directive{
				TagName: "@ctor",
				Decl: &typgen.Decl{
					Type: &typgen.FuncDecl{Name: "NewObject"},
					File: &typgen.File{Name: "pkg", Path: "project/pkg/file.go"},
				},
			},
			ExpectedImports: map[string]string{
				"github.com/user/project/project/pkg":         "a",
				"github.com/typical-go/typical-go/pkg/typapp": "",
			},
			ExpectedInitFunc: []string{
				`typapp.Provide("", a.NewObject)`,
			},
		},
		{
			TestName: "a function with constructor name",
			Directive: &typgen.Directive{
				TagName:  "@ctor",
				TagParam: `name:"ctor1"`,
				Decl: &typgen.Decl{
					Type: &typgen.FuncDecl{Name: "NewObject"},
					File: &typgen.File{Name: "pkg", Path: "project/pkg/file.go"},
				},
			},
			ExpectedImports: map[string]string{
				"github.com/user/project/project/pkg":         "a",
				"github.com/typical-go/typical-go/pkg/typapp": "",
			},
			ExpectedInitFunc: []string{
				`typapp.Provide("ctor1", a.NewObject)`,
			},
		},
		{
			TestName: "a method",
			Directive: &typgen.Directive{
				TagName: "@ctor",
				Decl: &typgen.Decl{
					Type: &typgen.FuncDecl{Name: "NewObject", Recv: &typgen.FieldList{}},
					File: &typgen.File{Name: "pkg", Path: "project/pkg/file.go"},
				},
			},
			ExpectedImports: map[string]string{
				"github.com/user/project/project/pkg":         "a",
				"github.com/typical-go/typical-go/pkg/typapp": "",
			},
			ExpectedInitFunc: []string{
				"// Method 'NewObject' is not supported",
			},
		},
		{
			TestName: "a struct",
			Directive: &typgen.Directive{
				TagName: "@ctor",
				Decl: &typgen.Decl{
					File: &typgen.File{Name: "pkg", Path: "project/pkg/file.go"},
					Type: &typgen.StructDecl{
						TypeDecl: typgen.TypeDecl{
							Name: "SomeStruct",
						},
						Fields: []*typgen.Field{
							{Names: []string{"args1", "args2"}, Type: "string"},
							{Names: []string{"args3"}, Type: "int64"},
						},
					},
				},
			},
			ExpectedImports: map[string]string{
				"github.com/typical-go/typical-go/pkg/typapp": "",
			},
			ExpectedInitFunc: []string{
				"// TODO",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			ctor := &typapp.CtorAnnot{}
			err := ctor.GenerateCode(tt.Context, tt.Directive)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.ExpectedImports, ctor.Imports().Map)
				stringers := ctor.InitFunc()
				require.EqualValues(t, len(tt.ExpectedInitFunc), len(stringers), "init funct is not match")
				for i, line := range tt.ExpectedInitFunc {
					require.Equal(t, line, stringers[i].String())
				}
			}
		})
	}
}
