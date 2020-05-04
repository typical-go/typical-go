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

	sampleInterface2Decl = &typast.Decl{
		Path: "sample_test.go",
		Pkg:  "typast_test",
		Type: typast.Interface,
		Name: "sampleInterface2",
	}

	sampleStruct2Decl = &typast.Decl{
		Path: "sample_test.go",
		Pkg:  "typast_test",
		Type: typast.Struct,
		Name: "sampleStruct2",
	}
)

func TestCreateASTStore(t *testing.T) {
	store, err := typast.CreateASTStore("sample_test.go")
	require.NoError(t, err)

	cnt := len(store.Decls)
	require.Equal(t, len(store.DeclNodes), cnt)
	require.Equal(t, len(store.Docs), cnt)

	require.EqualValues(t, []*typast.Decl{
		sampleInterfaceDecl,
		sampleStructDecl,
		sampleFunctionDecl,
		sampleInterface2Decl,
		sampleStruct2Decl,
	}, store.Decls)

	require.EqualValues(t, []*typast.Annotation{
		{
			Decl:    sampleStructDecl,
			TagName: "tag1",
		},
		{
			Decl:     sampleStructDecl,
			TagName:  "tag2",
			TagAttrs: []byte(`{"key1":"", "key2": "", "key3":"value3"}`),
		},
		{
			Decl:    sampleInterface2Decl,
			TagName: "tag3",
		},
		{
			Decl:    sampleStruct2Decl,
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
			typast.RetrRawAnnots(&initial, tt.doc)
			require.Equal(t, tt.expected, initial)
		})
	}

}
