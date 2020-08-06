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
	someStruct3Decl = &typannot.Decl{
		Path:    "sample_test.go",
		Package: "typannot_test",
		Type: &typannot.StructType{
			Fields: []*typannot.Field{
				{Name: "Name", Type: "string"},
				{Name: "Address", Type: "string"},
			},
		},
		Name: "sampleStruct3",
	}
)

func TestCompile(t *testing.T) {
	store, err := typannot.Compile("sample_test.go")
	require.NoError(t, err)

	require.EqualValues(t, []*typannot.Decl{
		someInterfaceDecl,
		someStructDecl,
		someFunctionDecl,
		someFunctionDecl2,
		someInterface2Decl,
		someStruct2Decl,
		someStruct3Decl,
	}, store.Decls)

	require.EqualValues(t, []*typannot.Annot{
		{
			Decl:    someStructDecl,
			TagName: "@tag1",
		},
		{
			Decl:     someStructDecl,
			TagName:  "@tag2",
			TagParam: `key1:"", key2: "", key3:"value3"`,
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

func TestCompile_FileNotFound(t *testing.T) {
	_, err := typannot.Compile("not_found.go")
	require.EqualError(t, err, "open not_found.go: no such file or directory")
}

func TestCreateAnnotation(t *testing.T) {
	testcases := []struct {
		TestName         string
		Raw              string
		ExpectedTagName  string
		ExpectedTagAttrs string
	}{
		{
			TestName:        "tag only",
			Raw:             `@tag1`,
			ExpectedTagName: "@tag1",
		},
		{
			TestName:        "tag only with space",
			Raw:             `@tag2 extra1`,
			ExpectedTagName: "@tag2",
		},
		{
			TestName:         "with attribute",
			Raw:              `@tag3("name":"wire1")`,
			ExpectedTagName:  "@tag3",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
		{
			TestName:         "there is space between tagname and attribute",
			Raw:              `@tag4 ("name":"wire1")`,
			ExpectedTagName:  "@tag4",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			tagName, tagAttrs := typannot.ParseAnnot(tt.Raw)
			require.Equal(t, tt.ExpectedTagName, tagName)
			require.Equal(t, tt.ExpectedTagAttrs, tagAttrs)
		})
	}
}
