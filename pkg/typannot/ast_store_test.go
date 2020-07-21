package typannot_test

import (
	"fmt"
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
		Type:    &typannot.StructType{},
		Name:    "sampleStruct",
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

	for _, decl := range store.Decls {
		fmt.Println(decl)
	}

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
			TagName: "tag1",
		},
		{
			Decl:     someStructDecl,
			TagName:  "tag2",
			TagAttrs: []byte(`{"key1":"", "key2": "", "key3":"value3"}`),
		},
		{
			Decl:    someFunctionDecl2,
			TagName: "ctor",
		},
		{
			Decl:    someInterface2Decl,
			TagName: "tag3",
		},
		{
			Decl:    someStruct2Decl,
			TagName: "tag4",
		},
	}, store.Annots)
}

func TestRetrRawAnnots(t *testing.T) {
	testcases := []struct {
		testname string
		doc      string
		expected []string
	}{
		{
			testname: "no annotation",
			doc:      `no annotation`,
		},
		{
			testname: "start with @",
			doc:      "@tag1",
			expected: []string{
				"@tag1",
			},
		},
		{
			testname: "start with @, multiple annotation",
			doc:      "@tag1\n@tag2",
			expected: []string{
				"@tag1",
				"@tag2",
			},
		},
		{
			testname: "start with @, multiple annotation, end with empty line",
			doc:      "@tag1\n@tag2\n",
			expected: []string{
				"@tag1",
				"@tag2",
			},
		},
		{
			testname: "start with @, have attribute",
			doc:      "@tag1{attribute}",
			expected: []string{
				"@tag1{attribute}",
			},
		},
		{
			testname: "start with @; multiple annotation; one has attribute",
			doc:      "@tag1{attribute}\n@tag2\n",
			expected: []string{
				"@tag1{attribute}",
				"@tag2",
			},
		},
		{
			testname: "start with @; multiple annotation; all have attribute",
			doc:      "@tag1{attribute}\n@tag2{attribute}\n",
			expected: []string{
				"@tag1{attribute}",
				"@tag2{attribute}",
			},
		},
		{
			testname: "start with @; multiple annotation; all have attribute; multiline annotation",
			doc:      "@tag1{\nattribute\n}\n@tag2{\nattribute\n}\n",
			expected: []string{
				"@tag1{\nattribute\n}",
				"@tag2{\nattribute\n}",
			},
		},
		{
			testname: "start with not annotation; multiple annotation; all have attribute; multiline annotation",
			doc:      "not annotation\n@tag1{\nattribute\n}\n@tag2{\nattribute\n}\n",
			expected: []string{
				"@tag1{\nattribute\n}",
				"@tag2{\nattribute\n}",
			},
		},
		{
			testname: "start with not annotation; multiple annotation; one have attribute; multiline annotation",
			doc:      "not annotation\n@tag1\n@tag2{\nattribute\n}\n",
			expected: []string{
				"@tag1",
				"@tag2{\nattribute\n}",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testname, func(t *testing.T) {
			var initial []string
			typannot.RetrieveRawAnnots(&initial, tt.doc)
			require.Equal(t, tt.expected, initial)
		})
	}

}
