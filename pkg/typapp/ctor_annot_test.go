package typapp_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typlog"
)

func TestGetCtorAnnot(t *testing.T) {
	var debugger strings.Builder
	astStore := &typast.ASTStore{
		Annots: []*typast.Annotation{
			{
				Decl: &typast.Decl{
					Name: "someFunc",
					Pkg:  "somePkg",
					Type: typast.Function,
				},
				TagName: "constructor",
			},
			{
				Decl: &typast.Decl{
					Name: "someFunc2",
					Pkg:  "somePkg",
					Type: typast.Function,
				},
				TagName:  "constructor",
				TagAttrs: []byte(`{"name": "noname"}`),
			},
			{
				Decl: &typast.Decl{
					Name: "someFunc3",
					Pkg:  "somePkg",
					Type: typast.Function,
				},
				TagName:  "constructor",
				TagAttrs: []byte(`{invalid-json`),
			},
			{
				Decl: &typast.Decl{
					Name: "someStruct",
					Pkg:  "somePkg",
					Type: typast.Struct,
				},
				TagName: "constructor",
			},
		},
	}

	c := &typbuildtool.PreconditionContext{
		Logger: typlog.Logger{Name: "PRECOND", Out: &debugger},
	}
	c.SetASTStore(astStore)

	require.Equal(t, []*typapp.CtorAnnot{
		{Name: "", Def: "somePkg.someFunc"},
		{Name: "noname", Def: "somePkg.someFunc2"},
	}, typapp.GetCtorAnnot(c))

	require.Equal(t, "PRECOND:WARN> CtorAnnot: Invalid tag attribute {invalid-json\n", debugger.String())
}
