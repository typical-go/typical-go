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
				Type: &typgen.Function{Name: "NewObject"},
				File: &typgen.File{Name: "pkg", Path: "project/pkg/file.go"},
			},
		},
		{
			TagName:  "@ctor",
			TagParam: `name:"obj2"`,
			Decl: &typgen.Decl{
				File: &typgen.File{Name: "pkg2", Path: "project/pkg2/file.go"},
				Type: &typgen.Function{Name: "NewObject2"},
			},
		},
	}

	require.NoError(t, ctorAnnot.Process(c, directives))

	b, _ := ioutil.ReadFile("internal/generated/ctor/ctor.go")
	require.Equal(t, `package ctor

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	a "github.com/user/project/project/pkg"
	b "github.com/user/project/project/pkg2"
)
// DO NOT EDIT. Code-generated file.
func init(){
	typapp.Provide("", a.NewObject)
	typapp.Provide("obj2", b.NewObject2)
}
`, string(b))

	require.Equal(t, `> Generate @ctor to internal/generated/ctor/ctor.go
> go build -o /bin/goimports golang.org/x/tools/cmd/goimports
`, out.String())

}
