package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	sampleInterfaceDecl = &typast.Decl{
		Path: "sample_test.go",
		Pkg:  "typast_test",
		Type: typast.Interface,
		Name: "sampleInterface",
	}

	sampleStructDecl = &typast.Decl{
		Path: "sample_test.go",
		Pkg:  "typast_test",
		Type: typast.Struct,
		Name: "sampleStruct",
	}

	sampleFunctionDecl = &typast.Decl{
		Path: "sample_test.go",
		Pkg:  "typast_test",
		Type: typast.Function,
		Name: "sampleFunction",
	}
)

func TestCreateASTStore(t *testing.T) {

	store := typast.CreateASTStore("sample_test.go")
	require.EqualValues(t, []*typast.Decl{
		sampleInterfaceDecl,
		sampleStructDecl,
		sampleFunctionDecl,
	}, store.Decls)

	require.EqualValues(t, []*typast.Annotation{
		{Decl: sampleStructDecl, TagName: "tag1", TagAttrs: map[string]string{}},
		{Decl: sampleStructDecl, TagName: "tag2", TagAttrs: map[string]string{
			"key1": "",
			"key2": "",
			"key3": "value3",
		}},
	}, store.Annots)

}
