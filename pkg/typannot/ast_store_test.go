package typannot_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

var (
	someInterfaceDecl = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type:    &typannot.InterfaceType{},
		Name:    "sampleInterface",
	}

	someStructDecl = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type: &typannot.StructType{
			Fields: []*typannot.Field{
				{Name: "sampleInt", Type: "int", StructTag: reflect.StructTag(`default:"value1"`)},
				{Name: "sampleString", Type: "string", StructTag: reflect.StructTag(`default:"value2"`)},
			},
		},
		Name: "sampleStruct",
	}

	someFunctionDecl = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type:    &typannot.FuncType{},
		Name:    "sampleFunction",
	}

	someFunctionDecl2 = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type:    &typannot.FuncType{},
		Name:    "sampleFunction2",
	}

	someInterface2Decl = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type:    &typannot.InterfaceType{},
		Name:    "sampleInterface2",
	}

	someStruct2Decl = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type:    &typannot.StructType{},
		Name:    "sampleStruct2",
	}
)

func TestCreateASTStore(t *testing.T) {
	store, err := typannot.CreateASTStore("sample_test.go")
	require.NoError(t, err)

	require.EqualValues(t, []*typannot.Decl{
		someInterfaceDecl,
		someStructDecl,
		someFunctionDecl,
		someFunctionDecl2,
		someInterface2Decl,
		someStruct2Decl,
	}, store.Decls)

	require.EqualValues(t, []*typannot.Annot{
		{
			Decl:    someStructDecl,
			TagName: "@tag1",
		},
		{
			Decl:     someStructDecl,
			TagName:  "@tag2",
			TagAttrs: `key1:"", key2: "", key3:"value3"`,
		},
		{
			Decl:    someFunctionDecl2,
			TagName: "@ctor",
		},
		{
			Decl:    someInterface2Decl,
			TagName: "@tag3",
		},
		{
			Decl:    someStruct2Decl,
			TagName: "@tag4",
		},
	}, store.Annots)
}
