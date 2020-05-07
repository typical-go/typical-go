package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestGetDtors(t *testing.T) {
	var (
		dtor  = &typast.Annot{Decl: someFunc, TagName: "destructor"}
		dtor2 = &typast.Annot{Decl: someFunc2, TagName: "destructor", TagAttrs: []byte(`{"name": "noname"}`)}
		dtor3 = &typast.Annot{Decl: someFunc3, TagName: "dtor"}
		dtor4 = &typast.Annot{Decl: someFunc4, TagName: "dtor", TagAttrs: []byte(`{invalid-json`)}
		dtor5 = &typast.Annot{Decl: someStruct, TagName: "dtor"}

		astStore = &typast.ASTStore{
			Annots: []*typast.Annot{dtor, dtor2, dtor3, dtor4, dtor5},
		}
	)

	dtors, errs := typannot.GetDtors(astStore)

	require.Equal(t, []*typannot.Dtor{
		{Annot: dtor},
		{Annot: dtor2},
		{Annot: dtor3},
	}, dtors)

	require.EqualError(t,
		errs.Unwrap(),
		"dtor: invalid character 'i' looking for beginning of object key string",
	)
}
