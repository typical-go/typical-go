package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestGetConstructor(t *testing.T) {
	var (
		ctor  = &typast.Annot{Decl: someFunc, TagName: "constructor"}
		ctor2 = &typast.Annot{Decl: someFunc2, TagName: "constructor", TagAttrs: []byte(`{"name": "noname"}`)}
		ctor3 = &typast.Annot{Decl: someFunc3, TagName: "ctor"}
		ctor4 = &typast.Annot{Decl: someFunc4, TagName: "ctor", TagAttrs: []byte(`{invalid-json`)}
		ctor5 = &typast.Annot{Decl: someStruct, TagName: "ctor"}

		astStore = &typast.ASTStore{
			Annots: []*typast.Annot{ctor, ctor2, ctor3, ctor4, ctor5},
		}
	)

	ctors, errs := typannot.GetCtors(astStore)

	require.Equal(t, []*typannot.Ctor{
		{
			Annot: ctor,
			Param: typannot.CtorParam{},
		},
		{
			Annot: ctor2,
			Param: typannot.CtorParam{
				Name: "noname",
			},
		},
		{
			Annot: ctor3,
			Param: typannot.CtorParam{},
		},
	}, ctors)

	require.EqualError(t,
		errs.Unwrap(),
		"ctor: invalid character 'i' looking for beginning of object key string",
	)
}
