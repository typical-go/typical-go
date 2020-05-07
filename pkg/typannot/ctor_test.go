package typannot_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestGetConstructor(t *testing.T) {
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

	ctors, errs := typannot.GetCtor(astStore)

	require.Equal(t, []*typannot.Ctor{
		{Name: "", Def: "somePkg.someFunc"},
		{Name: "noname", Def: "somePkg.someFunc2"},
	}, ctors)

	// require.EqualValues(t, []error{
	// 	errors.New(""),
	// }, errs)

	fmt.Println(errs[0].Error())

}
